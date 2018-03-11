package model

type DbScheme struct {
	Meta struct {
		Version      string `json:"version"`
		LastModified string `json:"last_modified"`
	}
	AccountData struct {
		MemberCount int       `json:"member_count"`
		Accounts    []Account `json:"accounts"`
	} `json:"account_data"`
}
