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

var gKeysDict map[string]*keys.Key
var gIdDict map[string]*keys.Key

func init() {
	gKeysDict = make(map[string]*keys.Key)
	gIdDict = make(map[string]*keys.Key)
}

func GetKey(keyName string) *keys.Key {
	key, ok := gKeysDict[keyName]
	if !ok {
		key, _ = keys.NewKey()
		gKeysDict[keyName] = key
	}
	return key
}

func GetIdentity(keyName string) (*keys.Key, *errors.Error) {
	idKey, ok := gIdDict[keyName]
	if !ok {
		idKey, _ = keys.NewKey()
		req := io.NewRequest[identity.Create](&identity.Create{
			Link:     domains.NewLink(fmt.Sprintf("AB_%d", time.Now().UnixMicro())),
			Password: domains.NewPassword("999909990"),
			Address:  idKey.Public.Address(),
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
			Address: idKey.Public.Address(),
			Payer:   true,
			Signer:  false,
		})
		err := req.Sign(idKey)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		resp := nineora.Nineora().Identity().Create(req)
		if !resp.Success {
			return nil, resp.Error
		}
		gIdDict[idKey.Public.Address().String()] = idKey
	}
	return idKey, nil
}
