/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v3

import (
	"configcenter/src/framework/common"
	"configcenter/src/framework/core/types"
	//"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
)

// CreateGroup create a group
func (cli *Client) CreateGroup(data types.MapStr) (int, error) {

	targetURL := fmt.Sprintf("%s/api/v3/objectattr/group/new", cli.GetAddress())

	rst, err := cli.httpCli.POST(targetURL, nil, data.ToJSON())
	if nil != err {
		return 0, err
	}

	gs := gjson.ParseBytes(rst)

	// check result
	if !gs.Get("result").Bool() {
		return 0, errors.New(gs.Get("bk_error_msg").String())
	}

	// parse id
	id := gs.Get("data.id").Int()

	return int(id), nil
}

// DeleteGroup delete a group by condition
func (cli *Client) DeleteGroup(cond common.Condition) error {

	data := cond.ToMapStr()
	id, err := data.Int("id")
	if nil != err {
		return err
	}

	targetURL := fmt.Sprintf("%s/api/v3/objectattr/group/groupid/%d", cli.GetAddress(), id)

	rst, err := cli.httpCli.DELETE(targetURL, nil, nil)
	if nil != err {
		return err
	}

	gs := gjson.ParseBytes(rst)

	// check result
	if !gs.Get("result").Bool() {
		return errors.New(gs.Get("bk_error_msg").String())
	}

	return nil
}

// UpdateGroup update a group by condition
func (cli *Client) UpdateGroup(data types.MapStr, cond common.Condition) error {

	return nil
}

// SearchGroups search some group by condition
func (cli *Client) SearchGroups(cond common.Condition) ([]types.MapStr, error) {
	return nil, nil
}
