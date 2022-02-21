package auth

type Auth interface {
	LoginEmail(email, password string) (string, error)
	GetPasswordByEmail(email string) (string, error)
	GetIdByEmail(email string) (int, error)
	GetIdRole(Email string) (int, error)
}
