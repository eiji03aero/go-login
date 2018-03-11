package action

import (
	"errors"
	"fmt"
	"os"

	"github.com/eiji03aero/go-login/util"
	"github.com/urfave/cli"
)

func AppBefore(c *cli.Context) error {
	if len(os.Args) > 3 {
		err := errors.New("You are asking for too much man")
		return err
	}
	fmt.Println("Welcome to go-login!🎉 🎉")
	util.NewLine()
	util.WaitMS(500)

	return nil
}

func AppAfter(c *cli.Context) error {
	util.WaitMS(500)

	fmt.Println("Bye!")
	return nil
}
