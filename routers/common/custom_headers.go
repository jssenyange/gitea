// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package common

import (
	"net/http"

	"code.gitea.io/gitea/modules/setting"
)

func CustomHeadersHandler(headers map[string]setting.CustomHeader) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			for _, header := range headers{
				resp.Header().Set(header.Name, header.Value)
			}
			next.ServeHTTP(resp, req)
		})
	}
}