package person

import "fmt"

type Person struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends Friends `json:"friends"`
}

type Friends []*Person

func (p *Person) FriendContains(id string) (int, bool) {
	for i, fs := range p.Friends {
		fmt.Println(fs.ID)
		if fs.ID == id {
			return i, true
		}
	}
	return -1, false
}

func (p *Person) FriendRemove(i int) string {
	name := p.Friends[i].Name
	p.Friends, p.Friends[len(p.Friends)-1] = append(p.Friends[:i], p.Friends[i+1:]...), nil
	return name
}

func (p *Person) AgeUpdate(age int) {
	p.Age = age
}
