package service

type IAppService interface {
	GetHello() string
}

type AppService struct{}

func NewAppService() AppService {
	return AppService{}
}

func (AppService) GetHello() string {
	return "hello, world!"
}
