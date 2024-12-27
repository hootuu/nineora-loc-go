package main

import (
	"fmt"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

func main() {
	//examples.IdentityCreate()
	//examples.TrusteeCreate()
	//examples.NetworkCreate()
	//examples.TokenCreate()
	//examples.TokenMint()
	//examples.NodeCreate()
	//examples.Trigger()
	//examples.TokenAccountCreate()
	//examples.TokenTransfer()
	mint := "AjFQUXezDmjNzK8gXw68a6wEDNoUDAg1VJotU44Nb36L"
	auth := "Dzwbdk39xH11RioJoAivGdFvJzYog9BXKgsTu1buihXZ"
	addr := domains.GetAtaLink(keys.Address(mint), keys.Address(auth))
	fmt.Println(addr)
	addrBy := domains.GetTokenAccountLink(domains.AtaLinkTpl, keys.Address(mint), keys.Address(auth))
	fmt.Println(addrBy)
}
