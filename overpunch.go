package overpunch

import (
  "fmt"
)

var overpunchSigns = map[rune]bool {
  	'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
		'{': true,
		'}': false,
		'A': true,
		'J': false,
		'B': true,
		'K': false,
		'C': true,
		'L': false,
		'D': true,
		'M': false,
		'E': true,
		'N': false,
		'F': true,
		'O': false,
		'G': true,
		'P': false,
		'H': true,
		'Q': false,
		'I': true,
		'R': false,
}

var overpunchDigits = map[rune]rune {
		'0': '0',
		'1': '1',
		'2': '2',
		'3': '3',
		'4': '4',
		'5': '5',
		'6': '6',
		'7': '7',
		'8': '8',
		'9': '9',
		'{': '0',
		'}': '0',
		'A': '1',
		'J': '1',
		'B': '2',
		'K': '2',
		'C': '3',
		'L': '3',
		'D': '4',
		'M': '4',
		'E': '5',
		'N': '5',
		'F': '6',
		'O': '6',
		'G': '7',
		'P': '7',
		'H': '8',
		'Q': '8',
		'I': '9',
		'R': '9',
}


func applyOverpunch(value string) string {
	if len(value) == 0 {
		return ""
	}
	char := rune(value[len(value) - 1])
	if isPositive, ok := overpunchSigns[char]; ok {
		if !isPositive {
value = fmt.Sprintf("-%v", value)
}
value = fmt.Sprintf("%v%v", value[0:len(value) - 1], string(overpunchDigits[char]))
}
return value
}
