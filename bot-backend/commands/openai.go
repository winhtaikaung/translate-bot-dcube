package commands

import (
	"context"
	"fmt"
	"strings"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/translate-bot-dcube/bot-backend/utils"
)

// return if the question is those we support and is the question need to be translate
func CheckLanguageSupport(question string) (bool, bool) {
	client := gogpt.NewClient(utils.GetOpenAIAPIKey())
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:  gogpt.GPT3TextDavinci003,
		Prompt: fmt.Sprintf("What language '%q' is?", question),
	}
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Println(err)
		return false, false
	}
	if strings.Contains(resp.Choices[0].Text, "Burmese") || strings.Contains(resp.Choices[0].Text, "Chinese") || strings.Contains(resp.Choices[0].Text, "Mandarin") {
		return true, true
	} else if strings.Contains(resp.Choices[0].Text, "English") {
		return true, false
	} else {
		return false, false
	}
}

func TranslateByOpenAI(question string) string {
	support, translateNeeded := CheckLanguageSupport(question)
	if support {
		if translateNeeded {
			client := gogpt.NewClient(utils.GetOpenAIAPIKey())
			ctx := context.Background()

			req := gogpt.CompletionRequest{
				Model:  gogpt.GPT3TextDavinci003,
				Prompt: fmt.Sprintf("Translate '%q' into English", question),
			}
			resp, err := client.CreateCompletion(ctx, req)
			if err != nil {
				fmt.Println(err)
				return "Sorry, currently we are facing some problem in translation. Please try again later."
			}
			return resp.Choices[0].Text
		}
		return question
	}
	return "Sorry, currently we only support Burmese, English, Chinese or Mandarin."
}
