package stage10

import (
	"context"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func Buildtest(ctx context.Context) (r compose.Runnable[map[string]any, *schema.Message], err error) {
	const (
		ChatTemplate1 = "ChatTemplate1"
		ChatModel1    = "ChatModel1"
	)
	g := compose.NewGraph[map[string]any, *schema.Message]()
	chatTemplate1KeyOfChatTemplate, err := newChatTemplate(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddChatTemplateNode(ChatTemplate1, chatTemplate1KeyOfChatTemplate)
	chatModel1KeyOfChatModel, err := newChatModel(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddChatModelNode(ChatModel1, chatModel1KeyOfChatModel)
	_ = g.AddEdge(compose.START, ChatTemplate1)
	_ = g.AddEdge(ChatModel1, compose.END)
	_ = g.AddEdge(ChatTemplate1, ChatModel1)
	r, err = g.Compile(ctx, compose.WithGraphName("test"))
	if err != nil {
		return nil, err
	}
	return r, err
}
