package day25

import (
	"fmt"
	"strings"
)

type Cursor struct {
	State string
	Value bool
}

type Rule struct {
	Value     bool
	NextState string
	Offset    int
}

type TuringMachine map[Cursor]Rule

type Tape []bool

func (t Tape) Checksum() (value int) {
	for _, v := range t {
		if v {
			value++
		}
	}
	return
}

func Part1(input string) (checksum int) {
	steps, computer := parseInput(input)
	tape := make(Tape, steps)

	pos, cursor := steps/2, Cursor{State: "A"}
	for i := 0; i < steps; i++ {
		cursor.Value = tape[pos]
		rule := computer[cursor]
		tape[pos] = rule.Value
		pos += rule.Offset
		cursor.State = rule.NextState
	}

	return tape.Checksum()
}

func parseInput(input string) (steps int, computer TuringMachine) {
	computer = make(TuringMachine)
	parts := strings.Split(input, "\n\n")

	fmt.Sscanf(strings.Split(parts[0], "\n")[1],
		"Perform a diagnostic checksum after %d steps.", &steps)

	for _, p := range parts[1:] {
		rows := strings.Split(p, "\n")

		var state string
		fmt.Sscanf(rows[0], "In state %1s:", &state)
		computer[Cursor{state, false}] = parseRule(rows[2:5])
		computer[Cursor{state, true}] = parseRule(rows[6:9])
	}

	return
}

func parseRule(rows []string) (rule Rule) {
	rule.Value = strings.HasSuffix(rows[0], "1.")
	rule.Offset = 1
	if strings.HasSuffix(rows[1], "left.") {
		rule.Offset = -1
	}
	rule.NextState = rows[2][26:27]
	return
}
