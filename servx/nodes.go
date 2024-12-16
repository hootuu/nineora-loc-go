package servx

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/network/restx"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/services/node"
)

type NodeService struct {
}

func (n *NodeService) Create(req *io.Request[node.Create]) *io.Response[node.CreateResult] {
	data := req.Data
	if !req.HasSigner(data.Authority) {
		return io.FailResponse[node.CreateResult](req.ID, errors.Verify("require signer for authority"))
	}
	if !req.HasSigner(data.Address) {
		return io.FailResponse[node.CreateResult](req.ID, errors.Verify("require signer for address"))
	}
	return restx.Rest[node.Create, node.CreateResult]("/nodes/create", req)
}
