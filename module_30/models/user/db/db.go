package db

import "skillbox/module_30/models/user"

type db struct {
	repository user.Storage
}

func (d *db) Create(user user.User) (string, error) {
	create, err := d.repository.Create(user)
	if err != nil {
		return "", err
	}
	return create, nil
}

func (d *db) Read(id string) (user.User, error) {
	read, err := d.repository.Read(id)
	if err != nil {
		return user.User{}, err
	}
	return read, nil
}

func (d *db) ReadAll() ([]user.User, error) {
	read, err := d.repository.ReadAll()
	if err != nil {
		return nil, err
	}
	return read, nil
}

func (d *db) Update(user user.User) error {
	err := d.repository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (d *db) Delete(id string) error {
	err := d.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
