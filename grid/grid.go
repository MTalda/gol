package grid

import "errors"
import "fmt"
import "sync"
//import "math/rand"

type Grid struct{
	// States
	currentState [][]bool
	nextState [][]bool

	// Numeric Information
	rows int
	columns int
	generation int

	// Lock
	lock sync.Mutex
}

//--------------------------------------------------------
// Name:			CreateGrid()
// Function:	Creates a 2D slice of Booleans
//--------------------------------------------------------
func CreateGrid(rows int, columns int) (Grid, error) {
	// Check Input
	if (rows < 1 || columns < 1) {
		return Grid{}, errors.New("rows and/or columns are invalid") 
	}
	
	// Allocate for Grid
	var grid Grid
	grid.rows = rows
	grid.columns = columns
	grid.generation = 0
	grid.currentState = make([][]bool, rows)
	cellsCurrent := make([]bool, columns*rows)
	grid.nextState = make([][]bool, rows)
	cellsNext := make([]bool, columns*rows)
	for i := range grid.currentState {
		grid.currentState[i], cellsCurrent = cellsCurrent[:columns], cellsCurrent[columns:]
		grid.nextState[i], cellsNext = cellsNext[:columns], cellsNext[columns:]
	}

	return grid, nil
}

//--------------------------------------------------------
// Name:			RandGrid()
// Function:	Randomizes the current grid 
//--------------------------------------------------------
func RandGrid(grid *Grid) {
	grid.currentState[5][5] = true
	grid.currentState[6][6] = true
	grid.currentState[7][4] = true
	grid.currentState[7][5] = true
	grid.currentState[7][6] = true
	
	/*
	for i := 0; i < grid.rows; i++ {
		for j := 0; j < grid.columns; j++ {
			grid.currentState[i][j] = (rand.Uint64()&(1<<63) == 0)
		}
	}
	*/
}

//--------------------------------------------------------
// Name:			PrintStates()
// Function:	Prints both Arrays 
//--------------------------------------------------------
func PrintStates(grid *Grid) {
	printCurrentGrid(grid)
	printNextGrid(grid)
}

//--------------------------------------------------------
// Name:			printCurrentGrid()
// Function:	Prints the current grid 
//--------------------------------------------------------
func printCurrentGrid(grid *Grid) {
	fmt.Printf("Current State\n-------------\n")
	for i := 0; i < grid.rows; i++ {
		for j := 0; j < grid.columns; j++ {
			if grid.currentState[i][j] {
				fmt.Printf("X")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

//--------------------------------------------------------
// Name:			printNextGrid()
// Function:	Prints the next grid 
//--------------------------------------------------------
func printNextGrid(grid *Grid) {
	fmt.Printf("Next State\n-------------\n")
	for i := 0; i < grid.rows; i++ {
		for j := 0; j < grid.columns; j++ {
			if grid.nextState[i][j] {
				fmt.Printf("X")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

//--------------------------------------------------------
// Name:			CellNextState()
// Function:	Determines if a Cell is dead or alive in
//					the next generation
//--------------------------------------------------------
func CellNextState(rowIndex int, columnIndex int, grid *Grid) error {
	if (rowIndex < 0 || columnIndex < 0 || rowIndex >= grid.rows || columnIndex >= grid.columns) {
		return errors.New("Indices are invalid")
	}

	cellsAdj := [3]int{-1, 0, 1}
	rowAdj := 0
	colAdj := 0
	liveCellsCount := 0

	for r := range cellsAdj {
		for c := range cellsAdj {
			// Skip Center
			if (cellsAdj[r] == 0 && cellsAdj[c] == 0) { continue }

			// Wrap Around Row
			if (rowIndex + cellsAdj[r] < 0) { 
				rowAdj = grid.rows - 1 
			} else if (rowIndex + cellsAdj[r] >= grid.rows) {
				rowAdj = 0
			} else {
				rowAdj = rowIndex + cellsAdj[r]
			}

			// Wrap Around Column
			if (columnIndex + cellsAdj[c] < 0) { 
				colAdj = grid.columns - 1
			} else if (columnIndex + cellsAdj[c] >= grid.columns) {
				colAdj = 0
			} else {
				colAdj = columnIndex + cellsAdj[c]
			}

			// Check if cell is alive
			if (grid.currentState[rowAdj][colAdj]) {
				liveCellsCount++
			}
		}
	}

	// Update Next State
	nextState := cellRule(grid.currentState[rowIndex][columnIndex], liveCellsCount)
	//grid.lock.Lock()
	//defer grid.lock.Unlock()
	grid.nextState[rowIndex][columnIndex] = nextState
	return nil
}

//--------------------------------------------------------
// Name:			cellRule()
// Function:	Determines cell' state based on live cells
//					adjacent to it
//--------------------------------------------------------
func cellRule(currentState bool, liveCells int) bool {
	/*
	    Any live cell with two or three neighbors survives.
    Any dead cell with three live neighbors becomes a live cell.
    All other live cells die in the next generation. Similarly, all other dead cells stay dead.
	 */
	
	if liveCells == 3 { 
		return true
	} else if currentState == true && liveCells == 2 {
		return true
	} else {
		return false
	}
}


/*
	r-1,c-1   r-1,c   r-1,c+1
	r,c-1             r,c+1              
   r+1,c-1   r+1,c   r+1,c+1
	
	// if r < 0 --> r = max
	// if r > max --> r = 0
*/