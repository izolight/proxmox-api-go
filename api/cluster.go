package api

type ACMEAccount struct {
	Account   interface{} `json:"account,omitempty"`
	Directory string      `json:"directory,omitempty"`
	Location  string      `json:"location,omitempty"`
	TOS       string      `json:"tos,omitempty"`
}

func (c *Client) GetACMEAccounts() ([]ACMEAccount, error) {
	acmeAccounts := []ACMEAccount{}
	err := c.get("/cluster/acme/account", nil, &acmeAccounts)
	return acmeAccounts, err
}

func (c *Client) GetACMEAccount(name string) (ACMEAccount, error) {
	acmeAccount := ACMEAccount{}
	err := c.get("/cluster/acme/account/" + name, nil, &acmeAccount)
	return acmeAccount, err
}

type ACMEDirectory struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

func (c *Client) GetACMEDirectory() (ACMEDirectory, error) {
	acmeDirectory := ACMEDirectory{}
	err := c.get("/cluster/acme/directories", nil, &acmeDirectory)
	return acmeDirectory, err
}

func (c *Client) GetACMEToS(url string) (string, error) {
	var tos string
	err := c.get("/cluster/acme/tos", map[string]string{"directory": url}, &tos)
	return tos, err
}

func (c *Client) GetBackupSchedule() ([]string, error) {
	var schedule []string
	err := c.get("/cluster/backup", nil, &schedule)
	return schedule, err
}

func (c *Client) GetBackupJobDefinition(id string) (string, error) {
	var backupJobDef string
	err := c.get("/cluster/backup/" + id, nil, &backupJobDef)
	return backupJobDef, err
}