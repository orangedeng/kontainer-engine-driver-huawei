package cce

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cnrancher/huaweicloud-sdk/common"
	"github.com/sirupsen/logrus"
)

type Client struct {
	common.Client
}

func (c *Client) CreateCluster(ctx context.Context, cluster *common.ClusterInfo) (*common.ClusterInfo, error) {
	logrus.Info("Creating Cluster")
	var clusterResp common.ClusterInfo
	_, err := c.DoRequest(
		ctx,
		http.MethodPost,
		c.GetURL("clusters"),
		cluster,
		&clusterResp,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating cluster: %v", err)
	}

	return &clusterResp, nil
}

func (c *Client) UpdateCluster(ctx context.Context, id string, info *common.ClusterInfo) (*common.ClusterInfo, error) {
	if id == "" {
		return nil, errors.New("cluster id is required")
	}
	return nil, nil
}

func (c *Client) GetCluster(ctx context.Context, id string) (*common.ClusterInfo, error) {
	if id == "" {
		return nil, errors.New("cluster id is required")
	}
	rtn := common.ClusterInfo{}
	_, err := c.DoRequest(
		ctx,
		http.MethodGet,
		c.GetURL("clusters", id),
		nil,
		&rtn,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting %s cluster: %v", id, err)
	}
	return &rtn, nil
}

func (c *Client) GetClusters(ctx context.Context) (*common.ClusterListInfo, error) {
	rtn := common.ClusterListInfo{}
	_, err := c.DoRequest(ctx,
		http.MethodGet,
		c.GetURL("clusters"),
		nil,
		&rtn,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting clusters")
	}
	return &rtn, err
}

func (c *Client) DeleteCluster(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("cluster id is required")
	}
	logrus.Infof("Deleting Cluster %s", id)
	_, err := c.DoRequest(
		ctx,
		http.MethodDelete,
		c.GetURL("clusters", id),
		nil,
		nil,
	)
	if err != nil {
		return fmt.Errorf("error deleting cluster: %v", err)
	}

	return common.WaitForDeleteComplete(ctx, func(ictx context.Context) error {
		_, err := c.GetCluster(ctx, id)
		return err
	})
}
