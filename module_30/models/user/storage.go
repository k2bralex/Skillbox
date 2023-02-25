package user

import "fmt"

type Storage interface {
	Add(e *User, s string) error
	Read(id string) (*User, error)
	ReadAll() []*User
	Update(e *User) error
	Delete(id string) error
}

type WorkStorage struct {
	repo map[string]*User
}

func NewWorkStorage() *WorkStorage {
	return &WorkStorage{
		repo: map[string]*User{},
	}
}

func (ws *WorkStorage) Add(e *User, s string) error {
	if _, ok := ws.repo[s]; !ok {
		ws.repo[s] = e
		return nil
	}
	return fmt.Errorf("entity already exist")
}

func (ws *WorkStorage) Read(id string) (*User, error) {
	if ws.contains(id) {
		return ws.repo[id], nil
	}
	return nil, fmt.Errorf("entity not exist")
}

func (ws *WorkStorage) ReadAll() []*User {
	values := make([]*User, 0, len(ws.repo))
	for _, val := range ws.repo {
		values = append(values, val)
	}
	return values
}

func (ws *WorkStorage) Update(e *User) error {
	if !ws.contains((*e).ID) {
		return fmt.Errorf("entity not exist, udate faild")
	}

	// ?????????? УТЕЧКА ПАМЯТИ ????????????
	ws.repo[(*e).ID] = e
	return nil
}

func (ws *WorkStorage) Delete(id string) error {
	if ws.contains(id) {
		delete(ws.repo, id)
		return nil
	}
	return fmt.Errorf("entity not exit")
}

func (ws *WorkStorage) contains(id string) bool {
	for entityID := range ws.repo {
		if entityID == id {
			return true
		}
	}
	return false
}
