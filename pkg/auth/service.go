package auth

type Service interface {
	Authenticate(identity, credential string) (*Identity, error)
}

