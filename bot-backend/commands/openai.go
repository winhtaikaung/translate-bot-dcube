package commands

import (
	"context"
	"fmt"
	"log"
	"strings"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/translate-bot-dcube/bot-backend/responses"
	"github.com/translate-bot-dcube/bot-backend/utils"
)

var BURMESE = "Burmese"
var CHINESE_MANDARIN = "Mandarin"

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

	if contains(resp.Choices[0].Text, "burmese") {
		support, translateNeeded, language = true, true, "Burmese"
	} else if contains(resp.Choices[0].Text, "chinese") || strings.Contains(resp.Choices[0].Text, "mandarin") {
		support, translateNeeded, language = true, true, "Chinese"
	} else if contains(resp.Choices[0].Text, "Malay") {
		support, translateNeeded, language = true, true, "Malay"
	} else if contains(resp.Choices[0].Text, "bahasa indonesia") {
		support, translateNeeded, language = true, true, "Bahasa Indonesia"
	} else if contains(resp.Choices[0].Text, "tagalog") || contains(resp.Choices[0].Text, "filipino") {
		support, translateNeeded, language = true, true, "Tagalog"
	} else if contains(resp.Choices[0].Text, "tamil") {
		support, translateNeeded, language = true, true, "Tamil"
	} else if contains(resp.Choices[0].Text, "English") {
		support, translateNeeded, language = true, false, "English"
	}
	return support, translateNeeded, language
}

// return if translated, translated/error text, language
func TranslateToMotherTongueByOpenAI(inPutText string, language string, translateNeeded bool) string {

	if translateNeeded {
		client := gogpt.NewClient(utils.GetOpenAIAPIKey())
		ctx := context.Background()

		req := gogpt.CompletionRequest{
			Model:     gogpt.GPT3TextDavinci003,
			Prompt:    fmt.Sprintf("Translate '%q' into %s", inPutText, language),
			MaxTokens: 256,
		}
		resp, err := client.CreateCompletion(ctx, req)
		if err != nil {
			log.Println(err)
			return "Sorry, currently we are facing some problem in translation. Please try again later."
		}
		return resp.Choices[0].Text
	}
	return inPutText

}

// return translated/error text
func IdentifyQuestionFromInputText(name string, question string) string {
	// if language == "" {
	// 	language = "English"
	// }
	support, translateNeeded, language := CheckLanguageSupport(question)
	if support {
		client := gogpt.NewClient(utils.GetOpenAIAPIKey())
		ctx := context.Background()

		req := gogpt.CompletionRequest{
			Model: gogpt.GPT3TextDavinci003,
			Prompt: fmt.Sprintf(`I have this question:
		%s
		Translate to English and tell me which of the following text matches the closest,		
		- What is my work pass status?
		- What is my short term visit pass staus?
		- How to fill up SG Arrival card?
		- How to apply short term visit pass ?
		- How to apply long term visit pass ?
		and reply here:
		`, question),
			MaxTokens: 256,
		}
		resp, err := client.CreateCompletion(ctx, req)
		if err != nil {
			log.Println(err)
			return "Sorry, currently we are facing some problem in translation. Please try again later."
		}
		log.Println(resp.Choices)
		return ResponseServiceByLanguage(name, resp.Choices[0].Text, language, translateNeeded)

	} else {
		return "Oops our bot is learning his best to serve services in more languages"
	}
}

func contains(match string, pattern string) bool {
	return strings.Contains(strings.ToLower(match), pattern)
}

func ResponseServiceByLanguage(name string, questionResp string, language string, translateNeeded bool) string {

	if contains(questionResp, "work pass status") {
		body := &responses.QueryStatusRequestBody{
			Name: name, Identifier: "",
		}

		return TranslateToMotherTongueByOpenAI(responses.QueryMomResponse(body, "Employment Pass"), language, translateNeeded)
	} else if contains(questionResp, "how to fill up sg arrival card") {
		body := &responses.QueryStatusRequestBody{
			Name: name, Identifier: "",
		}

		return TranslateToMotherTongueByOpenAI(responses.QueryICAHowToResponse(body, "entry"), language, translateNeeded)
	} else if contains(questionResp, "what is my short term visit pass staus") {
		body := &responses.QueryStatusRequestBody{
			Name: name, Identifier: "",
		}

		return TranslateToMotherTongueByOpenAI(responses.QueryICAResponse(body, "short-term-pass"), language, translateNeeded)
	} else if contains(questionResp, "how to apply short term visit") {
		body := &responses.QueryStatusRequestBody{
			Name: name, Identifier: "",
		}

		return TranslateToMotherTongueByOpenAI(responses.QueryICAHowToResponse(body, "short-term-visit"), language, translateNeeded)
	} else if contains(questionResp, "how to apply long term visit") {
		body := &responses.QueryStatusRequestBody{
			Name: name, Identifier: "",
		}

		return TranslateToMotherTongueByOpenAI(responses.QueryICAHowToResponse(body, "long-term-visit"), language, translateNeeded)
	} else {
		return "Oops our bot is learning his best to serve services in more languages"
	}
}
