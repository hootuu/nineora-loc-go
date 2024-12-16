package examples

import (
	"fmt"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/nineora"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/token"
	"time"
)

func TokenCreate() (*token.CreateResult, *errors.Error) {
	networkAddr, err := NetworkCreate()
	if err != nil {
		return nil, err
	}
	auth, _ := keys.NewKey()
	wallet, _ := keys.NewKey()
	req := io.NewRequest[token.Create](&token.Create{
		Authority: auth.Address(),
		Network:   networkAddr.Address,
		Address:   wallet.Address(),
		Symbol:    domains.TokenSymbol(fmt.Sprintf("TK%d", time.Now().Unix())),
		Decimals:  6,
		Ctrl:      nil,
		Tag:       nil,
		Meta: domains.MustNewMeta().
			MustSet(domains.MetaName, "TOKEN").
			MustSet(domains.MetaUri, "http://xx.xx/xx.json"),
	})
	req.AddPayer(wallet.Address()).AddSigner(auth.Address())

	err = req.Sign(auth, wallet)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp := nineora.Nineora().Token().Create(req)
	fmt.Println(resp.Json())
	if !resp.Success {
		return nil, err
	}
	return resp.Data, nil
}
