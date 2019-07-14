package api

import "strconv"

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
	err := c.get("/cluster/acme/account/"+name, nil, &acmeAccount)
	return acmeAccount, err
}

type ACMEDirectory struct {
	Name string `json:"name"`
	URL  string `json:"url"`
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
	err := c.get("/cluster/backup/"+id, nil, &backupJobDef)
	return backupJobDef, err
}

type LogEntry struct {
	User string `json:"user"`
	ID   string `json:"id"`
	PID  uint   `json:"pid"`
	Pri  int    `json:"pri"`
	Msg  string `json:"msg"`
	Time uint   `json:"time"`
	UID  string `json:"uid"`
	Tag  string `json:"tag"`
	Node string `json:"node"`
}

func (c *Client) GetClusterLog(max int) ([]LogEntry, error) {
	var entries []LogEntry
	err := c.get("/cluster/log", map[string]string{"max": strconv.Itoa(max)}, &entries)
	return entries, err
}

func (c *Client) GetNextVMID() (string, error) {
	var id string
	err := c.get("/cluster/nextid", nil, &id)
	return id, err
}

type Resource struct {
	CPUUtilization float64 `json:"cpu,omitempty"`
	DiskSpaceUsed  uint    `json:"disk,omitempty"`
	HAState        string  `json:"hastate,omitempty"`
	ID             string  `json:"id"`
	SupportLevel   string  `json:"level,omitempty"`
	MaxCPUs        uint    `json:"maxcpu,omitempty"`
	MaxDisk        uint    `json:"maxdisk,omitempty"`
	MaxMem         uint    `json:"maxmem,omitempty"`
	MemUsed        uint    `json:"mem,omitempty"`
	Node           string  `json:"node,omitempty"`
	Pool           string  `json:"pool,omitempty"`
	Status         string  `json:"status,omitempty"`
	StorageID      string  `json:"storage,omitempty"`
	Shared         uint    `json:"shared,omitempty"`
	Type           string  `json:"type"`
	Uptime         uint    `json:"uptime,omitempty"`
}

func (c *Client) GetResources() ([]Resource, error) {
	var resources []Resource
	err := c.get("/cluster/resources", nil, &resources)
	return resources, err
}

//func (c *Client) GetVMResources
//func (c *Client) GetStorageResources
//func (c *Client) GetNodeResources

func (c *Client) GetClusterStatus() ([]interface{}, error) {
	var status []interface{}
	err := c.get("/cluster/status", nil, &status)
	return status, err
}

type Task struct {
	Node      string `json:"node"`
	Saved     string `json:"saved"`
	EndTime   uint   `json:"endtime,omitempty"`
	StartTime uint   `json:"starttime"`
	UpID      string `json:"upid"`
	Type      string `json:"type"`
	ID        string `json:"id"`
	Status    string `json:"status"`
	User      string `json:"user"`
}

func (c *Client) GetClusterTasks() ([]Task, error) {
	var tasks []Task
	err := c.get("/cluster/tasks", nil, &tasks)
	return tasks, err
}
