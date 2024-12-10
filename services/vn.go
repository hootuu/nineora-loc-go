package services

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"github.com/hootuu/nineorai/services/vn"
)

type Service struct {
}

func (s Service) Create(req vn.Create) (keys.Address, io.Error) {
	//TODO implement me
	panic("implement me")
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
