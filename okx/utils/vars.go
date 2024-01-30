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

package utils

var (
	RestURL      = "https://www.okx.com"
	PublicWsURL  = "wss://ws.okx.com:8443/ws/v5/public"
	PrivateWsURL = "wss://ws.okx.com:8443/ws/v5/private"

	AWSRestURL      = "https://aws.okx.com"
	AWSPublicWsURL  = "wss://wsaws.okx.com:8443/ws/v5/public"
	AWSPrivateWsURL = "wss://wsaws.okx.com:8443/ws/v5/private"
)

type InstrumentType = string

const (
	Spot     = "SPOT"
	Margin   = "MARGIN"
	Swap     = "SWAP"
	Futures  = "FUTURES"
	Option   = "OPTION"
	Cash     = "cash"
	Isolated = "isolated"
	Cross    = "cross"
	Buy      = "buy"
	Sell     = "sell"
	Market   = "market"
	Limit    = "limit"
	PostOnly = "post_only"
	FOK      = "fok"
	IOC      = "ioc"
	Index    = "index"
)
