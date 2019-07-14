package api

func (c *Client) GetStorages() ([]map[string]string, error) {
	var storages []map[string]string
	err := c.get("/storage", nil, &storages)
	return storages, err
}

type DirStorage struct {
	Type     string `json:"type"`
	Digest   string `json:"digest"`
	Path     string `json:"path"`
	MaxFiles uint   `json:"maxfiles"`
	Storage  string `json:"storage"`
	Content  string `json:"content"`
}

func (c *Client) GetDirStorages() ([]DirStorage, error) {
	var dirStorage []DirStorage
	err := c.get("/storage", map[string]string{"type": "dir"}, &dirStorage)
	return dirStorage, err
}

type GlusterFSStorage struct {
	Volume  string `json:"volume"`
	Storage string `json:"storage"`
	Content string `json:"content"`
	Digest  string `json:"digest"`
	Shared  uint   `json:"shared"`
	Server  string `json:"server"`
	Type    string `json:"type"`
	Path    string `json:"path"`
}

func (c *Client) GetGlusterFSStorages() ([]GlusterFSStorage, error) {
	var glusterFSStorage []GlusterFSStorage
	err := c.get("/storage", map[string]string{"type": "dir"}, &glusterFSStorage)
	return glusterFSStorage, err
}

func (c *Client) GetStorage(storage string) (map[string]string, error) {
    var sto map[string]string
    err := c.get("/storage/" + storage, nil, &sto)
    return sto, err
}