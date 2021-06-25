package client_id

import (
	context "context"

	github_com_go_courier_courier "github.com/go-courier/courier"
	github_com_go_courier_metax "github.com/go-courier/metax"
)

type ClusterStatus struct {
}

func (ClusterStatus) Path() string {
	return "/id/status/cluster"
}

func (ClusterStatus) Method() string {
	return "GET"
}

// @StatusErr[NotInK8S][500999002][Not in K8s]
func (req *ClusterStatus) Do(ctx context.Context, c github_com_go_courier_courier.Client, metas ...github_com_go_courier_courier.Metadata) github_com_go_courier_courier.Result {

	ctx = github_com_go_courier_metax.ContextWith(ctx, "operationID", "id.ClusterStatus")
	return c.Do(ctx, req, metas...)

}

func (req *ClusterStatus) InvokeContext(ctx context.Context, c github_com_go_courier_courier.Client, metas ...github_com_go_courier_courier.Metadata) (*map[string]uint32, github_com_go_courier_courier.Metadata, error) {
	resp := new(map[string]uint32)

	meta, err := req.Do(ctx, c, metas...).Into(resp)

	return resp, meta, err
}

func (req *ClusterStatus) Invoke(c github_com_go_courier_courier.Client, metas ...github_com_go_courier_courier.Metadata) (*map[string]uint32, github_com_go_courier_courier.Metadata, error) {
	return req.InvokeContext(context.Background(), c, metas...)
}

type CurrentStatus struct {
}

func (CurrentStatus) Path() string {
	return "/id/status"
}

func (CurrentStatus) Method() string {
	return "GET"
}

func (req *CurrentStatus) Do(ctx context.Context, c github_com_go_courier_courier.Client, metas ...github_com_go_courier_courier.Metadata) github_com_go_courier_courier.Result {

	ctx = github_com_go_courier_metax.ContextWith(ctx, "operationID", "id.CurrentStatus")
	return c.Do(ctx, req, metas...)

}

func (req *CurrentStatus) InvokeContext(ctx context.Context, c github_com_go_courier_courier.Client, metas ...github_com_go_courier_courier.Metadata) (*GithubComGoCourierHttptransportHttpxPlain, github_com_go_courier_courier.Metadata, error) {
	resp := new(GithubComGoCourierHttptransportHttpxPlain)

	meta, err := req.Do(ctx, c, metas...).Into(resp)

	return resp, meta, err
}

func (req *CurrentStatus) Invoke(c github_com_go_courier_courier.Client, metas ...github_com_go_courier_courier.Metadata) (*GithubComGoCourierHttptransportHttpxPlain, github_com_go_courier_courier.Metadata, error) {
	return req.InvokeContext(context.Background(), c, metas...)
}

type GenerateID struct {
}

func (GenerateID) Path() string {
	return "/id/v0/id"
}

func (GenerateID) Method() string {
	return "GET"
}

func (req *GenerateID) Do(ctx context.Context, c github_com_go_courier_courier.Client, metas ...github_com_go_courier_courier.Metadata) github_com_go_courier_courier.Result {

	ctx = github_com_go_courier_metax.ContextWith(ctx, "operationID", "id.GenerateID")
	return c.Do(ctx, req, metas...)

}

func (req *GenerateID) InvokeContext(ctx context.Context, c github_com_go_courier_courier.Client, metas ...github_com_go_courier_courier.Metadata) (*UniqueID, github_com_go_courier_courier.Metadata, error) {
	resp := new(UniqueID)

	meta, err := req.Do(ctx, c, metas...).Into(resp)

	return resp, meta, err
}

func (req *GenerateID) Invoke(c github_com_go_courier_courier.Client, metas ...github_com_go_courier_courier.Metadata) (*UniqueID, github_com_go_courier_courier.Metadata, error) {
	return req.InvokeContext(context.Background(), c, metas...)
}

type GenerateIDs struct {
	N int32 `in:"query" default:"1" name:"n,omitempty"`
}

func (GenerateIDs) Path() string {
	return "/id/v0/id/batch"
}

func (GenerateIDs) Method() string {
	return "GET"
}

func (req *GenerateIDs) Do(ctx context.Context, c github_com_go_courier_courier.Client, metas ...github_com_go_courier_courier.Metadata) github_com_go_courier_courier.Result {

	ctx = github_com_go_courier_metax.ContextWith(ctx, "operationID", "id.GenerateIDs")
	return c.Do(ctx, req, metas...)

}

func (req *GenerateIDs) InvokeContext(ctx context.Context, c github_com_go_courier_courier.Client, metas ...github_com_go_courier_courier.Metadata) (*IDList, github_com_go_courier_courier.Metadata, error) {
	resp := new(IDList)

	meta, err := req.Do(ctx, c, metas...).Into(resp)

	return resp, meta, err
}

func (req *GenerateIDs) Invoke(c github_com_go_courier_courier.Client, metas ...github_com_go_courier_courier.Metadata) (*IDList, github_com_go_courier_courier.Metadata, error) {
	return req.InvokeContext(context.Background(), c, metas...)
}
