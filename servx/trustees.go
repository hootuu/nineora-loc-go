package servx

import (
	"github.com/hootuu/nineora-loc-go/network/restx"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/trustee"
)

type TrusteeService struct {
}

func (t *TrusteeService) Create(req *io.Request[trustee.Create]) *io.Response[trustee.CreateResult] {
	if !req.Data.Trustee {
		newKey, err := keys.NewKey()
		if err != nil {
			return io.FailResponse[trustee.CreateResult](req.ID, err)
		}
		return io.NewResponse[trustee.CreateResult](req.ID, &trustee.CreateResult{Key: keys.Key{
			Private: newKey.Private,
			Public:  newKey.Public,
		}})
	}
	return restx.Rest[trustee.Create, trustee.CreateResult]("/trustees/create", req)
}
