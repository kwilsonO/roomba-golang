package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
	"container/heap"
)


type Vertex struct{
	X, Y int
	index int
	DjDist int
	Visited bool
}
type PriorityQueue Q []*Vertex
var dist map[Vertex]int
var prev map[Vertex]int
var Rx, Ry int
var source Vertex


func (pq PriorityQueue) []*Vertex Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	//maybe needs to be <
	return pq[i].DjDist > pq[j].DjDist
}

func (pq PriorityQueue) Swap(i, j int) {

	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {

		n := len(*pq)
		vert := x.(*Vertex)
		vert.index = n
		*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {

		old := *pq
		n := len(old)
		v := old[n-1]
		v.index = -1 //for safety??
		*pq = old[0 : n -1]
		return v
}

func (pa *PriorityQueue) update(v *Vertex, value string, weight int) {
	v.weight = weight
	heap.Fix(pq, v.index)
}


func main() {

	if len(os.Args) != 2 {
		fmt.Println("Send a file path noob")
	}

	lines, err := ParseLines(os.Args[1], func(s string)(string,bool){
		return s, true
	})

	if err != nil {
		fmt.Println("you're file format sucks", err);
		return
	}
	dist = make(map[Vertex]int)
	prev = make(map[Vertex]int)

	var s []string
	for i, l := range lines {

		s = strings.Split(l, " ")

		x, err := strconv.Atoi(s[0])
		y, err := strconv.Atoi(s[1])

		if err != nil {
			fmt.Println("really can't even type integers")
			return
		}

		//Room Dimensions
		if i == 0 {
			Rx = x;
			Ry = y;
			continue
		} else if  i == 1 {
			source = Vertex{X; x, Y: y, Dist: 0, DjDist: 0, Visisted: true}
			continue
		}


		v := Vertex{X: x, Y: y, Dist: (x + y), DjDist: math.MaxInt32, Visited: false}

		dist[v] = math.MaxInt32
		prev[v] = -1
		Q = append(Q, v)
	}
}



func ParseLines(filePath string, parse func(string) (string, bool)) ([]string, error){
	inputFile, err := os.Open(filePath);
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile);
	var results []string
	for scanner.Scan() {
		if output, add := parse(scanner.Text()); add {
			results = append(results, output)
		}
	}

	if err  := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil

}
