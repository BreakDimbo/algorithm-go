package search

import (
	"container/list"
	"fmt"

	"github.com/golang-collections/collections/queue"
)

// 无向图
type Graph struct {
	vertex int // 顶点个数
	adj    []list.List
}

var found = false

func NewGraph(v int) *Graph {
	return &Graph{
		vertex: v,
		adj:    make([]list.List, v),
	}
}

func (g *Graph) Add(s, t int) {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}

// search the shortest route from s to t
func (g *Graph) BFS(s, t int) {
	if s == t {
		return
	}

	// visited 记录已经被访问的节点
	visited := make([]bool, g.vertex)
	visited[s] = true
	// queue 记录已经被访问但是后继节点没有被访问的节点
	queue := queue.New()
	queue.Enqueue(s)
	// prev 反向记录访问过的边
	// init prev
	prev := make([]int, g.vertex)
	for i := range prev {
		prev[i] = -1
	}

	for queue.Len() != 0 {
		w := queue.Dequeue().(int)

		for e := g.adj[w].Front(); e != nil; e = e.Next() {
			value := e.Value.(int)
			if visited[value] {
				continue
			}

			prev[value] = w
			if value == t {
				printRoute(prev, s, t)
				return
			}

			visited[value] = true
			queue.Enqueue(value)
		}
		fmt.Printf("prev: %v\n", prev)
	}
}

func (g *Graph) DFS(s, t int) {
	visited := make([]bool, g.vertex)
	prev := make([]int, g.vertex)

	for i := range prev {
		prev[i] = -1
	}

	g.recurDFS(s, t, visited, prev)
	printRoute(prev, s, t)
}

func (g *Graph) recurDFS(startP, endP int, visited []bool, prev []int) {
	if found {
		return
	}

	visited[startP] = true

	if startP == endP {
		found = true
		return
	}

	for e := g.adj[startP].Front(); e != nil; e = e.Next() {
		currentP := e.Value.(int)
		if !visited[currentP] {
			visited[currentP] = true
			prev[currentP] = startP
			g.recurDFS(currentP, endP, visited, prev)
		}
	}
}

func (g *Graph) Print() {
	for k, v := range g.adj {
		if v.Len() == 0 {
			continue
		}
		fmt.Printf("vertext: %d points to: ", k)
		for e := v.Front(); e != nil; e = e.Next() {
			fmt.Printf("%d|", e.Value)
		}
		fmt.Println()
	}
}

func printRoute(prev []int, s, t int) {
	if prev[t] != -1 && s != t {
		printRoute(prev, s, prev[t])
	}
	fmt.Printf("%d->", t)
}
