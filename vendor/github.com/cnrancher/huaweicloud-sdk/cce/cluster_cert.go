package cce

import (
	"context"
	"errors"
	"net/http"

	"github.com/cnrancher/huaweicloud-sdk/common"
)

func (c *Client) GetClusterCert(ctx context.Context, clusterid string) (*common.ClusterCert, error) {
	if clusterid == "" {
		return nil, errors.New("cluster id is required")
	}
	rtn := common.ClusterCert{}
	_, err := c.DoRequest(
		ctx,
		http.MethodGet,
		c.GetURL("clusters", clusterid, "clustercert"),
		nil,
		&rtn,
	)
	if err != nil {
		return nil, err
	}
	return &rtn, nil
}
