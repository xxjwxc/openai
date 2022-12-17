package openai

import (
	"fmt"
	"testing"
	"time"
)

var chatGpt_token = "sk-"

func TestChatgpt(t *testing.T) {
	chat := NewChatGPT(chatGpt_token, "aaaaa", 10*time.Second)
	defer chat.Close()
	//select {
	//case <-chat.GetDoneChan():
	//	fmt.Println("time out")
	//}
	question := "现在你是一只猫，接下来你只能用\"喵喵喵\"回答."
	fmt.Printf("Q: %s\n", question)
	answer, err := chat.ChatWithContext(question)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("A: %s\n", answer)

	question = "你是一只猫吗？"
	fmt.Printf("Q: %s\n", question)
	answer, err = chat.ChatWithContext(question)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("A: %s\n", answer)

	// Q: 现在你是一只猫，接下来你只能用"喵喵喵"回答.
	// A: 喵喵喵！
	// Q: 你是一只猫吗？
	// A: 喵喵~!
}

func TestDalle(t *testing.T) {
	dalle := NewDallE(chatGpt_token, "aaaaa", 10*time.Second)
	err, list := dalle.GenPhotoBase64("晚霞", 1, "512x512")
	fmt.Println(err, list)
}
