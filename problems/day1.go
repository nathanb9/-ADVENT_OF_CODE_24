package problems

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Day1 struct {
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    string  // The value of the item; arbitrary.
	priority float64 // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (Day1 *Day1) RunAndPrint() {
	pq1 := make(PriorityQueue, 1000) // 1k in
	pq2 := make(PriorityQueue, 1000)

	input_file, err := os.Open("problems/day1_input.txt")

	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	defer input_file.Close()

	line := bufio.NewReader(input_file)
	i := 0

	for {
		line, _, err := line.ReadLine()
		if err == io.EOF {
			break
		}
		curr_line := string(line)

		split_string := strings.Split(curr_line, "   ")

		casted_val_1, err := strconv.ParseFloat(split_string[0], 64)
		casted_val_2, err2 := strconv.ParseFloat(split_string[1], 64)

		if err != nil || err2 != nil {
			fmt.Println("Error casting string to int")
			panic(err)

		}
		pq1[i] = &Item{
			value:    string(split_string[0]),
			priority: casted_val_1,
			index:    i,
		}
		pq2[i] = &Item{
			value:    string(split_string[1]),
			priority: casted_val_2,
			index:    i,
		}
		i++
	}

	heap.Init(&pq1)
	heap.Init(&pq2)
	solution := 0.0

	for pq1.Len() > 0 && pq2.Len() > 0 {
		item1 := heap.Pop(&pq1).(*Item)
		item2 := heap.Pop(&pq2).(*Item)

		solution += math.Abs(item1.priority - item2.priority)
	}

	fmt.Println("Sol: ")
	fmt.Printf("%f\n", solution)
}
