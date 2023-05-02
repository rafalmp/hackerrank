package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type place struct {
	score    int64
	position int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	n, err := strconv.Atoi(strings.TrimSpace(readLine(reader)))
	checkErr(err)

	ranked := makeSlice(n, strings.TrimSpace(readLine(reader)))

	m, err := strconv.Atoi(strings.TrimSpace(readLine(reader)))
	checkErr(err)

	player := makeSlice(m, strings.TrimSpace(readLine(reader)))

	positions := calculatePositions(ranked)

	climb(positions, player)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func makeSlice(n int, s string) []int64 {
	result := make([]int64, 0, n)
	nums := strings.Split(strings.TrimSpace(s), " ")

	for _, num := range nums {
		tmp, err := strconv.ParseInt(num, 10, 64)
		checkErr(err)
		result = append(result, tmp)
	}
	return result
}

func calculatePositions(ranked []int64) []place {
	result := make([]place, 0, len(ranked))
	score := ranked[0]
	pos := 1
	for _, s := range ranked {
		if s < score {
			pos++
			score = s
		}
		result = append(result, place{s, pos})
	}
	return result
}

func climb(positions []place, player []int64) {
	for _, score := range player {
		cur := 1
		for _, place := range positions {
			if score < place.score {
				cur = place.position + 1
			}
		}
		fmt.Println(cur)
	}
}
