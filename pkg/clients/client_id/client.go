package client_id

import (
	context "context"

	github_com_go_courier_courier "github.com/go-courier/courier"
)

type ClientID interface {
	WithContext(context.Context) ClientID
	Context() context.Context
	ClusterStatus(metas ...github_com_go_courier_courier.Metadata) (*map[string]uint32, github_com_go_courier_courier.Metadata, error)
	CurrentStatus(metas ...github_com_go_courier_courier.Metadata) (*GithubComGoCourierHttptransportHttpxPlain, github_com_go_courier_courier.Metadata, error)
	GenerateID(metas ...github_com_go_courier_courier.Metadata) (*UniqueID, github_com_go_courier_courier.Metadata, error)
	GenerateIDs(req *GenerateIDs, metas ...github_com_go_courier_courier.Metadata) (*IDList, github_com_go_courier_courier.Metadata, error)
}

func NewClientID(c github_com_go_courier_courier.Client) *ClientIDStruct {
	return &(ClientIDStruct{
		Client: c,
	})
}

type ClientIDStruct struct {
	Client github_com_go_courier_courier.Client
	ctx    context.Context
}

func (c *ClientIDStruct) WithContext(ctx context.Context) ClientID {
	cc := new(ClientIDStruct)
	cc.Client = c.Client
	cc.ctx = ctx
	return cc
}

func (c *ClientIDStruct) Context() context.Context {
	if c.ctx != nil {
		return c.ctx
	}
	return context.Background()
}

func (c *ClientIDStruct) ClusterStatus(metas ...github_com_go_courier_courier.Metadata) (*map[string]uint32, github_com_go_courier_courier.Metadata, error) {
	return (&ClusterStatus{}).InvokeContext(c.Context(), c.Client, metas...)
}

func (c *ClientIDStruct) CurrentStatus(metas ...github_com_go_courier_courier.Metadata) (*GithubComGoCourierHttptransportHttpxPlain, github_com_go_courier_courier.Metadata, error) {
	return (&CurrentStatus{}).InvokeContext(c.Context(), c.Client, metas...)
}

func (c *ClientIDStruct) GenerateID(metas ...github_com_go_courier_courier.Metadata) (*UniqueID, github_com_go_courier_courier.Metadata, error) {
	return (&GenerateID{}).InvokeContext(c.Context(), c.Client, metas...)
}

func (c *ClientIDStruct) GenerateIDs(req *GenerateIDs, metas ...github_com_go_courier_courier.Metadata) (*IDList, github_com_go_courier_courier.Metadata, error) {
	return req.InvokeContext(c.Context(), c.Client, metas...)
}
