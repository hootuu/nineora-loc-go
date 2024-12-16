package nineora

import (
	"github.com/hootuu/nineora-loc-go/servx"
	"github.com/hootuu/nineorai/services"
	"github.com/hootuu/nineorai/services/asset"
	"github.com/hootuu/nineorai/services/identity"
	"github.com/hootuu/nineorai/services/network"
	"github.com/hootuu/nineorai/services/node"
	"github.com/hootuu/nineorai/services/stake"
	"github.com/hootuu/nineorai/services/token"
	"github.com/hootuu/nineorai/services/trustee"
	"sync"
)

type nineora struct {
	trustee  trustee.Service
	identity identity.Service
	network  network.Service
	node     node.Service
	token    token.Service
	stake    stake.Service
	asset    asset.Service
}

var instance *nineora
var once sync.Once

func Nineora() services.Nineora {
	once.Do(func() {
		instance = &nineora{
			trustee:  &servx.TrusteeService{},
			identity: &servx.IdentityService{},
			network:  &servx.NetworkService{},
			node:     &servx.NodeService{},
			token:    &servx.TokenService{},
			stake:    &servx.StakeService{},
			asset:    &servx.AssetService{},
		}
	})
	return instance
}

func (nineora *nineora) Trustee() trustee.Service {
	return nineora.trustee
}

func (nineora *nineora) Identity() identity.Service {
	return nineora.identity
}

func (nineora *nineora) Network() network.Service {
	return nineora.network
}

func (nineora *nineora) Node() node.Service {
	return nineora.node
}

func (nineora *nineora) Token() token.Service {
	return nineora.token
}

func (nineora *nineora) Stake() stake.Service {
	return nineora.stake
}

func (nineora *nineora) Asset() asset.Service {
	return nineora.asset
}
