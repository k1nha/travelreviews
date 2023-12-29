package adapter

type JwtAdapter interface {
	CreateToken(email string, username string) (string, error)
}
