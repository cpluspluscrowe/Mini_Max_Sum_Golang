package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	arrTemp := strings.Split(readLine(reader), " ")
	var arr []int64
	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int64(arrItemTemp)
		arr = append(arr, arrItem)
	}
	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getTotal(arr []int64) int64 {
	var total int64
	for _, item := range arr {
		total += item
	}
	return total
}

func miniMaxSum(arr []int64) {
	var total int64 = getTotal(arr)
	var max int64 = arr[0]
	var min int64 = total
	for _, item := range arr {
		var minusItem int64 = total - item
		if minusItem > max {
			max = minusItem
		}
		if minusItem < min {
			min = minusItem
		}
	}
	fmt.Println(min, max)
}
