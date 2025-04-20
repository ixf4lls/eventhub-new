package repository

type UserRepository interface {
	FindByUsername(username string) (*UserModel, error)
	Create(user *UserModel) error
	IsUsernameTaken(username string) (bool, error)
}
