package main

import "fmt"
import "time"
//import "sync"
import "github.com/MTalda/gol/grid"

type Element struct {
	row int
	col int
}

func worker(id int, elementChan <-chan Element, done chan<- int, grid1 *grid.Grid) {
	for element := range elementChan {
		grid.CellNextState(element.row, element.col, grid1)
		//fmt.Printf("Worker #%d: finished element[%d, %d]\n", id, element.row, element.col)
		done <- 1
	}
}

func main() {

	// Size
	rows := 10
	cols := 10
	numWorkers := 10

	// Create a Grid
	grid1, _ := grid.CreateGrid(rows, cols)
	grid.RandGrid(&grid1)

	// Channels
	elementChan := make(chan Element, rows * cols)
	done := make(chan int, rows * cols)

	// Spawn Goroutines
	for i := 0; i < numWorkers; i++ {
		go worker(i, elementChan, done, &grid1)
	}

	// Worker Pool
	t := time.Now()
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			elementChan <- Element{row, col}
		}
	}
	fmt.Println(time.Since(t))
	for i := 0; i < rows * cols; i++ {
		<- done
	}
	fmt.Println(time.Since(t))
	grid.PrintStates(&grid1)

	// Standard
	t = time.Now()
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			grid.CellNextState(row, col, &grid1)
		}
	}
	fmt.Println(time.Since(t))
	//grid.PrintStates(&grid1)
}