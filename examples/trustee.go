package examples

import (
	"fmt"
	"github.com/hootuu/nineora-loc-go/nineora"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/trustee"
	"time"
)

func TrusteeCreate() {
	req := io.NewRequest[trustee.Create](trustee.NewRandCreate(false))
	resp := nineora.Nineora().Trustee().Create(req)
	fmt.Println(resp.Data.Key.Public.ToBase58())
	fmt.Println(resp.Data.Key.Private.ToBase58())
	fmt.Println(resp.Data.Key.Private)

	trustee.NewCreate(true, fmt.Sprintf("NI%d", time.Now().Unix()), "xxxx")

	req = io.NewRequest[trustee.Create](trustee.NewRandCreate(true))
	k, _ := keys.NewKey()
	req.AddPayer(k.Address())
	_ = req.Sign(k)
	resp = nineora.Nineora().Trustee().Create(req)
	if resp.Success {
		fmt.Println(resp.Data.Key.Public.ToBase58())
		fmt.Println(resp.Data.Key.Private.ToBase58())
		fmt.Println(resp.Data.Key.Private)
	} else {
		fmt.Println(resp.Error)
	}
}
