package api

import (
	"net/http"
)

type TicketResponse struct {
	CSRFPreventionToken string `json:"CSRFPreventionToken,omitempty"`
	Clustername         string `json:"clustername,omitempty"`
	Ticket              string `json:"ticket,omitempty"`
	Username            string `json:"username"`
}

func (c *Client) Login(username, password string) error {
	req, err := c.newRequest(http.MethodPost, "/access/ticket", map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return err
	}
	ticket := &TicketResponse{}
	wrapper := &wrapper{
		Data: ticket,
	}
	_, err = c.do(req, wrapper)
	if err != nil {
		return err
	}
	c.ticket = ticket.Ticket
	c.csrfToken = ticket.CSRFPreventionToken

	return nil
}

type Domain struct {
	Comment string `json:"comment,omitempty"`
	Realm   string `json:"realm,omitempty"`
	Type    string `json:"type,omitempty"`
	Plugin  string `json:"plugin,omitempty"`
	Digest  string `json:"digest,omitempty"`
	TFA     string `json:"tfa,omitempty"`
}

func (c *Client) GetDomains() ([]Domain, error) {
	domains := []Domain{}
	err := c.get("/access/domains", nil, &domains)
	if err != nil {
		return nil, err
	}
	return domains, nil
}

func (c *Client) GetDomain(domain string) (Domain, error) {
	d := Domain{}
	err := c.get("/access/domains/" + domain, nil, &d)
	if err != nil {
		return d, err
	}
	return d, nil
}

type Group struct {
	Comment string `json:"comment,omitempty"`
	GroupId string `json:"groupid,omitempty"`
	Members []string `json:"members,omitempty"`
}

func (c *Client) GetGroups() ([]Group, error) {
	groups := []Group{}
	err := c.get("/access/groups", nil, &groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (c *Client) GetGroup(groupid string) (Group, error) {
	group := Group{}
	err := c.get("/access/groups/" + groupid, nil, &group)
	if err != nil {
		return group, err
	}
	return group, nil
}

type Role struct {
	Privs   string `json:"privs,omitempty"`
	RoleId  string `json:"roleid"`
	Special int    `json:"special,omitempty"`
}

func (c *Client) GetRoles() ([]Role, error) {
	roles := []Role{}
	err := c.get("/access/roles", nil, &roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

type User struct {
	Comment   string `json:"comment,omitempty"`
	Email     string `json:"email,omitempty"`
	Enable    int    `json:"enable,omitempty"`
	Expire    int    `json:"expire,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Keys      string `json:"keys,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Userid    string `json:"userid"`
}

func (c *Client) GetUsers() ([]User, error) {
	users := []User{}
	err := c.get("/access/users", nil, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

type ACL struct {
	Path      string `json:"path"`
	Propagate int    `json:"propagate,omitempty"`
	RoleId    string `json:"roleid"`
	Type      string `json:"type"`
	UGID      string `json:"ugid"`
}

func (c *Client) GetACLs() ([]ACL, error) {
	acls := []ACL{}
	err := c.get("/access/acl", nil, &acls)
	if err != nil {
		return nil, err
	}
	return acls, nil
}
