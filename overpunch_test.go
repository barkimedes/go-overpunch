package overpunch

import (
	"testing"
)

func TestOverpunchMarshalText(t *testing.T) {
	var tests = []struct {
		input int
		text  string
	}{
		{0, "{"},
		{1234, "123D"},
		{-1234, "123M"},
		{-240, "24}"},
	}
	for i, test := range tests {
		o := Overpunch(test.input)
		text, err := o.MarshalText()
		if string(text) != test.text {
			t.Errorf("test %d: got %v, expected %v", i, string(text), test.text)
		}
		if err != nil {
			t.Errorf("test %d: got unexpected error %v", i, err)
		}
	}
}

func TestOverpunchUnmarshalText(t *testing.T) {
	var tests = []struct {
		text   string
		number int
		err    string // empty err string -> nil error
	}{
		{"{", 0, ""},
		{"123D", 1234, ""},
		{"123m", -1234, ""},
		{"24}", -240, ""},
		{"A", 1, ""},
		{"j", -1, ""},
		{"", 0, "Unable to unmarshal an empty string"},
		{"AB", 0, `strconv.Atoi: parsing "A2": invalid syntax`},
		{"24Z", 0, "provided text (24Z) is not a valid overpunch value"},
	}
	for i, test := range tests {
		var o Overpunch
		err := (&o).UnmarshalText([]byte(test.text))
		if int(o) != test.number {
			t.Errorf("test %d: got %v, expected %v", i, int(o), test.number)
		}
		if err != nil {
			if test.err == "" {
				t.Errorf("test %d: got unexpected error %v", i, err)
			} else if test.err != err.Error() {
				t.Errorf("test %d: got err [%v], expected err [%v]", i, err, test.err)
			}
		}
	}
}

func TestOverpunchString(t *testing.T) {
	var tests = []struct {
		input int
		str   string
	}{
		{0, "0"},
		{1234, "1234"},
		{-1234, "-1234"},
		{-240, "-240"},
	}
	for i, test := range tests {
		o := Overpunch(test.input)
		str := o.String()
		if string(str) != test.str {
			t.Errorf("test %d: got %v, expected %v", i, str, test.str)
		}
	}
}
