package main

import (
	"fmt"
	"os"
	_ "strconv"
	"strings"
	_ "time"
)

/* func draw_frame(matrix [][]rune) {
	writer.WriteString("\033[H")

	for _, row := range matrix {
		writer.WriteString(string(row))
		writer.WriteByte('\n')
	}

	writer.Flush()
}
*/
/* func get_render_matrix(matrix [][]rune) [][]rune {
    cp := make([][]rune, len(matrix))
    for i := range matrix {
        cp[i] = make([]rune, 4 * len(matrix[i]) - 1)

		// fmt.Printf("\n%c, %c, %c, length: %d\n", matrix[i][0], matrix[i][1], matrix[i][2], len(matrix[i]))

		for j := range matrix[i] {
			letter := '.'

			cp[i][j*4] = '['
			cp[i][j*4 + 2] = ']'

			if matrix[i][j] != 0 {
				letter = matrix[i][j]
			} else {
			cp[i][j*4] = ' '
			cp[i][j*4 + 2] = ' '
			}


			if j != len(matrix[i])-1 { cp[i][j*4 + 3] = ' '}

			cp[i][j*4 + 1] = letter
		}
    }

    return cp
} */


func append_empty_lines(matrix [][]rune, n int) [][]rune {
	for i := 0; i < n; i++ {
		// fmt.Println("ADDING LINE")
		matrix = append([][]rune{make([]rune, len(matrix[0]))}, matrix...)
	}

	return matrix
}

func map_sum(m map[int]int) int {
	total := 0

	for _, v := range m {
		total += v
	}

	return total
}

func insert_by_xy[T any](matrix [][]T, arr []T, x, y int, vertical bool) [][]T {
	if vertical {
		for i := 0; i < len(arr); i++ {
			matrix[i+y][x] = arr[i]
		}
	} else {
		for i := 0; i < len(arr); i++ {
			matrix[y][i+x] = arr[i]
		}
	}
	
	return matrix
}

func move_block(matrix [][]rune, amount_to_move, from, to int) [][]rune {
	stack_map := init_map(matrix)
	// fmt.Printf("STACK MAP: %q\n", stack_map)

	from -= 1
	to -= 1
	buff := 2
	
	// matrix = append_empty_lines(matrix, map_sum(stack_map) - len(matrix) - 2)

	for i := 0; i < amount_to_move; i++ {
		if (stack_map[to] + amount_to_move > len(matrix)) { 
			// Add 3 empty lines at once
			matrix = append_empty_lines(matrix, stack_map[to] + amount_to_move - len(matrix) + buff)
		}
		
		from_index := len(matrix) - stack_map[from]
		to_index := len(matrix) - stack_map[to] - 1

		// fmt.Printf("%c --- FROM: (%d,%d) --- TO: (%d, %d)\n",matrix[from_index][from],  from_index, from, to_index, to)

		matrix[to_index][to] = matrix[from_index][from]
		matrix[from_index][from] = 0

		// fmt.Printf("\nMODIFIED ARRAY %q\n", matrix)

		stack_map[from] -= 1
		stack_map[to]   += 1

	}

	return matrix
}


func main() {
		dat, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("Error reading line:", err)
		return
	}

	content := strings.Trim(string(dat), "\n")


	res := parse_v2(content)

	fmt.Printf(result(res))
}


