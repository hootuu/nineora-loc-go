package servx

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/network/restx"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/services/token"
)

type TokenService struct {
}

func (t *TokenService) Create(req *io.Request[token.Create]) *io.Response[token.CreateResult] {
	data := req.Data
	if !req.HasSigner(data.Authority) {
		return io.FailResponse[token.CreateResult](req.ID, errors.Verify("require signer for authority"))
	}
	if !req.HasSigner(data.Address) {
		return io.FailResponse[token.CreateResult](req.ID, errors.Verify("require signer for address"))
	}
	return restx.Rest[token.Create, token.CreateResult]("/tokens/create", req)
}
