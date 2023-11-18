package svc

type ServiceContext struct {
}

//go:generate go run github.com/Bin-Huang/newc@v0.8.3
func NewServiceContext() *ServiceContext {
	return &ServiceContext{}
}

func (svc *ServiceContext) Run() error {
	return nil
}
