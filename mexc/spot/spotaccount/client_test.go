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

package spotaccount

import (
	"context"
	"os"
	"testing"

	"github.com/rluisr/nexapi/mexc/spot/spotaccount/types"
	spotutils "github.com/rluisr/nexapi/mexc/spot/utils"
	"github.com/stretchr/testify/assert"
)

func testNewAccountClient(t *testing.T) *SpotAccountClient {
	cli, err := NewSpotAccountClient(&SpotAccountClientCfg{
		BaseURL: spotutils.BaseURL,
		Key:     os.Getenv("MEXC_KEY"),
		Secret:  os.Getenv("MEXC_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create mexc client, %s", err)
	}

	return cli
}

func TestGetAccountInfo(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetAccountInfo(context.TODO())
	assert.Nil(t, err)
}

func TestTransfer(t *testing.T) {
	cli := testNewAccountClient(t)

	err := cli.Transfer(context.TODO(), types.TransferParam{
		FromAccountType: "SPOT",
		ToAccountType:   "FUTURES",
		Asset:           "USDT",
		Amount:          "5",
	})
	assert.Nil(t, err)
}
