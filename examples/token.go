package examples

import (
	"fmt"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/nineora"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/services/token"
	"time"
)

func TokenCreate() (*token.CreateResult, *errors.Error) {
	networkAddr, err := NetworkCreate()
	if err != nil {
		return nil, err
	}
	auth := GetKey("token_test_auth")
	wallet := GetKey("token_test_wallet")
	req := io.NewRequest[token.Create](&token.Create{
		Link:      domains.NewLink(fmt.Sprintf("token_%d", time.Now().Unix())),
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
	//fmt.Println(resp.Json())
	fmt.Println("create token success =====>>>>>>>>><<<<<<<<<<", resp.Success)
	if !resp.Success {
		fmt.Println(resp.Error)
		return nil, resp.Error
	}
	fmt.Println("create token success =====>>>>>>>>>")
	return resp.Data, nil
}

func TokenMint() (*token.MintResult, *errors.Error) {
	newTokenMintResult, err := TokenCreate()
	if err != nil {
		return nil, err
	}

	owner, err := IdentityCreate()
	if err != nil {
		return nil, err
	}

	auth := GetKey("token_test_auth")
	wallet := GetKey("token_test_wallet")
	for i := 0; i < 100; i++ {
		req := io.NewRequest[token.Mint](&token.Mint{
			Token:          newTokenMintResult.Address,
			Receive:        owner.Address,
			Amount:         100,
			Memo:           domains.NewMemo().MustSet("order.id", fmt.Sprintf("ORD_%d", time.Now().Unix())),
			TokenAuthority: auth.Public.Address(),
		})
		req.AddPayer(wallet.Address()).AddSigner(auth.Address())

		err = req.Sign(auth, wallet)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		resp := nineora.Nineora().Token().Mint(req)
		fmt.Println(resp.Json())
		if !resp.Success {
			return nil, err
		}
		//return resp.Data, nil
	}
	return nil, nil

}
