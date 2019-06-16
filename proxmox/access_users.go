package proxmox

func (c *Client)GetUsers() (list map[string]interface{}, err error) {
    err = c.GetJsonRetryable("/access/users", &list, 3)
    return
}

func (c *Client)GetUser(userid string) (user map[string]interface{}, err error) {
    err = c.GetJsonRetryable("/access/users/" + userid, &user, 3)
    return
}
