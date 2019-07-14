package api

type Node struct {
	CPUUtil        float64 `json:"cpu,omitempty"`
	SupportLevel   string  `json:"level,omitempty"`
	MaxCPU         uint    `json:"maxcpu,omitempty"`
	MaxMem         uint    `json:"maxmem,omitempty"`
	MaxDisk        uint    `json:"maxdisk,omitempty"`
	MemUsed        uint    `json:"mem,omitempty"`
	DiskSpaceUsed  uint    `json:"disk,omitempty"`
	Node           string  `json:"node"`
	SSLFingerprint string  `json:"ssl_fingerprint,omitempty"`
	Status         string  `json:"status"`
	Uptime         uint    `json:"uptime,omitempty"`
}

func (c *Client) GetNodes() ([]Node, error) {
	var nodes []Node
	err := c.get("/nodes", nil, &nodes)
	return nodes, err
}
