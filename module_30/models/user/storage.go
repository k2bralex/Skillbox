package user

type Storage interface {
	Create(user User) (string, error)
	Read(id string) (User, error)
	ReadAll() ([]User, error)
	Update(user User) error
	Delete(id string) error
}
