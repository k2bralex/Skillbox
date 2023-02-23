package user

import "fmt"

type User struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends Friends `json:"friends"`
}

type Friends []*User

func (u *User) FriendContains(id string) (int, bool) {
	for i, fs := range u.Friends {
		fmt.Println(fs.ID)
		if fs.ID == id {
			return i, true
		}
	}
	return -1, false
}

func (u *User) DeleteFriend(i int) {
	if i < 0 {
		return
	}
	u.Friends, u.Friends[len(u.Friends)-1] = append(u.Friends[:i], u.Friends[i+1:]...), nil
}

func (u *User) GetFriendIndex(id string) int {
	for i, user := range u.Friends {
		if user.ID == id {
			return i
		}
	}
	return -1
}

func (u *User) AgeUpdate(age int) {
	u.Age = age
}

func (u *User) String() string {
	return fmt.Sprintf(
		"ID: %s\nName: %s\nAge: %d\n",
		u.ID,
		u.Name,
		u.Age)
}
