package model

import (
	"context"
	"os"

	"github.com/baidubce/bce-qianfan-sdk/go/qianfan"
)

type Model struct {
	access_key string
	secret_key string
	model_name string
}

func NewModel(access_key, secret_key, model_name string) *Model {
	return &Model{
		access_key: access_key,
		secret_key: secret_key,
		model_name: model_name,
	}
}

func (m *Model) Input(text string) string {
	os.Setenv("QIANFAN_ACCESS_KEY", m.access_key)
	os.Setenv("QIANFAN_SECRET_KEY", m.secret_key)

	chat := qianfan.NewChatCompletion(
		qianfan.WithModel(m.model_name),
	)
	response, err := chat.Do(
		context.TODO(),
		&qianfan.ChatCompletionRequest{
			Messages: []qianfan.ChatCompletionMessage{
				qianfan.ChatCompletionUserMessage(text),
			},
		},
	)
	if err != nil {
		panic(err)
	}
	return response.Result
}
