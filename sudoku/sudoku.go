package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	// transformation des strings -> tableau 2D de int
	grid := parseInput(args)
	if InputCheck(grid) { // check si les strings sont résolvable
		if solvesodu(grid) { // résolution du sudoku et le print
			fmt.Println()
			for i := 0; i < 9; i++ {
				if i == 0 {
					fmt.Println("╔═══════╦═══════╦═══════╗")
				}
				if i == 3 || i == 6 {
					fmt.Println("╠═══════╬═══════╬═══════╣")
				}
				for j := 0; j < 9; j++ {
					if j == 0 {
						fmt.Print("║")
					}
					if j == 3 || j == 6 {
						fmt.Print(" ║")
					}
					fmt.Print(" ", grid[i][j])
				}
				fmt.Println(" ║")
			}
			fmt.Println("╚═══════╩═══════╩═══════╝")
			fmt.Println()
		} else { // si sudoku n'a aucune solution
			fmt.Println("Error: Sudoku cannot be solved")
		}
	} else { // si l'input n'est pas un puzzle sudoku résolvable
		fmt.Println("Error: Invalid Sudoku input")
	}
}

func solvesodu(grid [][]int) bool {
	// resoudre le sudoku
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				for nb := 1; nb <= 9; nb++ {
					// si le nombre respecte les regles du jeu
					if numCheck(grid, i, j, nb) {
						grid[i][j] = nb
						if solvesodu(grid) {
							return true
						}
						// backtracking algorithme
						grid[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func numCheck(grid [][]int, row, col, nb int) bool {
	// recherche par ligne et colonne si le nb respecte les regles du jeu
	for i := 0; i < 9; i++ {
		if grid[row][i] == nb || grid[i][col] == nb {
			return false
		}
	}
	// recherche par carré de 3x3 si le nb respecte les regles du jeu
	row1 := row - row%3
	col1 := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+row1][j+col1] == nb {
				return false
			}
		}
	}
	return true
}

func parseInput(args []string) [][]int {
	// si on a pas 9 strings
	if len(args) != 9 {
		fmt.Println("Error: Invalid number of rows")
		os.Exit(0)
	}
	grid := make([][]int, 9)
	for i, rowStr := range args {
		// si la longeur d'une des strings n'est pas 9
		if len(rowStr) != 9 {
			fmt.Println("Error: Invalid number of colums")
			os.Exit(0)
		}
		grid[i] = make([]int, 9)
		for j, char := range rowStr {
			if char == '.' {
				grid[i][j] = 0
			} else if char >= '1' && char <= '9' {
				digit := int(char - '0')
				grid[i][j] = digit
			} else {
				// si on a un character qui n'est pas numérique
				fmt.Println("Error: Invalid character in input")
				os.Exit(0)
			}
		}
	}
	return grid
}

func InputCheck(grid [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] != 0 {
				// Check lignes
				for j2 := 0; j2 < 9; j2++ {
					if j != j2 && grid[i][j] == grid[i][j2] {
						return false
					}
				}
				// Check colonne
				for i2 := 0; i2 < 9; i2++ {
					if i != i2 && grid[i][j] == grid[i2][j] {
						return false
					}
				}
				// Check 3x3 carré
				row1 := i - i%3
				col1 := j - j%3
				for x := 0; x < 3; x++ {
					for y := 0; y < 3; y++ {
						if grid[x+row1][y+col1] >= 1 && grid[x+row1][y+col1] <= 9 {
							for x2 := 0; x2 < 3; x2++ {
								for y2 := 0; y2 < 3; y2++ {
									if (x != x2 || y != y2) && grid[x+row1][y+col1] == grid[x2+row1][y2+col1] {
										return false
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return true
}
