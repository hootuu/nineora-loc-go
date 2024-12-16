package servx

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/network/restx"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/services/asset"
)

type AssetService struct {
}

func (a *AssetService) Create(req *io.Request[asset.Create]) *io.Response[asset.CreateResult] {
	data := req.Data
	if !req.HasSigner(data.Authority) {
		return io.FailResponse[asset.CreateResult](req.ID, errors.Verify("require signer for authority"))
	}
	if !req.HasSigner(data.Address) {
		return io.FailResponse[asset.CreateResult](req.ID, errors.Verify("require signer for address"))
	}
	return restx.Rest[asset.Create, asset.CreateResult]("/assets/create", req)
}
