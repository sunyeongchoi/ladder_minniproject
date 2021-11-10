package main

import "fmt"

const (
	row int = 12
	column int = 5
)
type Ladder struct {
	arr [row][column]int
}

func (ladder *Ladder) add(a int, b int) {
	var arr = &ladder.arr
	arr[a][b] = 1
	arr[a][b+1] = 2
}

func (ladder *Ladder) random() {
	ladder.add(2, 0)
	ladder.add(4, 0)
	ladder.add(8, 0)
	ladder.add(1, 1)
	ladder.add(6, 1)
	ladder.add(9, 1)
	ladder.add(3, 2)
	ladder.add(5, 2)
	ladder.add(8, 2)
	ladder.add(4, 3)
	ladder.add(7, 3)
	ladder.add(10, 3)
}

func (ladder *Ladder) start(index int) int {
	var visited = [row][column]int{}; var arr = &ladder.arr; var n = 0
	for n < row {
		if arr[n][index] == 1 || arr[n][index] == 2 {
			if arr[n][index] == 2 && arr[n][index-1] == 1 && visited[n][index-1] == 0{
				visited[n][index] = 1
				index--
			} else if arr[n][index] == 1 && arr[n][index+1] == 2 && visited[n][index+1] == 0 {
				visited[n][index] = 1
				index++
			} else {
				visited[n][index] = 1
				n++
			}
		} else {
			visited[n][index] = 1
			n++
		}
	}
	return index
}

func main()  {
	// [방법1] 배열
	//var result []string
	//member := []string{"김효경", "박재호", "박진휘", "이연지", "엄태준", "최수녕"}
	//rand.Seed(time.Now().Unix())
	//for _, number := range rand.Perm(5) {
	//	result = append(result, member[number])
	//}
	//fmt.Println(result)

	// [방법2] 사다리타기
	ladder := &Ladder{}
	ladder.random()
	var result []string
	member := []string{"김효경", "박재호", "박진휘", "이연지", "엄태준", "최수녕"}
	number := 0
	for number < column {
		result = append(result, member[ladder.start(number)])
		number++
	}
	fmt.Println(result)
}