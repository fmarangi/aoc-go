package day04

import (
	"regexp"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

var rules = map[string]validation{
	"byr": between(1920, 2002),
	"iyr": between(2010, 2020),
	"eyr": between(2020, 2030),
	"hgt": func(s string) bool {
		var valid validation
		if strings.HasSuffix(s, "cm") {
			valid = between(150, 193)
		} else {
			valid = between(59, 76)
		}
		return s != "" && valid(s[:len(s)-2])
	},
	"hcl": re(`#[0-9a-f]{6}`),
	"ecl": re(`(amb|blu|brn|gry|grn|hzl|oth)`),
	"pid": re(`\d{9}`),
}

type Passport map[string]string

func (p Passport) HasAllFields() bool {
	for field := range rules {
		if _, ok := p[field]; !ok {
			return false
		}
	}
	return true
}

func (p Passport) Valid() bool {
	for k, v := range rules {
		if !v(p[k]) {
			return false
		}
	}
	return true
}

type validation func(string) bool

func between(from, to int) validation {
	return func(s string) bool {
		value := utils.ToInt(s)
		return value >= from && value <= to
	}
}

func re(pattern string) validation {
	p := regexp.MustCompile("^" + pattern + "$")
	return func(s string) bool { return p.MatchString(s) }
}

func Part1(input string) (valid int) {
	for _, passport := range parseInput(input) {
		if passport.HasAllFields() {
			valid++
		}
	}
	return
}

func Part2(input string) (valid int) {
	for _, passport := range parseInput(input) {
		if passport.Valid() {
			valid++
		}
	}
	return
}

func parseInput(input string) (passports []Passport) {
	for _, data := range strings.Split(input, "\n\n") {
		passport := make(Passport)
		for _, field := range strings.Fields(data) {
			separator := strings.Index(field, ":")
			passport[field[:separator]] = field[separator+1:]
		}
		passports = append(passports, passport)
	}
	return
}
