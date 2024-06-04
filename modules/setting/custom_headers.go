// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"strings"
)

type CustomHeader struct {
	Name  string
	Value string
}

// Custom headers settings
var CustomHeaders = struct {
	API map[string]CustomHeader
	WEB map[string]CustomHeader
}{
	API: make(map[string]CustomHeader),
	WEB: make(map[string]CustomHeader),
}

func loadCustomHeaders(rootCfg ConfigProvider) {
	sec := rootCfg.Section("web_custom_headers")
	mapSectionHeaders(sec, CustomHeaders.WEB)

	sec = rootCfg.Section("api_custom_headers")
	mapSectionHeaders(sec, CustomHeaders.API)
}

func mapSectionHeaders(section ConfigSection, sectionHeaders map[string]CustomHeader) {	
	for _, header := range section.Keys() {
		if len(header.Value()) > 0 {
			headerName := header.Name()			
			customHeader := CustomHeader{}
			customHeader.Name = headerName
			customHeader.Value = header.Value()
			sectionHeaders[strings.ToLower(headerName)] = customHeader
		}
	}
}
