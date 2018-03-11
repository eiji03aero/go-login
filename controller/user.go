package controller

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/eiji03aero/go-login/action"
	"github.com/eiji03aero/go-login/util"
	"github.com/urfave/cli"
)

func UserController(c *cli.Context) {
	fmt.Println("Please select what you wanna do!! ctrl-i to select option")
	fmt.Println("Available Actions: [index, search, quit]")
	userInput := prompt.Input(">> ", userControllerCompleter)

	util.NewLine()

	switch userInput {
	case "index":
		if cont := action.ShowAllAccounts(c); cont == true {
			UserController(c)
		}

	case "search":
		if cont := action.SearchAccounts(c); cont == true {
			UserController(c)
		}

	case "quit":
		fmt.Println("Looking forward to seeing you next time!")

	default:
		fmt.Println("I dont know what you sayin!")
	}
}

func userControllerCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "index", Description: "Show all the users"},
		{Text: "search", Description: "Search user"},
		{Text: "quit", Description: "Quit this awesome app"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
