package commands

import (
	"context"
	"fmt"
	"log"
	"strings"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/translate-bot-dcube/bot-backend/utils"
)

// return if the question is those we support, is the question need to be translate and language
func CheckLanguageSupport(question string) (bool, bool, string) {
	support, translateNeeded, language := false, false, ""

	client := gogpt.NewClient(utils.GetOpenAIAPIKey())
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:  gogpt.GPT3TextDavinci003,
		Prompt: fmt.Sprintf("What language '%q' is?", question),
	}
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		log.Println(err)
		return support, translateNeeded, language
	}

	if strings.Contains(resp.Choices[0].Text, "Burmese") {
		support, translateNeeded, language = true, true, "Burmese"
	} else if strings.Contains(resp.Choices[0].Text, "Chinese") || strings.Contains(resp.Choices[0].Text, "Mandarin") {
		support, translateNeeded, language = true, true, "Chinese"
	} else if strings.Contains(resp.Choices[0].Text, "Malay") {
		support, translateNeeded, language = true, true, "Malay"
	} else if strings.Contains(resp.Choices[0].Text, "Tamil") {
		support, translateNeeded, language = true, true, "Tamil"
	} else if strings.Contains(resp.Choices[0].Text, "English") {
		support, translateNeeded, language = true, false, "English"
	}
	return support, translateNeeded, language
}

// return if translated, translated/error text, language
func TranslateToEnglishByOpenAI(question string) (bool, string, string) {
	support, translateNeeded, language := CheckLanguageSupport(question)
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
				log.Println(err)
				return false, "Sorry, currently we are facing some problem in translation. Please try again later.", "English"
			}
			return true, resp.Choices[0].Text, language
		}
		return true, question, language
	}
	return false, "Sorry, currently we only support Burmese, English, Chinese/Mandarin, Malay and Tamil.", "English"
}

// return translated/error text
func TranslateBackByOpenAI(question string, language string) string {
	if language == "" {
		language = "English"
	}
	client := gogpt.NewClient(utils.GetOpenAIAPIKey())
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:  gogpt.GPT3TextDavinci003,
		Prompt: fmt.Sprintf("Translate '%q' into %s", question, language),
	}
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		log.Println(err)
		return "Sorry, currently we are facing some problem in translation. Please try again later."
	}
	return resp.Choices[0].Text
}
