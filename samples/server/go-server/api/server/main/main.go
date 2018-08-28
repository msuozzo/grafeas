// Copyright 2017 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"log"

	"github.com/grafeas/grafeas/samples/server/go-server/api/server/api"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/config"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/storage"
	server "github.com/grafeas/grafeas/server-go"
)

var (
	configFile = flag.String("config", "", "Path to a config file")
)

func main() {
	flag.Parse()
	config, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load config file: %s", err)
	}
	var storager server.Storager
	switch config.StorageType {
	case "memstore":
		storager = storage.NewMemStore()
	case "postgres":
		storager = storage.NewPgSQLStore(config.PgSQLConfig)
	case "embedded":
		storager = storage.NewEmbeddedStore(config.EmbeddedConfig)
	default:
		log.Fatalf("Storage type unsupported: %s", config.StorageType)
	}
	api.Run(config.API, &storager)
}
