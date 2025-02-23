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

import "github.com/jl1/nexapi/mexc/utils"

type TransferParam struct {
	FromAccountType string `url:"fromAccountType" validate:"required,oneof=SPOT FUTURES"`
	ToAccountType   string `url:"toAccountType" validate:"required,oneof=SPOT FUTURES"`
	Asset           string `url:"asset" validate:"required"`
	Amount          string `url:"amount" validate:"required"`
}

type TransferParams struct {
	TransferParam
	utils.DefaultParam
}
