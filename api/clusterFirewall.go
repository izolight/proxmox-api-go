package api

import "strconv"

type FirewallAlias struct {
    CIDR    string `json:"cidr"`
    Comment string `json:"comment"`
    Digest  string `json:"digest"`
    Name    string `json:"name"`
}

func (c *Client) GetFirewallAliases() ([]FirewallAlias, error) {
    firewallAliases := []FirewallAlias{}
    err := c.get("/cluster/firewall/aliases", nil, &firewallAliases)
    return firewallAliases, err
}

func (c *Client) GetFirewallAlias(name string) (FirewallAlias, error) {
    firewallAlias := FirewallAlias{}
    err := c.get("/cluster/firewall/aliases/"+name, nil, &firewallAlias)
    return firewallAlias, err
}

type FirewallGroup struct {
    Comment string `json:"comment"`
    Digest  string `json:"digest"`
    Group   string `json:"group"`
}

func (c *Client) GetFirewallGroups() ([]FirewallGroup, error) {
    firewallGroups := []FirewallGroup{}
    err := c.get("/cluster/firewall/groups", nil, &firewallGroups)
    return firewallGroups, err
}

func (c *Client) GetFirewallGroup(group string) (FirewallGroup, error) {
    firewallGroup := FirewallGroup{}
    err := c.get("/cluster/firewall/groups/"+group, nil, &firewallGroup)
    return firewallGroup, err
}

type FirewallRule struct {
    Action    string `json:"action"`
    Comment   string `json:"comment"`
    Dest      string `json:"dest"`
    Dport     string `json:"dport"`
    Enable    int    `json:"enable"`
    Iface     string `json:"iface"`
    IPVersion int    `json:"ipversion"`
    Log       string `json:"log"`
    Macro     string `json:"macro"`
    Pos       int    `json:"pos"`
    Proto     string `json:"proto"`
    Source    string `json:"source"`
    SPort     string `json:"sport"`
    Type      string `json:"type"`
}

func (c *Client) GetFirewallGroupRule(group string, pos int) (FirewallRule, error) {
    firewallRule := FirewallRule{}
    err := c.get("/cluster/firewall/groups/"+group+"/"+strconv.Itoa(pos), nil, &firewallRule)
    return firewallRule, err
}

type IPSet struct {
    Comment string `json:"comment"`
    Digest  string `json:"digest"`
    Name    string `json:"name"`
    CIDR    string `json:"cidr"`
    NoMatch bool   `json:"nomatch"`
}

func (c *Client) GetIPSets() ([]IPSet, error) {
    ipsets := []IPSet{}
    err := c.get("/cluster/firewall/ipset", nil, &ipsets)
    return ipsets, err
}

func (c *Client) GetIPSet(name string) (IPSet, error) {
    ipset := IPSet{}
    err := c.get("/cluster/firewall/ipset/"+name, nil, &ipset)
    return ipset, err
}

func (c *Client) GetIPSetSettings(cidr, name string) (string, error) {
    var settings string
    err := c.get("/cluster/firewall/ipset/"+name+"/"+cidr, nil, &settings)
    return settings, err
}

func (c *Client) GetFirewallRules() ([]FirewallRule, error) {
    firewallRules := []FirewallRule{}
    err := c.get("/cluster/firewall/rules", nil, &firewallRules)
    return firewallRules, err
}

func (c *Client) GetFirewallRule(pos int) (FirewallRule, error) {
    firewallRule := FirewallRule{}
    err := c.get("/cluster/firewall/rules/"+strconv.Itoa(pos), nil, &firewallRule)
    return firewallRule, err
}

type FirewallMacro struct {
    Description string `json:"descr"`
    Macro       string `json:"macro"`
}

func (c *Client) GetFirewallMacros() ([]FirewallMacro, error) {
    firewallMacros := []FirewallMacro{}
    err := c.get("/cluster/firewall/macros", nil, &firewallMacros)
    return firewallMacros, err
}

type FirewallOptions struct {
    EbTables     bool `json:"ebtables"`
    Enable       int  `json:"enable"`
    LogRatelimit struct {
        Description string `json:"description"`
        Format      struct {
            Burst  int    `json:"burst"`
            Enable bool   `json:"enable"`
            Rate   string `json:"rate"`
        } `json:"format"`
    } `json:"log_ratelimit"`
    PolicyIn  string `json:"policy_in"`
    PolicyOut string `json:"policy_out"`
}

func (c *Client) GetFirewallOptions() (FirewallOptions, error) {
    firewallOptions := FirewallOptions{}
    err := c.get("/cluster/firewall/options", nil, &firewallOptions)
    return firewallOptions, err
}

type FirewallReference struct {
    Comment string `json:"comment"`
    Name    string `json:"name"`
    Type    string `json:"type"`
}

func (c *Client) GetFirewallReferences() ([]FirewallReference, error) {
    firewallReferences := []FirewallReference{}
    err := c.get("/cluster/firewall/refs", nil, &firewallReferences)
    return firewallReferences, err
}