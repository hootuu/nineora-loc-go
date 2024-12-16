package servx

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/network/restx"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/services/identity"
)

type IdentityService struct {
}

func (i *IdentityService) Create(req *io.Request[identity.Create]) *io.Response[identity.CreateResult] {
	if !req.HasSigner(req.Data.Address) {
		return io.FailResponse[identity.CreateResult](req.ID, errors.Verify("the address is invalid or no signer for it"))
	}
	return restx.Rest[identity.Create, identity.CreateResult]("/identities/create", req)
}
