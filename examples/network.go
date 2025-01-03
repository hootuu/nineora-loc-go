package examples

import (
	"fmt"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/nineora"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/network"
	"time"
)

func NetworkCreate() (*network.CreateResult, *errors.Error) {
	auth, _ := keys.NewKey()
	wallet, _ := keys.NewKey()
	req := io.NewRequest[network.Create](&network.Create{
		Link:      domains.NewLink(fmt.Sprintf("VN_%d", time.Now().Unix())),
		Authority: auth.Address(),
		Address:   wallet.Address(),
		Symbol:    domains.NetworkSymbol(fmt.Sprintf("VN%d", time.Now().Unix())),
		Ctrl:      nil,
		Tag:       nil,
		Meta: domains.MustNewMeta().
			MustSet(domains.MetaName, "NET aWORK").
			MustSet(domains.MetaUri, "http://xx.xx/xx.json"),
	})
	req.AddPayer(wallet.Address()).AddSigner(auth.Address())

	err := req.Sign(auth, wallet)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp := nineora.Nineora().Network().Create(req)
	fmt.Println(resp.Json())
	if !resp.Success {
		return nil, resp.Error
	}
	return resp.Data, nil
}
