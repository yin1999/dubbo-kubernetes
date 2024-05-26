/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handler

import (
	"context"
	mesh_proto "github.com/apache/dubbo-kubernetes/api/mesh/v1alpha1"
	"github.com/apache/dubbo-kubernetes/pkg/admin/model"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/apis/mesh"
	res_model "github.com/apache/dubbo-kubernetes/pkg/core/resources/model"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/store"
	core_runtime "github.com/apache/dubbo-kubernetes/pkg/core/runtime"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ConfiguratorSearchWithPath set in context as key to mark searchPath
type ConfiguratorSearchWithPath struct{}

func ConfiguratorSearch(rt core_runtime.Runtime) gin.HandlerFunc {
	return func(c *gin.Context) {
		resList := &mesh.DynamicConfigResourceList{}
		if err := rt.ResourceManager().List(rt.AppContext(), resList); err != nil {
			c.JSON(http.StatusBadRequest, model.NewErrorResp(err.Error()))
			return
		}
		resp := model.ConfiguratorSearchResp{
			Code:    200,
			Message: "success",
			Data:    make([]model.ConfiguratorSearchResp_Data, 0, len(resList.Items)),
		}
		for _, item := range resList.Items {
			resp.Data = append(resp.Data, model.ConfiguratorSearchResp_Data{
				RuleName:   item.Meta.GetName(),
				Scope:      item.Spec.GetScope(),
				CreateTime: item.Meta.GetCreationTime().String(),
				Enabled:    item.Spec.GetEnabled(),
			})
		}
		c.JSON(http.StatusOK, resp)
	}
}

func GetConfiguratorWithRuleName(rt core_runtime.Runtime) gin.HandlerFunc {
	return func(c *gin.Context) {
		ruleName := c.Param("ruleName")
		res := &mesh.DynamicConfigResource{}
		ctx := context.WithValue(rt.AppContext(), ConfiguratorSearchWithPath{}, ruleName)
		if err := rt.ResourceManager().Get(ctx, res, store.GetByKey(res_model.DefaultMesh, res_model.DefaultMesh)); err != nil {
			c.JSON(http.StatusBadRequest, model.NewErrorResp(err.Error()))
			return
		}
		c.JSON(http.StatusOK, model.GenDynamicConfigToResp(http.StatusOK, "success", res.Spec))
	}
}

func PutConfiguratorWithRuleName(rt core_runtime.Runtime) gin.HandlerFunc {
	return func(c *gin.Context) {
		ruleName := c.Param("ruleName")
		res := &mesh.DynamicConfigResource{
			Meta: nil,
			Spec: &mesh_proto.DynamicConfig{},
		}
		err := c.Bind(res.Spec)
		if err != nil {
			c.JSON(http.StatusBadRequest, model.NewErrorResp(err.Error()))
			return
		}
		ctx := context.WithValue(rt.AppContext(), ConfiguratorSearchWithPath{}, ruleName)
		if err = rt.ResourceManager().Update(ctx, res, store.UpdateByKey(res_model.DefaultMesh, res_model.DefaultMesh)); err != nil {
			c.JSON(http.StatusBadRequest, model.NewErrorResp(err.Error()))
			return
		} else {
			c.JSON(http.StatusOK, model.GenDynamicConfigToResp(http.StatusOK, "success", nil))
		}
	}
}

func PostConfiguratorWithRuleName(rt core_runtime.Runtime) gin.HandlerFunc {
	return func(c *gin.Context) {
		ruleName := c.Param("ruleName")
		res := &mesh.DynamicConfigResource{
			Meta: nil,
			Spec: &mesh_proto.DynamicConfig{},
		}

		err := c.Bind(res.Spec)
		if err != nil {
			c.JSON(http.StatusBadRequest, model.NewErrorResp(err.Error()))
			return
		}
		ctx := context.WithValue(rt.AppContext(), ConfiguratorSearchWithPath{}, ruleName)
		if err = rt.ResourceManager().Create(ctx, res, store.CreateByKey(res_model.DefaultMesh, res_model.DefaultMesh)); err != nil {
			c.JSON(http.StatusBadRequest, model.NewErrorResp(err.Error()))
			return
		} else {
			c.JSON(http.StatusCreated, model.GenDynamicConfigToResp(http.StatusCreated, "success", nil))
		}
	}
}

func DeleteConfiguratorWithRuleName(rt core_runtime.Runtime) gin.HandlerFunc {
	return func(c *gin.Context) {
		ruleName := c.Param("ruleName")
		res := &mesh.DynamicConfigResource{}
		ctx := context.WithValue(rt.AppContext(), ConfiguratorSearchWithPath{}, ruleName)
		if err := rt.ResourceManager().Delete(ctx, res, store.DeleteByKey(res_model.DefaultMesh, res_model.DefaultMesh)); err != nil {
			c.JSON(http.StatusBadRequest, model.NewErrorResp(err.Error()))
			return
		}
		c.JSON(http.StatusNoContent, model.GenDynamicConfigToResp(http.StatusNoContent, "success", nil))
	}
}
