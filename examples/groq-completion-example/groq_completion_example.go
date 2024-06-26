package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ankit-arora/langchaingo/llms"
	"github.com/ankit-arora/langchaingo/llms/openai"
)

func main() {
	llm, err := openai.New(
		openai.WithModel("llama3-8b-8192"),
		openai.WithBaseURL("https://api.groq.com/openai/v1"),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	_, err = llms.GenerateFromSinglePrompt(ctx,
		llm,
		"Write a long poem about how golang is a fantastic language.",
		llms.WithTemperature(0.8),
		llms.WithMaxTokens(4096),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	fmt.Println()
	if err != nil {
		log.Fatal(err)
	}
}
