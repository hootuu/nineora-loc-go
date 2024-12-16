package examples

import (
	"fmt"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/nineora"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/asset"
	"time"
)

func AssetCreate() (*asset.CreateResult, *errors.Error) {
	networkAddr, err := NetworkCreate()
	if err != nil {
		return nil, err
	}
	auth, _ := keys.NewKey()
	wallet, _ := keys.NewKey()
	req := io.NewRequest[asset.Create](&asset.Create{
		Link:      domains.NewLink(fmt.Sprintf("LK_%d", time.Now().Unix())),
		Authority: auth.Address(),
		Network:   networkAddr.Address,
		Address:   wallet.Address(),
		Symbol:    domains.AssetSymbol(fmt.Sprintf("TK%d", time.Now().Unix())),
		Ctrl:      nil,
		Tag:       nil,
		Meta: domains.MustNewMeta().
			MustSet(domains.MetaName, "ASSET").
			MustSet(domains.MetaUri, "http://xx.xx/xx.json").
			MustSet(domains.MetaDescription, "xxxxxx"),
	})
	req.AddPayer(wallet.Address()).AddSigner(auth.Address())

	err = req.Sign(auth, wallet)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp := nineora.Nineora().Asset().Create(req)
	fmt.Println(resp.Json())
	if !resp.Success {
		return nil, err
	}
	return resp.Data, nil
}
