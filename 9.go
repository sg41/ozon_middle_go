package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type edge struct {
	to     int
	weight int
}

type node struct {
	id       int
	priority int
	index    int
}

type priorityQueue []*node

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func minimumCost(numPoints int, routerCosts []int, connections [][3]int) int {
	graph := make([][]edge, numPoints)
	for _, conn := range connections {
		u, v, w := conn[0]-1, conn[1]-1, conn[2]
		graph[u] = append(graph[u], edge{to: v, weight: w})
		graph[v] = append(graph[v], edge{to: u, weight: w})
	}

	visited := make([]bool, numPoints)
	totalCost := 0

	pq := make(priorityQueue, 0)
	heap.Init(&pq)

	startNode := 0 // You can choose any starting node
	heap.Push(&pq, &node{id: startNode, priority: routerCosts[startNode]})

	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(*node)
		if visited[currentNode.id] {
			continue
		}

		visited[currentNode.id] = true
		totalCost += currentNode.priority

		for _, nextEdge := range graph[currentNode.id] {
			if !visited[nextEdge.to] {
				heap.Push(&pq, &node{id: nextEdge.to, priority: nextEdge.weight})
			}
		}
	}

	return totalCost
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, k, m int
		fmt.Fscan(reader, &n, &k)

		fmt.Fscan(reader, &m)
		boxes := make([]int, m)
		for j := 0; j < m; j++ {
			var a int
			fmt.Fscan(reader, &a)
			boxes[j] = 1 << a // Calculate box weight as 2^a
		}
	}

	minCostResult := minimumCost(n, routerCosts, connections)
	fmt.Fprintln(writer, minCostResult)
	// ...
}
