package api

func (c *Client) GetHAGroups() ([]string, error) {
	var groups []string
	err := c.get("/cluster/ha/groups", nil, groups)
	return groups, err
}

func (c *Client) GetHAGroup(group string) (string, error) {
	var groupConfig string
	err := c.get("/cluster/ha/groups/"+group, nil, &groupConfig)
	return groupConfig, err
}

func (c *Client) GetHAResources() ([]string, error) {
	var resources []string
	err := c.get("/cluster/ha/resources", nil, &resources)
	return resources, err
}

type HAStatus struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Type    string `json:"type"`
	Node    string `json:"node"`
	Quorate int    `json:"quorate"`
}

func (c *Client) GetHACurrentStatus() (HAStatus, error) {
	var haStatus HAStatus
	err := c.get("/cluster/ha/status/current", nil, &haStatus)
	return haStatus, err
}

type HAManagerStatus struct {
	ManagerStatus struct {
		NodeStatus interface{} `json:"node_status"`
	} `json:"manager_status"`
	Quorum struct {
		Node    string `json:"node"`
		Quorate string `json:"quorate"`
	} `json:"quorum"`
}

func (c *Client) GetHAManagerStatus() (HAManagerStatus, error) {
	var status HAManagerStatus
	err := c.get("/cluster/ha/status/manager_status", nil, &status)
	return status, err
}
