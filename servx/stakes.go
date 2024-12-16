package servx

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/network/restx"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/services/stake"
)

type StakeService struct {
}

func (s *StakeService) Create(req *io.Request[stake.Create]) *io.Response[stake.CreateResult] {
	data := req.Data
	if !req.HasSigner(data.Authority) {
		return io.FailResponse[stake.CreateResult](req.ID, errors.Verify("require signer for authority"))
	}
	if !req.HasSigner(data.Address) {
		return io.FailResponse[stake.CreateResult](req.ID, errors.Verify("require signer for address"))
	}
	return restx.Rest[stake.Create, stake.CreateResult]("/stakes/create", req)
}
