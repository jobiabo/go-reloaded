package helper

import (
	"strconv"
	"strings"
)

func ModAphaNum(input string) string {
	sliceInput := strings.Fields(input)

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
				count := 1

				num, err := strconv.Atoi(strings.TrimSpace(numAttach))
				if err == nil {
					count = num
				}

				for j := 1; j <= count; j++ {
					if i-j < 0 {
						break
					}
					w := strings.ToLower(sliceInput[i-j])
					if len(w) > 0 {
						sliceInput[i-j] = strings.ToUpper(string(w[0])) + w[1:]

						if len(sliceInput) > i {
							sliceInput = append(sliceInput[:i], sliceInput[i:]...)
						}
					}

				}
			case "(cap)":
				w := strings.ToLower(sliceInput[i-1])
				if len(w) > 0 {
					sliceInput[i-1] = strings.ToUpper(string(w[0])) + w[1:]

					if len(sliceInput) > i {
						sliceInput = append(sliceInput[:i], sliceInput[i+1:]...)
					}

				}

			}
		}
	}
	return strings.Join(sliceInput, " ")
}
