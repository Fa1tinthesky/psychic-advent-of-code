package main

import (
	"fmt"
	_ "maps"
	"strconv"
	"strings"
)

func fmt_arr[T any](arr []T) {
	for _, v := range arr {
		fmt.Printf("%s\n", v)
	}
}

func parse_movement(move string) (int, int, int) {
	divided := strings.Fields(move)

	amount_to_move, _ := strconv.Atoi(divided[1])
	from, _ := strconv.Atoi(divided[3])
	to, _ := strconv.Atoi(divided[5])

	return amount_to_move, from, to
}

func init_map(matrix [][]rune) map[int]int {
	mp := make(map[int]int)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if (mp[j] == 0) { mp[j] = 0 }
			if (matrix[i][j] != 0) { mp[j] += 1 }
		}
	}

	return mp
}

func result(matrix [][]rune) string {
	results := make([]rune, 3)

	flag := false
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			if (matrix[j][i] == 0) { flag = true }
			if (flag == true && matrix[j][i] != 0) {
				results = append(results, matrix[j][i])
				flag = false
				j++
			}
		}
	}

	return string(results)
}

func copyMatrix(matrix [][]rune) [][]rune {
    cp := make([][]rune, len(matrix))
    for i := range matrix {
        cp[i] = make([]rune, len(matrix[i]))
        copy(cp[i], matrix[i])  // built-in copy for each row
    }
    return cp
}

func parse_v2(content string) [][]rune {	
	f := strings.Split(content, "\n")
	
	m := (len(f[0]) - len(f[0]) / 4) / 3
	var n int

	// fmt.Printf("%q\n", f[0][1])
	for i := 0; i < len(f); i++ {
		if strings.Index(f[i], "1   2   3") != -1 { n = i + 1 }
	}

	// Creating [NxM] matrix filled with empty chars
	matrix := make([][]rune, n)
	for i := range matrix {
		matrix[i] = make([]rune, m)
	}
	
	// Mapping last line of matrix which is static:
	last_line := matrix[len(matrix) - 1]
	for i := 0; i < len(last_line); i++ {
		matrix[len(matrix) - 1][i] = rune(i + 49)
	}
	
	// Maping file contents to matrix
	for i, v := range f[:n-1] {
		// fmt.Println(i, v)

		flag := false
		for j := 0; j < len(v); j++ {
			if rune(v[j]) == ']' {
				flag = false
			} else if rune(v[j]) == '[' {
				flag = true
			} else if flag {
				matrix[i][(j - j / 4) / 3] = rune(v[j])
			}
		}
	} 

	// Moving everything

	for _, v := range f[n+1:] {
		amount_to_move, from, to := parse_movement(v)

		matrix = move_block(matrix, amount_to_move, from, to)
	}

	return matrix
}


