package crossword

type Crossword struct {
	Grid    [][]string
	Numbers [][]string
	Rows    int
	Cols    int
}

func New() *Crossword {
	return &Crossword{}
}

func (c *Crossword) SetGrid(grid [][]string) {
	c.Grid = grid
	if len(grid) > 0 {
		c.Rows = len(grid)
		c.Cols = len(grid[0])
	}
}

func (c *Crossword) SetNumbers(numbers [][]string) {
	c.Numbers = numbers
}

func (c *Crossword) GetCell(row, col int) string {
	if row < 0 || row >= c.Rows || col < 0 || col >= c.Cols {
		return ""
	}
	return c.Grid[row][col]
}

func (c *Crossword) GetNumber(row, col int) string {
	if c.Numbers == nil || row < 0 || row >= len(c.Numbers) || col < 0 || col >= len(c.Numbers[row]) {
		return ""
	}
	return c.Numbers[row][col]
}