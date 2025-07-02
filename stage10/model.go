package stage10

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/model"
)

// newChatModel component initialization function of node 'ChatModel1' in graph 'test'
func newChatModel(ctx context.Context) (cm model.BaseChatModel, err error) {
	// TODO Modify component configuration here.
	config := &ark.ChatModelConfig{
		Model:  os.Getenv("MODEL"),
		APIKey: os.Getenv("ARK_API_KEY"),
	}
	cm, err = ark.NewChatModel(ctx, config)
	if err != nil {
		return nil, err
	}
	return cm, nil
}
