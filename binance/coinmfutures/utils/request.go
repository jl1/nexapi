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

package cmutils

import (
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
)

type HTTPRequest struct {
	// SecurityType each endpoint has a security type that determines how you will interact with it
	// docs: https://binance-docs.github.io/apidocs/delivery/en/#endpoint-security-type
	SecurityType usdmutils.SecurityType

	BaseURL string
	Path    string
	Method  string
	Headers map[string]string
	Query   any
	Body    any
}
