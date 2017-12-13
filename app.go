package main

import (
	"fmt"
)

func main() {
	answers, err := askMezzanineDetails()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	mezzanineURL := answers["url"]
	mezzanineKey := answers["oauth2-key"]

	answers, err = askGhostDetails()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ghostURL := answers["url"]
	ghostAuthMethod := answers["auth-choice"]

	switch ghostAuthMethod {
	case "username+password":
		answers, err = askGhostLogin()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	case "API Token":
		answers, err = askGhostToken()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	fmt.Printf("URL: %s, OAuth2 key: %s", mezzanineURL, mezzanineKey)
	fmt.Printf("URL: %s, Auth method: %s", ghostURL, ghostAuthMethod)
}
