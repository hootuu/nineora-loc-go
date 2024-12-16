package servx

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/network/restx"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/services/network"
)

type NetworkService struct {
}

func (s *NetworkService) Create(req *io.Request[network.Create]) *io.Response[network.CreateResult] {
	data := req.Data
	if !req.HasSigner(data.Authority) {
		return io.FailResponse[network.CreateResult](req.ID, errors.Verify("require signer for authority"))
	}
	if !req.HasSigner(data.Address) {
		return io.FailResponse[network.CreateResult](req.ID, errors.Verify("require signer for address"))
	}
	return restx.Rest[network.Create, network.CreateResult]("/networks/create", req)
}
