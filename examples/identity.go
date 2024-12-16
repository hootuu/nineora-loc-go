package examples

import (
	"fmt"
	"github.com/hootuu/nineora-loc-go/nineora"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/identity"
	"time"
)

func IdentityCreate() {
	userKey, _ := keys.NewKey()
	req := io.NewRequest[identity.Create](&identity.Create{
		CustomID: fmt.Sprintf("AB_%d", time.Now().UnixMicro()),
		Password: []byte("999909990"),
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
		return
	}
	resp := nineora.Nineora().Identity().Create(req)
	if resp.Success {
		fmt.Println(resp.Data.NineoraID)
	} else {
		fmt.Println(resp.Error.Message)
	}
}
