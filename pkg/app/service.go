package app

import "sync"

type Service struct {}

var (
	service *Service
	serviceOnce sync.Once
)

func NewService() *Service {
	return &Service{}
}

func GetAppService() *Service {
	serviceOnce.Do(func(){
		service = NewService()
	})
	return service
}

func (s *Service) SignUp() (bool, error) {
	return false, nil
}

func (s *Service) SignIn() (string, error) {
 	return "", nil
}

func (s *Service) SignOut() error {
	return nil
}

func (s *Service) IsAuthenticated() (bool, error) {
	return false, nil
}
