package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func HexBin(s string) string {
	ssplit := strings.Fields(s)

	for i := 0; i < len(ssplit); i++ {

		if strings.HasPrefix(ssplit[i], "(hex") || strings.HasPrefix(ssplit[i], "(bin") {

			cleanedi := strings.Trim(ssplit[i], "()")
			sepi := strings.Split(cleanedi, ",")
			pointv := strings.TrimSpace(sepi[0])
			count := 1

			if len(sepi) > 1 {
				num, err := strconv.Atoi(strings.TrimSpace(sepi[1]))
				if err != nil {
					fmt.Println("an error occured")
					count = num
				}
			}
			for j := 1; j <= count; j++ {
				if i-j < 0 {
					break
				}
				switch pointv {
				case "bin":
					numi, err := strconv.ParseInt(ssplit[i-j], 2, 64)
					if err != nil {
						return err.Error()
					}
					ssplit[i-j] = strconv.FormatInt(numi, 10)

				case "hex":
					numi, err := strconv.ParseInt(ssplit[i-j], 16, 64)
					if err != nil {
						return err.Error()
					}
					ssplit[i-j] = strconv.FormatInt(numi, 10)
				}
			}
			ssplit = append(ssplit[:i], ssplit[i+1:]...)
		}
	}
	return strings.Join(ssplit, " ")
}

func CapAlphaMod(s string) string {
	spcomma := strings.ReplaceAll(s, ", ", ",")
	ssplit := strings.Fields(spcomma)

	for i := 0; i < len(ssplit); i++ {
		if strings.HasPrefix(ssplit[i], "(up") ||
			strings.HasPrefix(ssplit[i], "(low") ||
			strings.HasPrefix(ssplit[i], "(cap") {

			cleanedi := strings.Trim(ssplit[i], "()")
			sepi := strings.Split(cleanedi, ",")
			point := strings.TrimSpace(sepi[0])
			count := 1

			if len(sepi) > 1 {
				num, err := strconv.Atoi(strings.TrimSpace(sepi[1]))
				if err == nil {
					count = num
				}
			}

			for j := 1; j <= count; j++ {
				if i-j < 0 {
					break
				}

				switch point {
				case "up":
					ssplit[i-j] = strings.ToUpper(ssplit[i-j])

				case "low":
					ssplit[i-j] = strings.ToLower(ssplit[i-j])

				case "cap":
					w := strings.ToLower(ssplit[i-j])
					if len(w) > 0 {
						ssplit[i-j] = strings.ToUpper(string(w[0])) + w[1:]
					}
				}
			}
			ssplit = append(ssplit[:i], ssplit[i+1:]...)
		}
	}
	return strings.Join(ssplit, " ")
}

func Vowels(s string) string {
	lows := strings.ToLower(s)
	ssplit := strings.Fields(lows)

	for i := 1; i < len(ssplit); i++ {
		if strings.HasPrefix(ssplit[i], "a") ||
			strings.HasPrefix(ssplit[i], "e") ||
			strings.HasPrefix(ssplit[i], "e") ||
			strings.HasPrefix(ssplit[i], "o") ||
			strings.HasPrefix(ssplit[i], "u") ||
			strings.HasPrefix(ssplit[i], "h") {

			cleanedi := strings.TrimSpace(string(ssplit[i][0]))

			switch cleanedi {
			case "a", "e", "i", "o", "u", "h":

				if ssplit[i-1] == "a" {
					ssplit[i-1] = "an"
				}

			}
		}
	}
	return strings.Join(ssplit, " ")
}

func PunctAAte(s string) string {
	s = strings.ReplaceAll(s, " ...", "... ")
	s = strings.ReplaceAll(s, " ,", ", ")
	s = strings.ReplaceAll(s, " ?", "?")
	return s
}
