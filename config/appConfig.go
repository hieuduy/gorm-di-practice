package config

import (
	"MyGo/service"
	"MyGo/service/impl"
	"github.com/golobby/container/pkg/container"
)

var Container container.Container

func InitDIContainer() {
	Container = container.NewContainer()
}

func BindingDependencies() {
	//Container.Singleton(func() service.OrderService {
	//	return &impl.OrderServiceImpl{}
	//})

	Container.Singleton(func() service.UserService {
		return &impl.UserServiceImpl{}
	})
}