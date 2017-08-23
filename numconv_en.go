package numconv

import "strings"

func init() {
	for num, text := range enNumberMap {
		enTextMap[text] = num
	}
}

type InvalidNumberString struct {
	error
}

func (invalid InvalidNumberString) Error() string {
	return "Bad number"
}

var enTextMap = map[string]int{}

var enNumberMap = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  `four`,
	5:  `five`,
	6:  `six`,
	7:  `seven`,
	8:  `eight`,
	9:  `nine`,
	10: `ten`,
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: `fourteen`,
	15: `fifteen`,
	16: `sixteen`,
	17: `seventeen`,
	18: `eighteen`,
	19: `nineteen`,
	20: `twenty`,
	30: `thirty`,
	40: `forty`,
	50: `fifty`,
	60: `sixty`,
	70: `seventy`,
	80: `eighty`,
	90: `ninety`,
}

// Itoa takes an integer and returns the textual representation. This function returns the English value
func Itoa(i int) (string, error) {
	var remainder int
	var returnValue string
	if i > 20 {
		returnValue = enNumberMap[(i/10)*10]
		if remainder = i % 10; remainder != 0 {
			returnValue += " " + enNumberMap[remainder]
		}
	} else {
		returnValue = enNumberMap[i]
	}
	return returnValue, nil
}

// Atoi takes the textual representation of a number and returns an integer.
func Atoi(a string) (int, error) {
	var returnValue int
	a = strings.ToLower(strings.TrimSpace(a))

	if a == "" {
		return 0, new(InvalidNumberString)
	}

	for _, part := range strings.Split(a, " ") {
		i := enTextMap[part]
		if i >= 11 && i <= 19 && enNumberMap[i] != a {
			return 0, new(InvalidNumberString)
		}
		returnValue += i
	}
	return returnValue, nil
}
