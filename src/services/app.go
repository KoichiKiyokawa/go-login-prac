package services

type IAppService interface {
	GetHello() string
}

type AppService struct{}

func (AppService) GetHello() string {
	return "hello, world!"
}
