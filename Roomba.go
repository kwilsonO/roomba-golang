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


var allDirt []*Vertex
var prev map[*Vertex]*Vertex
var dist map[*Vertex]int
var Rx, Ry, Dx, Dy int
var source Vertex


type Vertex struct{
	X, Y int
	index int
	Weight int
	Visited bool
}

type PriorityQueue []*Vertex

var pq PriorityQueue 

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
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
		*pq = append(*pq, vert)
}

func (pq *PriorityQueue) Pop() interface{} {

		old := *pq
		n := len(old)
		v := old[n-1]
		v.index = -1 //for safety??
		*pq = old[0 : n -1]
		return v
}

func (pq *PriorityQueue) update(v *Vertex, weight int) {
	v.Weight = weight
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
	dist = make(map[*Vertex]int)
	prev = make(map[*Vertex]*Vertex)
	var s []string
	pq = make(PriorityQueue, len(lines))
	in := 0
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
			Dx = x;
			Dy = y;
			continue
		} else if  i == 1 {
			Rx = x;
			Ry = y;
			source = Vertex{X: x, Y: y, Weight: 0, Visited: true}
		}


		v := Vertex{X: x, Y: y, Weight: math.MaxInt32, Visited: false}
		allDirt = append(allDirt, &v)
		pq[in] = &v
		in++
		dist[&v] = math.MaxInt32
		prev[&v] = nil
	}

	pq[in] = &source
	in++
	heap.Init(&pq)

	FindPath()

	fmt.Println(dist)
	fmt.Println(prev)

}


func FindPath(){

	for pq.Len() > 0 {

		u := heap.Pop(&pq);
		u = u.(*Vertex)

		for _, v := range allDirt {

			if v == u {
				continue
			}

			alt := dist[u.(*Vertex)] + GetDist(u.(*Vertex), v)
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u.(*Vertex)
				pq.update(v, alt)
			}

		}


	}
}

func GetDist(u, v *Vertex) int {

	ret := int(math.Abs(float64((u.X - v.X)) + float64((u.Y - v.Y))))
	return ret
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
