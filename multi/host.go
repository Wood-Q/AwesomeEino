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
	"context"
	"fmt"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/flow/agent/multiagent/host"
)
 
 func newHost(ctx context.Context) (*host.Host, error) {
	arkAPIKey := "56a6b406-8b6b-4bb5-b169-92117a5caa72"
	arkModelName := "doubao-1-5-pro-32k-250115"
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: arkAPIKey,
		Model:  arkModelName,
	})
	if err != nil {
		fmt.Printf("failed to create chat model: %v", err)
		return nil,err
	}

	return &host.Host{
		ChatModel:    chatModel,
		SystemPrompt: "你可以同时计算加法和减法。当用户提问时，你需要回答问题。",
	}, nil
}

 