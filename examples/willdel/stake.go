package willdel

import (
	"fmt"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/examples"
	"github.com/hootuu/nineora-loc-go/nineora"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/stake"
	"time"
)

func StakeCreate() (*stake.CreateResult, *errors.Error) {
	networkAddr, err := examples.NetworkCreate()
	if err != nil {
		return nil, err
	}
	auth, _ := keys.NewKey()
	wallet, _ := keys.NewKey()
	req := io.NewRequest[stake.Create](&stake.Create{
		Link:      domains.NewLink(fmt.Sprintf("stake.%d", time.Now().Unix())),
		Authority: auth.Address(),
		Network:   networkAddr.Address,
		Address:   wallet.Address(),
		Symbol:    domains.StakeSymbol(fmt.Sprintf("TK%d", time.Now().Unix())),
		Total:     1,
		Ctrl:      nil,
		Tag:       nil,
		Meta: domains.MustNewMeta().
			MustSet(domains.MetaName, "STAKE").
			MustSet(domains.MetaUri, "http://xx.xx/xx.json"),
	})
	req.AddPayer(wallet.Address()).AddSigner(auth.Address())

	err = req.Sign(auth, wallet)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp := nineora.Nineora().Stake().Create(req)
	fmt.Println(resp.Json())
	if !resp.Success {
		return nil, err
	}
	return resp.Data, nil
}
