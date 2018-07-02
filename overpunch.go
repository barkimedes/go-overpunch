package overpunch

import (
	"fmt"
	"strconv"
	"strings"
)

// positivityMap tells whether the overpunch character indicates a positive
// value (true) or a negative value (false).
var positivityMap = map[rune]bool{
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

// digitMap gives the numeric digit that the overpunch character maps to.
var digitMap = map[rune]rune{
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

// positiveMap gives the overpunch character if the overall number is positive.
var positiveMap = map[rune]rune{
	'0': '{',
	'1': 'A',
	'2': 'B',
	'3': 'C',
	'4': 'D',
	'5': 'E',
	'6': 'F',
	'7': 'G',
	'8': 'H',
	'9': 'I',
}

// negativeMap gives the overpunch character if the overall number is negative.
var negativeMap = map[rune]rune{
	'0': '}',
	'1': 'J',
	'2': 'K',
	'3': 'L',
	'4': 'M',
	'5': 'N',
	'6': 'O',
	'7': 'P',
	'8': 'Q',
	'9': 'R',
}

// Overpunch is an integer value that can be encoded/decoded from an overpunch
// text value.
// Note that this doesn't support decimal values since overpunch-encoded strings
// often do not explicitly support decimals.
type Overpunch int

// abs returns the absolute value of the given integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// MarshalText marshals the integer into a textual overpunched format.
// See https://golang.org/pkg/encoding/#TextMarshaler.
func (o Overpunch) MarshalText() (text []byte, err error) {
	sign := int(o) >= 0
	value := strconv.Itoa(abs(int(o)))
	lastDigit := rune(value[len(value)-1])
	var overpunchedLastDigit rune
	if sign {
		overpunchedLastDigit = positiveMap[lastDigit]
	} else {
		overpunchedLastDigit = negativeMap[lastDigit]
	}
	encoded := fmt.Sprintf("%v%v", value[0:len(value)-1], string(overpunchedLastDigit))
	return []byte(encoded), nil
}

// UnmarshalText unmarshals an overpunch-encoded string into the integer value.
// Note that whatever value was stored in o will be overwritten.
// If there is an error in the unmarshaling process, it will be set to 0.
// an error in the unmarshaling process.
// See https://golang.org/pkg/encoding/#TextUnmarshaler.
func (o *Overpunch) UnmarshalText(text []byte) error {
	value := strings.ToUpper(string(text))
	if len(value) == 0 {
		*o = 0
		return fmt.Errorf("Unable to unmarshal an empty string")
	}
	char := rune(value[len(value)-1])
	isPositive, ok := positivityMap[char]
	if !ok {
		*o = 0
		return fmt.Errorf("provided text (%v) is not a valid overpunch value", value)
	}
	if !isPositive {
		value = fmt.Sprintf("-%v", value)
	}
	value = fmt.Sprintf("%v%v", value[0:len(value)-1], string(digitMap[char]))
	decoded, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	*o = Overpunch(decoded)
	return nil
}

// String returns the integer as a normal integer string.
func (o Overpunch) String() string {
	return strconv.Itoa(int(o))
}
