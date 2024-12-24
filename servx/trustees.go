package servx

import (
	"github.com/hootuu/nineora-loc-go/network/restx"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/trustee"
)

type TrusteeService struct {
}

func (t *TrusteeService) Exists(req *io.Request[trustee.Exists]) *io.Response[trustee.ExistsResult] {
	return restx.Rest[trustee.Exists, trustee.ExistsResult]("/trustees/exists", req)
}

func (t *TrusteeService) Create(req *io.Request[trustee.Create]) *io.Response[trustee.CreateResult] {
	if !req.Data.Trustee {
		newKey, err := keys.NewKey()
		if err != nil {
			return io.FailResponse[trustee.CreateResult](req.ID, err)
		}
		return io.NewResponse[trustee.CreateResult](req.ID, &trustee.CreateResult{Address: newKey.Address()})
	}
	return restx.Rest[trustee.Create, trustee.CreateResult]("/trustees/create", req)
}
