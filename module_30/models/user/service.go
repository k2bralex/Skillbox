package user

type Service interface {
	GetAll() []*User
	CreateUser(e *User) error
	GetById(id string) (*User, error)
	DeleteById(id string) error
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (r *service) GetAll() []*User {
	return r.storage.ReadAll()
}

func (r *service) CreateUser(e *User) error {
	return r.storage.Add(e, (*e).ID)
}

func (r *service) GetById(id string) (*User, error) {
	return r.storage.Read(id)
}

func (r *service) DeleteById(id string) error {
	return r.storage.Delete(id)
}
