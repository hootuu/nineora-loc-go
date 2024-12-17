package servx

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/network/restx"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/services/identity"
)

type IdentityService struct {
}

func (i *IdentityService) Get(req *io.Request[identity.Get]) *io.Response[identity.GetResult] {
	return restx.Rest[identity.Get, identity.GetResult]("/identities/get", req)
}

func (i *IdentityService) GetByLink(req *io.Request[identity.GetByLink]) *io.Response[identity.GetResult] {
	return restx.Rest[identity.GetByLink, identity.GetResult]("/identities/get/by/link", req)
}

func (i *IdentityService) GetByNineoraID(req *io.Request[identity.GetByNineoraID]) *io.Response[identity.GetResult] {
	return restx.Rest[identity.GetByNineoraID, identity.GetResult]("/identities/get/by/nid", req)
}

func (i *IdentityService) Create(req *io.Request[identity.Create]) *io.Response[identity.CreateResult] {
	if !req.HasSigner(req.Data.Address) {
		return io.FailResponse[identity.CreateResult](req.ID, errors.Verify("the address is invalid or no signer for it"))
	}
	return restx.Rest[identity.Create, identity.CreateResult]("/identities/create", req)
}
