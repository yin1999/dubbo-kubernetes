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

package v3

import (
	envoy_core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"

	util_proto "github.com/apache/dubbo-kubernetes/pkg/util/proto"
	envoy_tls "github.com/apache/dubbo-kubernetes/pkg/xds/envoy/tls/v3"
)

type ServerSideStaticTLSConfigurer struct {
	CertPath string
	KeyPath  string
}

var _ FilterChainConfigurer = &ServerSideStaticTLSConfigurer{}

func (c *ServerSideStaticTLSConfigurer) Configure(filterChain *envoy_listener.FilterChain) error {
	tlsContext := envoy_tls.StaticDownstreamTlsContextWithPath(c.CertPath, c.KeyPath)

	pbst, err := util_proto.MarshalAnyDeterministic(tlsContext)
	if err != nil {
		return err
	}
	filterChain.TransportSocket = &envoy_core.TransportSocket{
		Name: "envoy.transport_sockets.tls",
		ConfigType: &envoy_core.TransportSocket_TypedConfig{
			TypedConfig: pbst,
		},
	}
	return nil
}
