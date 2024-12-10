package iasset

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/services/asset"
)

type Service struct {
}

func (s Service) SetCtrl(ctrl domains.Ctrl, ctx io.Ctx) io.Error {
	//TODO implement me
	panic("implement me")
}

func (s Service) AddTag(tag domains.Tag, ctx io.Ctx) io.Error {
	//TODO implement me
	panic("implement me")
}

func (s Service) RemoveTag(tag domains.Tag, ctx io.Ctx) io.Error {
	//TODO implement me
	panic("implement me")
}

func (s Service) SetMeta(dict domains.Dict, ctx io.Ctx) io.Error {
	//TODO implement me
	panic("implement me")
}

func (s Service) RemoveMeta(keys []string, ctx io.Ctx) io.Error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Create(req asset.Create, ctx asset.CreateCtx) (domains.AssetAddr, io.Error) {
	//TODO implement me
	panic("implement me")
}
