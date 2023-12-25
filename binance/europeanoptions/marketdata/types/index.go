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

package types

import "github.com/linstohu/nexapi/utils"

type GetUnderlyingIndexPriceParams struct {
	Underlying string `url:"underlying" validate:"required"`
}

type GetUnderlyingIndexPriceResp struct {
	Http *utils.ApiResponse
	Body *UnderlyingIndexPrice
}

type UnderlyingIndexPrice struct {
	Time       int64  `json:"time"`
	IndexPrice string `json:"indexPrice"`
}
