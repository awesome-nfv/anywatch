/*
 * Copyright (C) 2016 Red Hat, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package packet_injector

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	shttp "github.com/connetos/anywatch/http"
	"github.com/connetos/anywatch/logging"
)

type PacketInjectorReply struct {
	TrackingID string
	err        error
}

type PacketInjectorClient struct {
	shttp.DefaultWSServerEventHandler
	WSServer       *shttp.WSServer
	replyChanMutex sync.RWMutex
	replyChan      map[string]chan PacketInjectorReply
}

func (pc *PacketInjectorClient) OnMessage(c *shttp.WSClient, m shttp.WSMessage) {
	pc.replyChanMutex.RLock()
	defer pc.replyChanMutex.RUnlock()

	ch, ok := pc.replyChan[m.UUID]
	if !ok {
		logging.GetLogger().Errorf("Unable to send reply, chan not found for %s, available: %v", m.UUID, pc.replyChan)
		return
	}

	var reply PacketInjectorReply
	if err := json.Unmarshal([]byte(*m.Obj), &reply); err != nil {
		ch <- PacketInjectorReply{err: fmt.Errorf("Failed to parse response from %s: %s", c.Host, err.Error())}
		return
	}

	ch <- reply
}

func (pc *PacketInjectorClient) injectPacket(host string, pp *PacketParams) (string, error) {
	msg := shttp.NewWSMessage(Namespace, "PIRequest", pp)

	ch := make(chan PacketInjectorReply)
	defer close(ch)

	pc.replyChanMutex.Lock()
	pc.replyChan[msg.UUID] = ch
	pc.replyChanMutex.Unlock()

	defer func() {
		pc.replyChanMutex.Lock()
		delete(pc.replyChan, msg.UUID)
		pc.replyChanMutex.Unlock()
	}()

	if !pc.WSServer.SendWSMessageTo(msg, host) {
		return "", fmt.Errorf("Unable to send message to agent: %s", host)
	}

	var reply PacketInjectorReply
	select {
	case reply = <-ch:
		return reply.TrackingID, reply.err
	case <-time.After(time.Second * 10):
		return "", fmt.Errorf("Timeout while reading PIReply from: %s", host)
	}
}

func (pc *PacketInjectorClient) InjectPacket(host string, pp *PacketParams) (string, error) {
	trackingID, err := pc.injectPacket(host, pp)
	if err != nil {
		logging.GetLogger().Errorf(err.Error())
		return "", err
	}

	return trackingID, err
}

func NewPacketInjectorClient(w *shttp.WSServer) *PacketInjectorClient {
	pic := &PacketInjectorClient{
		WSServer:  w,
		replyChan: make(map[string]chan PacketInjectorReply),
	}
	w.AddEventHandler(pic, []string{Namespace})

	return pic
}
