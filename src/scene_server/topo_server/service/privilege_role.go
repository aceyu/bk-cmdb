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

package service

import (
	frtypes "configcenter/src/common/mapstr"
	"configcenter/src/scene_server/topo_server/core/types"
)

// CreatePrivilege search user goup
func (s *Service) CreatePrivilege(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (interface{}, error) {

	datas := make([]string, 0)
	data.ForEach(func(key string, val interface{}) {
		datas = append(datas, key)
	})

	err := s.core.PermissionOperation().Role(params).CreatePermission(params.Header.OwnerID, pathParams("bk_obj_id"), pathParams("bk_property_id"), datas)
	return nil, err
}

// GetPrivilege search user goup
func (s *Service) GetPrivilege(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (interface{}, error) {
	return s.core.PermissionOperation().Role(params).GetPermission(params.Header.OwnerID, pathParams("bk_obj_id"), pathParams("bk_property_id"))
}