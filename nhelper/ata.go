package nhelper

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/sys"
	"github.com/hootuu/nineora-loc-go/nineora"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/token"
)

func MustGetAtaAccount(mint keys.Address, auth keys.Address, linkStr string, payer *keys.Key) (*domains.TokenAccount, *errors.Error) {
	acc, err := GetTokenAccount(mint, auth, linkStr, payer)
	if err != nil {
		return nil, err
	}
	if acc == nil {
		_, err := CreateTokenAccount(mint, auth, linkStr, payer)
		if err != nil {
			return nil, err
		}
		acc, err = GetTokenAccount(mint, auth, linkStr, payer)
		if err != nil {
			return nil, err
		}
		if acc == nil {
			return nil, errors.System("system. error")
		}
	}
	return acc, nil
}

func GetTokenAccount(mint keys.Address, auth keys.Address, linkStr string, payer *keys.Key) (*domains.TokenAccount, *errors.Error) {
	req := io.NewRequest[token.AccLoadByLink](&token.AccLoadByLink{
		Link:      linkStr,
		Mint:      mint,
		Authority: auth,
	})
	req.AddPayer(payer.Address())
	_ = req.Sign(payer)
	resp := nineora.Nineora().Token().AccLoadByLink(req)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp.Data.One(), nil
}

func CreateTokenAccount(mint keys.Address, auth keys.Address, linkStr string, payer *keys.Key) (keys.Address, *errors.Error) {
	req := io.NewRequest(&token.AccountCreate{
		Link:      domains.GetTokenAccountLink(linkStr, mint, auth),
		Authority: auth,
		Mint:      mint,
		Ctrl:      nil,
		Tag:       nil,
		Meta:      nil,
	}).AddPayer(payer.Address())
	err := req.Sign(payer)
	if err != nil {
		sys.Error("token account create err-sign err:", err.Error())
		return "", err
	}
	resp := nineora.Nineora().Token().AccCreate(req)
	if resp.Error != nil {
		sys.Error("token account create err-serv err:", resp.Error.Error())
		return "", resp.Error
	}
	sys.Success("create token account success =====>>>>>>>>>>>>:", resp.Data.Address)
	return resp.Data.Address, nil
}
