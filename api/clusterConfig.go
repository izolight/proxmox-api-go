package api

func (c *Client) GetClusterNodes() ([]string, error) {
    var nodes []string
    err := c.get("/cluster/config/nodes", nil, &nodes)
    return nodes, err
}

type JoinInfo struct {
    ConfigDigest string `json:"config_digest"`
    NodeList     struct {
        Items []struct {
            Name        string `json:"name"`
            NodeId      int    `json:"nodeid"`
            PveAddr     string `json:"pve_addr"`
            PveFP       string `json:"pve_fp"`
            QuorumVotes int    `json:"quorum_votes"`
            Ring0Addr   string `json:"ring0_addr"`
        } `json:"items"`
    } `json:"nodelist"`
    PreferredNode string `json:"preferred_node"`
    Totem         string `json:"totem"`
}

func (c *Client) GetJoinInfo(node string) (JoinInfo, error) {
    joinInfo := JoinInfo{}
    err := c.get("/cluster/config/join", map[string]string{"node": node}, &joinInfo)
    return joinInfo, err
}

func (c *Client) GetTotem() (string, error) {
    var totem string
    err := c.get("/cluster/config/totem", nil, &totem)
    return totem, err
}