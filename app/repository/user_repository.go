package repository

type Repository interface {
	FindByID(ID int) (User, error)
	FindByUsername(username string) (User, error)
}
