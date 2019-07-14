package api

func (c *Client) GetReplicationJobs() ([]interface{}, error) {
	var replicationJobs []interface{}
	err := c.get("/cluster/replication", nil, &replicationJobs)
	return replicationJobs, err
}

func (c *Client) GetReplicationJob(id string) (interface{}, error) {
	var replicationJob interface{}
	err := c.get("/cluster/replication/"+id, nil, &replicationJob)
	return replicationJob, err
}
