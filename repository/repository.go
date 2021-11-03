package repository

type Repository interface {
	Close()
	FindById(id string) (*UserModel, error)
	Find() ([]*UserModel, error)
	Create(user *UserModel) error
	Update(user *UserModel) error
	Delete(id string) error
}

type UserModel struct {
	ID    string
	Name  string
	Email string
	Phone string
}
