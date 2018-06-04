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
 
package models

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	dbStorage "configcenter/src/storage"
)

func InitSystemData(tableName string, insCli dbStorage.DI) error {
	blog.Errorf("add data for  %s table ", tableName)
	data := map[string]interface{}{
		common.HostCrossBizField: common.HostCrossBizValue}
	isExist, err := insCli.GetCntByCondition(tableName, data)
	if nil != err {
		blog.Errorf("add data for  %s table error  %s", tableName, err)
		return err
	}
	if isExist > 0 {
		return nil
	}
	_, err = insCli.Insert(tableName, data)
	if nil != err {
		blog.Errorf("add data for  %s table error  %s", tableName, err)
		return err
	}

	return nil

	blog.Errorf("add data for  %s table  ", tableName)
	return nil
}

func ModifySystemData(tableName, ownerID string, insCli dbStorage.DI) error {
	blog.Errorf("modify data for  %s table ", tableName)
	cond := map[string]interface{}{
		common.HostCrossBizField: common.HostCrossBizValue}
	data := map[string]interface{}{
		common.HostCrossBizField: common.HostCrossBizValue + ownerID}

	err := insCli.UpdateByCondition(tableName, data, cond)
	if nil != err {
		blog.Errorf("modify data for  %s table error  %s", tableName, err)
		return err
	}

	return nil
}
