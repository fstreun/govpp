//  Copyright (c) 2020 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package binapigen

import (
	"fmt"
	"strings"
)

type Plugin struct {
	Name         string
	GenerateFile func(*Generator)
}

var registeredPlugins = map[string]*Plugin{}

func (g *Generator) RunPlugin(name string) error {
	plugin, ok := registeredPlugins[name]
	if !ok {
		return fmt.Errorf("plugin not found: %q", name)
	}
	plugin.GenerateFile(g)
	return nil
}

func RegisterPlugin(name string, genfn func(*Generator)) {
	if name == "" {
		panic("plugin name empty")
	}
	if _, ok := registeredPlugins[name]; ok {
		panic("plugin name reused")
	}
	registeredPlugins[name] = &Plugin{
		Name:         name,
		GenerateFile: genfn,
	}
}

func GetAvailablePluginNames() string {
	s := make([]string, 0)
	for k, _ := range registeredPlugins {
		s = append(s, k)
	}
	return strings.Join(s, ",")
}
