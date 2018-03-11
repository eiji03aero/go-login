package action

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/eiji03aero/go-login/db"
	"github.com/eiji03aero/go-login/util"
	"github.com/urfave/cli"
)

func ShowAllAccounts(c *cli.Context) (cont bool) {
	accounts := db.GetAllAccounts()

	for i, a := range accounts {
		util.WaitMS(200)
		fmt.Printf("========== User#%d [%s] ==========\n", i, a.Name)
		a.Introduce()
		util.NewLine()
	}

	return true
}

func SearchAccounts(c *cli.Context) (cont bool) {
	var _cont bool
	fmt.Println("Type the name of the user that you wanna search...")
	userInput := prompt.Input(">> ", util.DefaultCompleter)

	util.NewLine()

	if account, err := db.GetAccount(string(userInput)); err != true {
		fmt.Println("========== Here comes the hero!! ==========")
		account.Introduce()
		_cont = true
	} else {
		fmt.Printf("No such account found man! [%v]\n", userInput)
		_cont = false
	}

	util.NewLine()
	util.WaitMS(200)

	return _cont
}
