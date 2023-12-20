/*
 * Copyright (c) 2023, LinstoHu
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package marketws

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fastjson"
)

type Request struct {
	ID    string `json:"id,omitempty"`
	Sub   string `json:"sub,omitempty"`
	UnSub string `json:"unsub,omitempty"`
}

// AnyMessage represents either a JSON Response or SubscribedMessage.
type AnyMessage struct {
	Ping              *PingMessage
	Response          *Response
	SubscribedMessage *SubscribedMessage
}

type PingMessage struct {
	Ping int64 `json:"ping,omitempty"`
}

type Response struct {
	ID     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Subbed string `json:"subbed,omitempty"`
	Ts     int64  `json:"ts,omitempty"`
}

type SubscribedMessage struct {
	Channel string          `json:"ch,omitempty"`
	Ts      int64           `json:"ts,omitempty"`
	Data    json.RawMessage `json:"tick"`
}

func (m AnyMessage) MarshalJSON() ([]byte, error) {
	var v any

	switch {
	case m.Response != nil && m.SubscribedMessage == nil:
		v = m.Response
	case m.Response == nil && m.SubscribedMessage != nil:
		v = m.SubscribedMessage
	}

	if v != nil {
		return json.Marshal(v)
	}

	return nil, errors.New("message must have exactly one of the Response or SubscribedMessage fields set")
}

func (m *AnyMessage) UnmarshalJSON(data []byte) error {
	var p fastjson.Parser
	v, err := p.ParseBytes(data)
	if err != nil {
		return err
	}

	if v.Exists("ping") {
		var resp PingMessage

		if err := json.Unmarshal(data, &resp); err != nil {
			return err
		}

		m.Ping = &resp

		return nil
	}

	if v.Exists("id") {
		var resp Response

		if err := json.Unmarshal(data, &resp); err != nil {
			return err
		}

		m.Response = &resp

		return nil
	}

	if v.Exists("ch") {
		msg := &SubscribedMessage{
			Channel: string(v.GetStringBytes("ch")),
			Ts:      v.GetInt64("ts"),
		}

		if v.Get("tick") != nil {
			msg.Data = v.Get("tick").MarshalTo(nil)
		}

		m.SubscribedMessage = msg

		return nil
	}

	return nil
}
