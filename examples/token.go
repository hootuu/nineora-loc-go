package examples

import (
	"fmt"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/hootuu/gelato/sys"
	"github.com/hootuu/nineora-loc-go/nineora"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/token"
	"time"
)

func TokenCreate() (*token.CreateResult, *errors.Error) {
	networkAddr, err := NetworkCreate()
	if err != nil {
		return nil, err
	}
	auth := GetKey("token_test_auth")
	wallet := GetKey("token_test_wallet")
	req := io.NewRequest[token.Create](&token.Create{
		Link:      domains.NewLink(fmt.Sprintf("token2_%d", time.Now().Unix())),
		Authority: auth.Address(),
		Network:   networkAddr.Address,
		Address:   wallet.Address(),
		Symbol:    domains.TokenSymbol(fmt.Sprintf("TK%d", time.Now().Unix())),
		Decimals:  0,
		Ctrl: domains.MustNewCtrl().
			MustSet(domains.TokenCtrlNonDivisible, true),
		Tag: domains.NewTag("SPACE"),
		Meta: domains.MustNewMeta().
			MustSet(domains.MetaName, "TOKEN").
			MustSet(domains.MetaUri, "http://xx.xx/xx.json").
			MustSet("address", "zhejian..."),
	})
	req.AddPayer(wallet.Address()).AddSigner(auth.Address())

	err = req.Sign(auth, wallet)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp := nineora.Nineora().Token().Create(req)
	//fmt.Println(resp.Json())
	fmt.Println("create token success =====>>>>>>>>><<<<<<<<<<", resp.Success)
	if !resp.Success {
		fmt.Println(resp.Error)
		return nil, resp.Error
	}
	fmt.Println("create token success =====>>>>>>>>>")
	return resp.Data, nil
}

func TokenMint() (*token.MintResult, *errors.Error) {
	newTokenMintResult, err := TokenCreate()
	if err != nil {
		return nil, err
	}

	owner, err := GetIdentity("token_test_id")
	if err != nil {
		return nil, err
	}

	auth := GetKey("token_test_auth")
	wallet := GetKey("token_test_wallet")
	for i := 0; i < 3; i++ {
		req := io.NewRequest[token.Mint](&token.Mint{
			Token:          newTokenMintResult.Address,
			Receive:        owner.Address(),
			Amount:         100,
			Memo:           domains.NewMemo().MustSet("order.id", fmt.Sprintf("ORD_%d", time.Now().Unix())),
			TokenAuthority: auth.Public.Address(),
		})
		req.AddPayer(wallet.Address()).AddSigner(auth.Address())

		err = req.Sign(auth, wallet)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		resp := nineora.Nineora().Token().Mint(req)
		fmt.Println(resp.Json())
		if !resp.Success {
			return nil, err
		}
		//return resp.Data, nil
	}

	//TEST ACC FIND
	req := io.NewRequest[token.AccLoadByAuth](&token.AccLoadByAuth{
		Authority: owner.Address(),
		Networks:  []domains.NetworkAddr{},
		Page:      *pagination.PageALL(),
	})
	req.AddPayer(wallet.Address()).AddSigner(auth.Address())
	err = req.Sign(auth, wallet)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp := nineora.Nineora().Token().AccLoadByAuth(req)
	if resp.Error != nil {
		fmt.Println(resp.Error)
		return nil, resp.Error
	}
	fmt.Println(resp.Json())
	return nil, nil

}

func doMint(mint keys.Address, receiver keys.Address, amount uint64) *errors.Error {
	auth := GetKey("token_test_auth")
	wallet := GetKey("token_test_wallet")
	req := io.NewRequest[token.Mint](&token.Mint{
		Token:          mint,
		Receive:        receiver,
		Amount:         amount,
		Memo:           domains.NewMemo().MustSet("order.id", fmt.Sprintf("ORD_%d", time.Now().Unix())),
		TokenAuthority: auth.Public.Address(),
	})
	req.AddPayer(wallet.Address()).AddSigner(auth.Address())

	err := req.Sign(auth, wallet)
	if err != nil {
		sys.Error(err)
		return err
	}
	resp := nineora.Nineora().Token().Mint(req)
	sys.Info(resp.Json())
	if !resp.Success {
		return err
	}

	return nil
}

func TokenAccountCreate() (*token.AccountCreateResult, *errors.Error) {
	tcr, err := TokenCreate()
	if err != nil {
		sys.Error("token create err:", err.Error())
		return nil, err
	}
	return doTokenAccountCreate(tcr.Address, "fee_loc")

}

func doTokenAccountCreate(mint keys.Address, linkStr string) (*token.AccountCreateResult, *errors.Error) {
	auth := GetKey("token_test_auth")
	wallet := GetKey("token_test_wallet")
	req := io.NewRequest(&token.AccountCreate{
		Link:      domains.GetTokenAccountLink(linkStr, mint, wallet.Address()),
		Authority: wallet.Address(),
		Mint:      mint,
		Ctrl:      nil,
		Tag:       nil,
		Meta:      nil,
	}).AddPayer(wallet.Address()).AddSigner(auth.Address())
	err := req.Sign(auth, wallet)
	if err != nil {
		sys.Error("token account create err-sign err:", err.Error())
		return nil, err
	}
	resp := nineora.Nineora().Token().AccCreate(req)
	if resp.Error != nil {
		sys.Error("token account create err-serv err:", resp.Error.Error())
		return nil, resp.Error
	}
	sys.Success("create token account success =====>>>>>>>>>>>>:", resp.Data.Address)
	return resp.Data, nil

}

func doGetTokenAccount(mint keys.Address, auth keys.Address, linkStr string) (*domains.TokenAccount, *errors.Error) {
	req := io.NewRequest[token.AccLoadByLink](&token.AccLoadByLink{
		Link:      linkStr,
		Mint:      mint,
		Authority: auth,
	})
	uniPayer := GetKey("uni_payer")
	req.AddPayer(uniPayer.Address())
	_ = req.Sign(uniPayer)
	resp := nineora.Nineora().Token().AccLoadByLink(req)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp.Data.One(), nil
}

func TokenTransfer() *errors.Error {
	tcr, err := TokenCreate()
	if err != nil {
		sys.Error("token create err:", err.Error())
		return err
	}
	wallet := GetKey("token_test_wallet")
	err = doMint(tcr.Address, wallet.Address(), 1000000)
	if err != nil {
		sys.Error("token mint err:", err.Error())
		return err
	}
	fromAtaAcc, err := doGetTokenAccount(tcr.Address, wallet.Address(), domains.AtaLinkTpl)
	if err != nil {
		sys.Error("doGetTokenAccount err:", err.Error())
		return err
	}
	toAccResp, err := doTokenAccountCreate(tcr.Address, "transfer_t")
	if err != nil {
		sys.Error("to token account create err:", err.Error())
		return err
	}
	req := io.NewRequest(&token.Transfer{
		FromAddress: fromAtaAcc.Address,
		ToAddress:   toAccResp.Address,
		Authority:   wallet.Address(),
		Amount:      899,
		Memo:        domains.NewMemo().MustSet(domains.MemoMemo, fmt.Sprintf("ORD_%d", time.Now().Unix())),
	}).AddPayer(wallet.Address())
	err = req.Sign(wallet)
	if err != nil {
		sys.Error("token transfer err-sign err:", err.Error())
		return err
	}
	resp := nineora.Nineora().Token().Transfer(req)
	if resp.Error != nil {
		sys.Error("token transfer err-serv err:", resp.Error.Error())
		return resp.Error
	}
	sys.Success("token transfer success =====>>>>>>>>>>>>:", resp.Data)

	toAcc, err := doGetTokenAccount(tcr.Address, wallet.Address(), "transfer_t")
	if err != nil {
		sys.Error("doGetTokenAccount err:", err.Error())
		return err
	}
	sys.Info("to account:", toAcc.Balance)

	_ = doGetTx(toAcc.Mint, toAcc.Address)

	return nil
}

func doGetTx(mint keys.Address, auth keys.Address) *errors.Error {
	req := io.NewRequest[token.TxLoad](&token.TxLoad{
		Authority: auth,
		Mint:      mint,
	})

	wallet := GetKey("token_test_wallet")
	req.AddPayer(wallet.Address())
	_ = req.Sign(wallet)
	resp := nineora.Nineora().Token().TxLoad(req)
	if resp.Error != nil {
		sys.Error("token load err:", resp.Error.Error())
		return resp.Error
	}
	sys.Info("load tx success =====>>>>>>>>>>>>:", resp.Data)
	return nil
}
