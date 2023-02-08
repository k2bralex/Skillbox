package storage

import (
	"fmt"
	i "skillbox/module_28/instance"
)

type Group map[string]i.Worker

func (g Group) PrintGroup() {
	for _, v := range g {
		fmt.Println(v.String())
	}
}

func (g Group) Put(w i.Worker) error {
	if _, ok := g[w.GetName()]; ok {
		return fmt.Errorf("account exist")
	}
	g[w.GetName()] = w
	return nil
}

func (g Group) Get(s string) (i.Worker, error) {
	if val, ok := g[s]; !ok {
		return val, fmt.Errorf("account not found")
	}
	return nil, nil
}
