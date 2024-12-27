package servx

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineora-loc-go/network/restx"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/services/trigger"
)

type TriggerService struct {
}

func (t *TriggerService) Trigger(req *io.Request[trigger.Trigger]) *io.Response[trigger.Result] {
	data := req.Data
	for k, account := range data.Accounts {
		if account.Payer && (!req.HasPayer(account.Address)) {
			return io.FailResponse[trigger.Result](req.ID,
				errors.Verify("require payer: "+k+":"+account.Address.String()))
		}
		if account.Signer && (!req.HasSigner(account.Address)) {
			return io.FailResponse[trigger.Result](req.ID,
				errors.Verify("require signer: "+k+":"+account.Address.String()))
		}
	}
	return restx.Rest[trigger.Trigger, trigger.Result]("/triggers/trigger", req)
}
