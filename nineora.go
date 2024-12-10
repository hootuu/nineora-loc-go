package nineora_loc_go

import (
	"github.com/hootuu/nineora-loc-go/iasset"
	"github.com/hootuu/nineora-loc-go/inode"
	"github.com/hootuu/nineorai/services"
	"github.com/hootuu/nineorai/services/asset"
	"github.com/hootuu/nineorai/services/identity"
	"github.com/hootuu/nineorai/services/node"
	"github.com/hootuu/nineorai/services/stake"
	"github.com/hootuu/nineorai/services/token"
	"github.com/hootuu/nineorai/services/vn"
	"sync"
)

type nineora struct {
	identity identity.Service
	network  vn.Service
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
			identity: nil,
			network:  nil,
			node:     &inode.Service{},
			token:    nil,
			stake:    nil,
			asset:    &iasset.Service{},
		}
	})
	return instance
}

func (nineora *nineora) Identity() identity.Service {
	//TODO implement me
	panic("implement me")
}

func (nineora *nineora) Network() vn.Service {
	return nineora.network
}

func (nineora *nineora) Node() node.Service {
	//TODO implement me
	panic("implement me")
}

func (nineora *nineora) Token() token.Service {
	//TODO implement me
	panic("implement me")
}

func (nineora *nineora) Stake() stake.Service {
	//TODO implement me
	panic("implement me")
}

func (nineora *nineora) Asset() asset.Service {
	//TODO implement me
	panic("implement me")
}
