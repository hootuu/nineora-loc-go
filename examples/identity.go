package examples

import (
	"fmt"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/nineora"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/identity"
	"time"
)

func IdentityCreate() (*identity.CreateResult, *errors.Error) {
	userKey, _ := keys.NewKey()
	//should use trustee.Create
	req := io.NewRequest[identity.Create](&identity.Create{
		Link:     domains.NewLink(fmt.Sprintf("AB_%d", time.Now().UnixMicro())),
		Password: domains.NewPassword("999909990"),
		Address:  userKey.Public.Address(),
		Ctrl: domains.MustNewCtrl().
			MustSet(2, true).
			MustSet(3, true).
			MustSet(8, true),
		Tag: domains.NewTag("A", "B"),
		Meta: domains.MustNewMeta().
			MustSet("mobi", "123456789").
			MustSet("icon", "aa").
			MustSet("master", true),
	})
	_ = req.Accounts.AddAccount(io.Account{
		Address: userKey.Public.Address(),
		Payer:   true,
		Signer:  false,
	})
	err := req.Sign(userKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp := nineora.Nineora().Identity().Create(req)
	if !resp.Success {
		return nil, resp.Error
	}
	fmt.Println(resp.Data.NineoraID)
	getReq := io.NewRequest[identity.Get](&identity.Get{Address: resp.Data.Address})
	getReq.AddPayer(userKey.Public.Address())
	_ = getReq.Sign(userKey)
	getResp := nineora.Nineora().Identity().Get(getReq)
	if !getResp.Success {
		return nil, getResp.Error
	}
	fmt.Println(getResp.Json())
	return resp.Data, nil
}
