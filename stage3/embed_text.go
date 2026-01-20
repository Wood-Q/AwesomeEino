package stage3

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

func EmbedText() {
	ctx := context.Background()
	apiKey := os.Getenv("ARK_API_KEY")
	endpointID := "doubao-embedding-vision-250615" // 需要用多模态向量模型
	timeout := 60 * time.Second

	// 初始化 Ark Runtime 客户端
	client := arkruntime.NewClientWithApiKey(apiKey,
		arkruntime.WithTimeout(timeout),
	)

	encodingFormat := model.EmbeddingEncodingFormatFloat
	texts := []string{
		"这是第一段示例文本",
		"这是第二段示例文本",
	}
	for i, text := range texts {
		req := model.MultiModalEmbeddingRequest{
			Model:          endpointID,
			EncodingFormat: &encodingFormat,
			Input: []model.MultimodalEmbeddingInput{
				{Type: model.MultiModalEmbeddingInputTypeText, Text: &text},
			},
		}

		resp, err := client.CreateMultiModalEmbeddings(ctx, req)
		if err != nil {
			panic(fmt.Errorf("第 %d 条文本 embedding 失败: %w", i+1, err))
		}

		if len(resp.Data.Embedding) == 0 {
			panic(fmt.Errorf("第 %d 条文本返回空向量", i+1))
		}
		fmt.Println("文本", i+1, "的向量维度:", len(resp.Data.Embedding))
	}
}
