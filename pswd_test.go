package pswd

import (
	"strings"
	"testing"
)

func TestPasswordLength(t *testing.T) {
	cases := []struct {
		in     int
		want   int
		weight PasswordWeight
	}{
		{24, 24, PasswordWeight{Lower: 4, Upper: 3, Digit: 3, Symbol: 2}},
		{3, 3, PasswordWeight{Lower: 4, Upper: 3, Digit: 3, Symbol: 2}},
		{200, 200, PasswordWeight{Lower: 4, Upper: 3, Digit: 3, Symbol: 2}},
		{2000, 2000, PasswordWeight{Lower: 4, Upper: 3, Digit: 3, Symbol: 2}},
		{50, 0, PasswordWeight{Lower: 0, Upper: 0, Digit: 0, Symbol: 0}},
		{10, 10, PasswordWeight{Lower: 4, Upper: 3, Digit: 3, Symbol: 2}},
		{10, 10, PasswordWeight{Lower: 1, Upper: 1, Digit: 1, Symbol: 1}},
		{10, 10, PasswordWeight{Lower: 0, Upper: 0, Digit: 3, Symbol: 3}},
		{10, 10, PasswordWeight{Lower: 1, Upper: 0, Digit: 0, Symbol: 0}},
		{100, 100, PasswordWeight{Lower: 4, Upper: 3, Digit: 3, Symbol: 2}},
		{100, 100, PasswordWeight{Lower: 100, Upper: 100, Digit: 100, Symbol: 100}},
		{0, 0, PasswordWeight{Lower: 4, Upper: 3, Digit: 3, Symbol: 2}},
	}
	for _, c := range cases {
		got := len(Password(c.in, c.weight))
		if got != c.want {
			t.Errorf("len(Password(%q)) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestPasswordOnlySym(t *testing.T) {
	cases := []struct {
		in     int
		weight PasswordWeight
	}{
		{24, PasswordWeight{Lower: 0, Upper: 0, Digit: 0, Symbol: 1}},
		{3, PasswordWeight{Lower: 0, Upper: 0, Digit: 0, Symbol: 1}},
		{200, PasswordWeight{Lower: 0, Upper: 0, Digit: 0, Symbol: 1}},
		{2000, PasswordWeight{Lower: 0, Upper: 0, Digit: 0, Symbol: 1}},
	}
	for _, c := range cases {
		got := Password(c.in, c.weight)
		pw_array := []rune(got)
		for i := 0; i < len(pw_array); i++ {
			if !strings.Contains(Symbols, string(pw_array[i])) {
				t.Errorf("Password(%q) containers something other than symbols: %q", c.in, got)
			}
		}
	}
}

func TestPasswordOnlyDigit(t *testing.T) {
	cases := []struct {
		in     int
		weight PasswordWeight
	}{
		{24, PasswordWeight{Lower: 0, Upper: 0, Digit: 1, Symbol: 0}},
		{3, PasswordWeight{Lower: 0, Upper: 0, Digit: 1, Symbol: 0}},
		{200, PasswordWeight{Lower: 0, Upper: 0, Digit: 1, Symbol: 0}},
		{2000, PasswordWeight{Lower: 0, Upper: 0, Digit: 1, Symbol: 0}},
	}
	for _, c := range cases {
		got := Password(c.in, c.weight)
		pw_array := []rune(got)
		for i := 0; i < len(pw_array); i++ {
			if !strings.Contains(Digits, string(pw_array[i])) {
				t.Errorf("Password(%q) containers something other than symbols: %q", c.in, got)
			}
		}
	}
}

func TestPasswordOnlyLower(t *testing.T) {
	cases := []struct {
		in     int
		weight PasswordWeight
	}{
		{24, PasswordWeight{Lower: 1, Upper: 0, Digit: 0, Symbol: 0}},
		{3, PasswordWeight{Lower: 1, Upper: 0, Digit: 0, Symbol: 0}},
		{200, PasswordWeight{Lower: 1, Upper: 0, Digit: 0, Symbol: 0}},
		{2000, PasswordWeight{Lower: 1, Upper: 0, Digit: 0, Symbol: 0}},
	}
	for _, c := range cases {
		got := Password(c.in, c.weight)
		pw_array := []rune(got)
		for i := 0; i < len(pw_array); i++ {
			if !strings.Contains(LowerLetters, string(pw_array[i])) {
				t.Errorf("Password(%q) containers something other than symbols: %q", c.in, got)
			}
		}
	}
}

func TestPasswordOnlyUpper(t *testing.T) {
	cases := []struct {
		in     int
		weight PasswordWeight
	}{
		{24, PasswordWeight{Lower: 0, Upper: 1, Digit: 0, Symbol: 0}},
		{3, PasswordWeight{Lower: 0, Upper: 1, Digit: 0, Symbol: 0}},
		{200, PasswordWeight{Lower: 0, Upper: 1, Digit: 0, Symbol: 0}},
		{2000, PasswordWeight{Lower: 0, Upper: 1, Digit: 0, Symbol: 0}},
	}
	for _, c := range cases {
		got := Password(c.in, c.weight)
		pw_array := []rune(got)
		for i := 0; i < len(pw_array); i++ {
			if !strings.Contains(UpperLetters, string(pw_array[i])) {
				t.Errorf("Password(%q) containers something other than symbols: %q", c.in, got)
			}
		}
	}
}
