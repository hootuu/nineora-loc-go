package examples

import (
	"fmt"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/nineora"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/node"
	"time"
)

func NodeCreate() (*node.CreateResult, *errors.Error) {
	networkAddr, err := NetworkCreate()
	if err != nil {
		return nil, err
	}
	auth, _ := keys.NewKey()
	wallet, _ := keys.NewKey()
	req := io.NewRequest[node.Create](&node.Create{
		Link:      domains.NewLink(fmt.Sprintf("node_%d", time.Now().Unix())),
		Authority: auth.Address(),
		Network:   networkAddr.Address,
		Address:   auth.Address(),
		Ctrl:      nil,
		Tag:       nil,
		Meta:      nil,
	})
	req.AddPayer(wallet.Address()).AddSigner(auth.Address())

	err = req.Sign(auth, wallet)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp := nineora.Nineora().Node().Create(req)
	fmt.Println(resp.Json())
	if !resp.Success {
		return nil, err
	}
	return resp.Data, nil
}
