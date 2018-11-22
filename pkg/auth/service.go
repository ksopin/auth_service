package auth

type Service interface {
	Authenticate(r *Request) (*Identity, error)
}

