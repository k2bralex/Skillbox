package user

import (
	"fmt"
	"strconv"
)

var Id = 1010

type User struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends Friends `json:"friends"`
}

type Friends []*User

func NewUser() *User {
	Id++
	return &User{ID: strconv.Itoa(Id)}
}

func (u *User) FriendContains(id string) error {
	for _, fs := range u.Friends {
		fmt.Println(fs.ID)
		if fs.ID == id {
			return fmt.Errorf("already in a friend list")
		}
	}
	return nil
}

func (u *User) AddFriend(t *User) error {
	if err := u.FriendContains(t.ID); err != nil {
		return err
	}
	u.Friends = append(u.Friends, t)
	return nil
}

func (u *User) getFriendIndex(id string) int {
	for i, user := range u.Friends {
		if user.ID == id {
			return i
		}
	}
	return -1
}

func (u *User) DeleteFriend(id string) error {
	i := u.getFriendIndex(id)
	if i < 0 {
		return fmt.Errorf("not in friend list")
	}
	u.Friends, u.Friends[len(u.Friends)-1] = append(u.Friends[:i], u.Friends[i+1:]...), nil
	return nil
}

func (u *User) String() string {
	return fmt.Sprintf(
		"ID: %s\nName: %s\nAge: %d\n",
		u.ID,
		u.Name,
		u.Age)
}
