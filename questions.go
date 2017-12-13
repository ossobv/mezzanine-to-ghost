package main

import (
	"github.com/AlecAivazis/survey"
)

var mezzanineQuestions = []*survey.Question{
	{
		Name:      "url",
		Prompt:    &survey.Input{Message: "URL of Mezzanine CMS/blog"},
		Validate:  IsURL,
		Transform: survey.ToLower, // force URLs to be lowercase
	},
	{
		Name:   "oauth2-key",
		Prompt: &survey.Input{Message: "Mezzanine API OAuth2 key"},
	},
}

var ghostQuestions = []*survey.Question{
	{
		Name:      "url",
		Prompt:    &survey.Input{Message: "URL of Ghost blog"},
		Validate:  IsURL,
		Transform: survey.ToLower, // force URLs to be lowercase
	},
	{
		Name: "auth-choice",
		Prompt: &survey.Select{
			Message: "Authentication to Ghost",
			Options: []string{"username+password", "API Token"},
		},
	},
}

var ghostLoginQuestions = []*survey.Question{
	{
		Name: "username",
		Prompt: &survey.Input{
			Message: "Ghost username",
		},
	},
	{
		Name: "password",
		Prompt: &survey.Password{
			Message: "Ghost password",
		},
	},
}

var ghostTokenQuestions = []*survey.Question{
	{
		Name: "token",
		Prompt: &survey.Password{
			Message: "Ghost API Token",
		},
	},
}

func askMezzanineDetails() (map[string]string, error) {
	questions := mezzanineQuestions
	answers := struct {
		URL       string `survey:"url"`
		OAuth2Key string `survey:"oauth2-key"`
	}{}

	// perform the questions
	err := survey.Ask(questions, &answers)
	if err != nil {
		return nil, err
	}

	kv := map[string]string{}

	kv["url"] = answers.URL
	kv["oauth2-key"] = answers.OAuth2Key

	return kv, nil
}

func askGhostDetails() (map[string]string, error) {
	questions := ghostQuestions
	answers := struct {
		URL        string `survey:"url"`
		AuthChoice string `survey:"auth-choice"`
	}{}

	// perform the questions
	err := survey.Ask(questions, &answers)
	if err != nil {
		return nil, err
	}

	kv := map[string]string{}

	kv["url"] = answers.URL
	kv["auth-choice"] = answers.AuthChoice

	return kv, nil
}

func askGhostLogin() (map[string]string, error) {
	return nil, nil
}

func askGhostToken() (map[string]string, error) {
	return nil, nil
}
