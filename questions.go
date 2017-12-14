package main

import (
	"fmt"

	"github.com/AlecAivazis/survey"
	"upper.io/db.v3/mysql"
)

var mezzanineQuestions = []*survey.Question{
	{
		Name:   "db-user",
		Prompt: &survey.Input{Message: "Mezzanine database user"},
	},
	{
		Name:   "db-pass",
		Prompt: &survey.Password{Message: "Mezzanine database password"},
	},
	{
		Name:   "db-name",
		Prompt: &survey.Input{Message: "Mezzanine database name"},
	},
	{
		Name:   "db-host",
		Prompt: &survey.Input{Message: "Mezzanine database host", Default: "localhost"},
	},
	{
		Name:   "db-port",
		Prompt: &survey.Input{Message: "Mezzanine database port", Default: "3306"},
	},
}

var ghostQuestions = []*survey.Question{
	{
		Name:   "db-user",
		Prompt: &survey.Input{Message: "Ghost database user"},
	},
	{
		Name:   "db-pass",
		Prompt: &survey.Password{Message: "Ghost database password"},
	},
	{
		Name:   "db-name",
		Prompt: &survey.Input{Message: "Ghost database name"},
	},
	{
		Name:   "db-host",
		Prompt: &survey.Input{Message: "Ghost database host", Default: "localhost"},
	},
	{
		Name:   "db-port",
		Prompt: &survey.Input{Message: "Ghost database port", Default: "3306"},
	},
}

func askDatabaseDetails(questions []*survey.Question) (*mysql.ConnectionURL, error) {
	answers := struct {
		User string `survey:"db-user"`
		Pass string `survey:"db-pass"`
		Name string `survey:"db-name"`
		Host string `survey:"db-host"`
		Port string `survey:"db-port"`
	}{}

	// perform the questions
	err := survey.Ask(questions, &answers)
	if err != nil {
		return nil, err
	}

	settings := mysql.ConnectionURL{
		User:     answers.User,
		Password: answers.Pass,
		Host:     fmt.Sprintf("%s:%s", answers.Host, answers.Port),
		Database: answers.Name,
		Options: map[string]string{
			"charset": "utf8mb4",
		},
	}

	return &settings, nil
}

func mezzanineDBConfig() (*mysql.ConnectionURL, error) {
	return askDatabaseDetails(mezzanineQuestions)
}

func ghostDBConfig() (*mysql.ConnectionURL, error) {
	return askDatabaseDetails(ghostQuestions)
}

func askGhostLogin() (map[string]string, error) {
	return nil, nil
}

func askGhostToken() (map[string]string, error) {
	return nil, nil
}
