package graph

import "fmt"

type graph struct {
	length  int
	adjlist map[interface{}][]interface{}
}

func New() *graph {
	return &graph{0, make(map[interface{}][]interface{})}
}

func (g *graph) Addnode(val interface{}) {
	g.adjlist[val] = nil
	g.length++
}

func (g *graph) Addedge(n1, n2 interface{}) {
	g.adjlist[n1] = append(g.adjlist[n1], n2)
	g.adjlist[n2] = append(g.adjlist[n2], n1)
}

func (g *graph) Print() {
	for k, v := range g.adjlist {
		fmt.Printf("%v--->", k)
		for _, val := range v {
			fmt.Printf("%v ", val)
		}
		fmt.Println()
	}
}
