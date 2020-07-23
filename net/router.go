package net

import "zinx/iface"

type BaseRouter struct {

}

func (b BaseRouter) PreHandle(request iface.IRequest) {

}

func (b BaseRouter) Handle(request iface.IRequest) {

}

func (b BaseRouter) PostHandle(request iface.IRequest) {

}

var _ iface.IRouter = new(BaseRouter)