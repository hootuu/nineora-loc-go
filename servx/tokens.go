package servx

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/network/restx"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/services/token"
)

type TokenService struct {
}

func (t *TokenService) TxLoad(req *io.Request[token.TxLoad]) *io.Response[token.TxLoadResult] {
	return restx.Rest[token.TxLoad, token.TxLoadResult]("/tokens/tx", req)
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

func (t *TokenService) AccLoadByAuth(req *io.Request[token.AccLoadByAuth]) *io.Response[token.AccLoadResult] {
	return restx.Rest[token.AccLoadByAuth, token.AccLoadResult]("/tokens/acc/load/by/auth", req)
}

func (t *TokenService) AccLoadByLink(req *io.Request[token.AccLoadByLink]) *io.Response[token.AccLoadResult] {
	return restx.Rest[token.AccLoadByLink, token.AccLoadResult]("/tokens/acc/load/by/link", req)
}

func (t *TokenService) AccCreate(req *io.Request[token.AccountCreate]) *io.Response[token.AccountCreateResult] {
	return restx.Rest[token.AccountCreate, token.AccountCreateResult]("/tokens/acc/create", req)
}

func (t *TokenService) Transfer(req *io.Request[token.Transfer]) *io.Response[token.TransferResult] {
	if !req.HasSigner(req.Data.Authority) {
		return io.FailResponse[token.TransferResult](req.ID, errors.Verify("require signer for authority"))
	}
	return restx.Rest[token.Transfer, token.TransferResult]("/tokens/transfer", req)
}
