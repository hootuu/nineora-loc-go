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
	if !req.HasSigner(req.Data.Authority) {
		return io.FailResponse[token.CreateResult](req.ID, errors.Verify("require signer for authority"))
	}
	if !req.HasSigner(req.Data.Address) {
		return io.FailResponse[token.CreateResult](req.ID, errors.Verify("require signer for address"))
	}
	return restx.Rest[token.Create, token.CreateResult]("/tokens/create", req)
}

func (t *TokenService) Mint(req *io.Request[token.Mint]) *io.Response[token.MintResult] {
	if !req.HasSigner(req.Data.TokenAuthority) {
		return io.FailResponse[token.MintResult](req.ID, errors.Verify("require signer for token authority"))
	}
	return restx.Rest[token.Mint, token.MintResult]("/tokens/mint", req)
}

func (t *TokenService) LoadAccountByAuthority(req *io.Request[token.LoadByAuthority]) *io.Response[token.LoadResult] {
	return restx.Rest[token.LoadByAuthority, token.LoadResult]("/tokens/load/by/authority", req)
}
