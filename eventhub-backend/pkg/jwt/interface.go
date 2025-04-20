package customJwt

type Manager interface {
	GenerateAccessToken(userID uint) (string, error)
	GenerateRefreshToken(userID uint) (string, error)
	ParseToken(tokenStr string) (uint, error)
}
