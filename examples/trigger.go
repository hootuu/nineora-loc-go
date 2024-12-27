package examples

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/sys"
	"github.com/hootuu/nineora-loc-go/nineora"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/services/trigger"
)

func Trigger() *errors.Error {
	auth := GetKey("token_test_auth")
	wallet := GetKey("token_test_wallet")
	req := io.NewRequest[trigger.Trigger](&trigger.Trigger{
		Contract: "2UocKkSqAfsvkS1st2iCgXRJD2vYHoSoLn921RaZJWgT",
		Code:     "CUSTOM_ORD",
		Ctx:      domains.MustNewDict().MustSet("", ""),
		Accounts: io.NewAccounts().Put("a", io.Account{
			Address: "",
			Payer:   false,
			Signer:  false,
		}).Put("b", io.Account{
			Address: "",
			Payer:   false,
			Signer:  false,
		}),
		Memo: domains.NewMemo().MustSet(domains.MemoMemo, "xxx"),
	})
	req.AddPayer(wallet.Address()).AddSigner(auth.Address())

	err := req.Sign(auth, wallet)
	if err != nil {
		sys.Error(err)
		return err
	}
	resp := nineora.Nineora().Trigger().Trigger(req)
	if !resp.Success {
		sys.Error(resp.Error.Error())
		return resp.Error
	}
	return nil
}
