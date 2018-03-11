package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/eiji03aero/go-login/model"
)

const BUFSIZE = 1024
const DBPATH = "./db/accounts.json"

func GetAllAccounts() []model.Account {
	data := fetchData()
	return data.AccountData.Accounts
}

func GetAccount(name string) (account model.Account, err bool) {
	nameReg := regexp.MustCompile(name)
	data := fetchData()

	for _, a := range data.AccountData.Accounts {
		if nameReg.MatchString(a.Name) {
			return a, false
		}
	}

	return model.Account{}, true
}

func fetchData() (data model.DbScheme) {
	var resultData model.DbScheme
	bytes, dbErr := ioutil.ReadFile(DBPATH)

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	if jsonErr := json.Unmarshal(bytes, &resultData); jsonErr != nil {
		fmt.Println("Error happend: ", jsonErr)
	}
	return resultData
}
