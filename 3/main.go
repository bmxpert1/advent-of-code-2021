package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func mode(nums []int64, prefer int64) (maxKey int64) {
	occurrences := map[int64]int{}
	max := 0

	for _, num := range nums {
		occurrences[num]++
	}

	for k, v := range occurrences {
		if v == max && prefer > -1 {
			maxKey = prefer
		} else if v > max {
			max = v
			maxKey = k
		}
	}

	return maxKey
}

func invert(str string) (inverted string) {
	for _, c := range str {
		if c == '0' {
			inverted += "1"
		} else {
			inverted += "0"
		}
	}

	return inverted
}

func main() {
	// read in json data
	var data []string
	content, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		fmt.Println(err)
	}

	////////////////////////////////////////////
	// challenge 1
	gammaBinStr := ""

	for i := 0; i < len(data[0]); i++ {
		posVals := make([]int64, len(data))

		for ii, binStr := range data {
			posVals[ii], _ = strconv.ParseInt(string(binStr[i]), 10, 64)
		}

		gammaBinStr += strconv.Itoa(int(mode(posVals, -1)))
	}

	epsilonBinStr := invert(gammaBinStr)

	gammaDecimal, _ := strconv.ParseInt(gammaBinStr, 2, 64)
	epsilonDecimal, _ := strconv.ParseInt(epsilonBinStr, 2, 64)

	fmt.Println(gammaDecimal * epsilonDecimal)

	//////////////////////////////////////////////
	// challenge 2
	runningO2 := make([]string, len(data))
	copy(runningO2, data)
	var finalO2 int64
	//
	runningCO2 := make([]string, len(data))
	copy(runningCO2, data)
	var finalCO2 int64

	// o2
	for i := 0; i < len(runningO2[0]); i++ {
		posVals := make([]int64, len(runningO2))

		for ii, binStr := range runningO2 {
			posVals[ii], _ = strconv.ParseInt(string(binStr[i]), 10, 64)
		}

		mode := mode(posVals, 1)
		newRunningO2 := []string{}

		// filter o2
		for _, binStr := range runningO2 {
			if string(binStr[i]) == strconv.Itoa(int(mode)) {
				newRunningO2 = append(newRunningO2, binStr)
			}
		}

		runningO2 = newRunningO2

		if len(runningO2) == 1 {
			finalO2, _ = strconv.ParseInt(runningO2[0], 2, 64)
			break
		}
	}

	// co2
	for i := 0; i < len(runningCO2[0]); i++ {
		posVals := make([]int64, len(runningCO2))

		for ii, binStr := range runningCO2 {
			posVals[ii], _ = strconv.ParseInt(string(binStr[i]), 10, 64)
		}

		mode := mode(posVals, 1)
		antimode := invert(strconv.Itoa(int(mode)))
		newRunningCO2 := []string{}

		// filter co2
		for _, binStr := range runningCO2 {
			if string(binStr[i]) == antimode {
				newRunningCO2 = append(newRunningCO2, binStr)
			}
		}

		runningCO2 = newRunningCO2

		if len(runningCO2) == 1 {
			finalCO2, _ = strconv.ParseInt(runningCO2[0], 2, 64)
			break
		}
	}

	fmt.Println(finalO2 * finalCO2)
}
