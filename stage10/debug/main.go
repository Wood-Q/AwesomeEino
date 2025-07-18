/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"AwesomeEino/stage10"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudwego/eino-ext/devops"
)

func main() {
	ctx := context.Background()

	// Init eino devops server
	err := devops.Init(ctx)
	if err != nil {
		log.Printf("[eino dev] init failed, err=%v", err)
		return
	}

	// Register chain, graph and state_graph for demo use
	_, err = stage10.Buildtest(ctx)
	if err != nil {
		log.Printf("[eino dev] build test failed, err=%v", err)
		return
	}
	// Blocking process exits
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Exit
	log.Printf("[eino dev] shutting down\n")
}
