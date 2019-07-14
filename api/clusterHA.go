package api

func (c *Client) GetHAGroups() ([]string, error) {
    var groups []string
    err := c.get("/cluster/ha/groups", nil, groups)
    return groups, err
}

func (c *Client) GetHAGroup(group string) (string, error) {
    var groupConfig string
    err := c.get("/cluster/ha/groups/" + group, nil, &groupConfig)
    return groupConfig, err
}

func (c *Client) GetHAResources() ([]string, error) {
    var resources []string
    err := c.get("/cluster/ha/resources", nil, &resources)
    return resources, err
}

