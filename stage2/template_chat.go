package stage2

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func TemplateChat() {
	ctx := context.Background()
	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}"),
		schema.MessagesPlaceholder("history_key", false),
		&schema.Message{
			Role:    schema.User,
			Content: "请帮帮我，史瓦罗先生，{task}",
		},
	)
	params := map[string]any{
		"role":        "机器人史瓦罗先生",
		"task":        "写一首诗",
		"history_key": []*schema.Message{{Role: schema.User, Content: "告诉我油画是什么?"}, {Role: schema.Assistant, Content: "油画是xxx"}},
	}
	messages, err := template.Format(ctx, params)
	if err != nil {
		panic(err)
	}

	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
	if err != nil {
		panic(err)
	}

	answer, err := model.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}

	print(answer.Content)
}
