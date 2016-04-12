import (
	"bufio"
	"fmt"
	"os"
)


type Vertex struct{
	X, Y int
	dist int
	prev int
}

var verts []Vertex
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

	verts = make(map[Vertex]int)

	for _, 1 := range lines {
		//do some data handling
	}
}



func ParseLines(filePath string, parse func(strin) (string, bool)){
	inputFile, err = os.Open(filePath);
	if err != nil {
		return nil, err
	}

	defer inputFile.close()

	scanner := bufio.NewScanner(inputFile);
	var results []string
	for scanner.Scan() {
		if output, add := parse(scanner.Text()); add {
			results.append(results, output)
		}
	}

	if err  := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil

}
