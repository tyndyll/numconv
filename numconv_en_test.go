package numconv_test

import (
	"fmt"
	"testing"

	"github.com/tyndyll/numconv"
)

func TestItoa(t *testing.T) {
	testCases := []struct {
		i   int
		a   string
		err error
	}{
		{1, `one`, nil},
		{2, `two`, nil},
		{3, `three`, nil},
		{4, `four`, nil},
		{5, `five`, nil},
		{6, `six`, nil},
		{7, `seven`, nil},
		{8, `eight`, nil},
		{9, `nine`, nil},
		{10, `ten`, nil},
		{11, `eleven`, nil},
		{12, `twelve`, nil},
		{13, `thirteen`, nil},
		{14, `fourteen`, nil},
		{15, `fifteen`, nil},
		{16, `sixteen`, nil},
		{17, `seventeen`, nil},
		{18, `eighteen`, nil},
		{19, `nineteen`, nil},
		{20, `twenty`, nil},
		{21, `twenty one`, nil},
		{30, `thirty`, nil},
		{31, `thirty one`, nil},
		{40, `forty`, nil},
		{41, `forty one`, nil},
		{51, `fifty one`, nil},
		{61, `sixty one`, nil},
		{71, `seventy one`, nil},
		{81, `eighty one`, nil},
		{91, `ninety one`, nil},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`Converting %d to %s`, tc.i, tc.a), func(t *testing.T) {
			returnedValue, returnedError := numconv.Itoa(tc.i)
			if returnedValue != tc.a {
				t.Errorf(`got "%s", want "%s"`, returnedValue, tc.a)
			}

			if returnedError != tc.err {
				t.Errorf(`got "%s", want "%s"`, returnedError, tc.err)
			}
		})
	}
}

func TestAtoi(t *testing.T) {
	testCases := []struct {
		i      int
		a      string
		errNil bool
	}{
		{0, ``, false},
		{1, `ONE`, true},
		{1, `one`, true},
		{2, `two`, true},
		{3, `three`, true},
		{4, `four`, true},
		{5, `five`, true},
		{6, `six`, true},
		{7, `seven`, true},
		{8, `eight`, true},
		{9, `nine`, true},
		{10, `ten`, true},
		{11, `eleven`, true},
		{12, `twelve`, true},
		{13, `thirteen`, true},
		{14, `fourteen`, true},
		{15, `fifteen`, true},
		{16, `sixteen`, true},
		{17, `seventeen`, true},
		{18, `eighteen`, true},
		{19, `nineteen`, true},
		{0, `nineteen nineteen`, false},
		{20, `twenty`, true},
		{0, `twenty eleven`, false},
		{21, `twenty one`, true},
		{30, `thirty`, true},
		{31, `thirty one`, true},
		{40, `forty`, true},
		{41, `forty one`, true},
		{51, `fifty one`, true},
		{61, `sixty one`, true},
		{71, `seventy one`, true},
		{81, `eighty one`, true},
		{91, `ninety one`, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`Converting %s to %d`, tc.a, tc.i), func(t *testing.T) {
			returnedValue, returnedError := numconv.Atoi(tc.a)
			if returnedValue != tc.i {
				t.Errorf(`got "%d", want "%d"`, returnedValue, tc.i)
			}

			if (returnedError != nil) && tc.errNil {
				t.Errorf(`got "%s", want "%s"`, returnedError, tc.errNil)
			}
		})
	}
}
