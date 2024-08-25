package api

import util_http "github.com/apache/dubbo-kubernetes/pkg/util/http"

type AuthnPlugin interface {
	Validate(map[string]string) error
	DecorateClient(util_http.Client, map[string]string) (util_http.Client, error)
}
