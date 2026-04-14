package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func ModAphaNum(input string) string {
	sliceInput := strings.Fields(input)
	var result []string

	for i := 0; i < len(sliceInput); i++ {
		if strings.HasPrefix(sliceInput[i], "(hex") || strings.HasPrefix(sliceInput[i], "(bin") ||
			strings.HasPrefix(sliceInput[i], "(cap") || strings.HasPrefix(sliceInput[i], "(low") ||
			strings.HasPrefix(sliceInput[i], "(low") {
			switch sliceInput[i] {

			case "(hex)":
				num, err := strconv.ParseInt(sliceInput[i-1], 16, 64)
				if err == nil {
					sliceInput[i-1] = strconv.FormatInt(num, 10)

				}
				sliceInput = append(sliceInput[:i], sliceInput[i+1:]...)
				i--

			case "(bin)":
				num, err := strconv.ParseInt(sliceInput[i-1], 2, 64)
				if err == nil {
					sliceInput[i-1] = strconv.FormatInt(num, 10)
				}
				sliceInput = append(sliceInput[:i], sliceInput[i+1:]...)
				i--

			case "(cap,":
				numAttach := strings.Trim(sliceInput[i+1], ")")

				num, err := strconv.Atoi(strings.TrimSpace(numAttach))
				if err == nil {
					fmt.Println(err)
				}

				for j := len(result) - num; j < len(result); j++ {
					if len(result[j]) > 0 {
						result[j] = strings.ToUpper(string(result[j][0])) + strings.ToLower(result[j][1:])

						result = append(result[:i], result[i+1:]...)
						i--

					} else {
						result = append(result, sliceInput[i])
					}

				}

			}
		}
	}
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(ModAphaNum("my name is ben"))
}
