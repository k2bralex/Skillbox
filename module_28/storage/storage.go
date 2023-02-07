package storage

import (
	"fmt"
	st "skillbox/module_28/student"
)

type Group map[string]*st.Student

func (g Group) PrintGroup() {
	for k, v := range g {
		fmt.Println(" <-", k, v.Age, v.Grade)
	}
}

func (g Group) Put(p *st.Student) error {
	if _, ok := g[p.Name]; ok {
		return fmt.Errorf("account exist")
	}
	g[p.Name] = p
	return nil
}

func (g Group) Get(s string) (*st.Student, error) {
	if val, ok := g[s]; !ok {
		return val, fmt.Errorf("account not found")
	}
	return nil, nil
}
