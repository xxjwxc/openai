package openai

import (
	"time"

	"github.com/xxjwxc/openai/chatgpt"
	"github.com/xxjwxc/openai/dalle"
)

// NewChatGPT 新建一个chartgpt
func NewChatGPT(ApiKey, UserId string, timeOut time.Duration) *chatgpt.ChatGPT {
	return chatgpt.New(ApiKey, UserId, timeOut)
}

// NewDallE 新建一个图片库
func NewDallE(ApiKey, UserId string, timeOut time.Duration) *dalle.DallE {
	return dalle.NewDallE(ApiKey, UserId, timeOut)
}
