package auth

type IdentityRow struct {
	Id int
	Identity string
	Credential string
}

type Identity struct {
	Id int
	Name string
}

type Request struct {
	Identity string
	Credential string
}