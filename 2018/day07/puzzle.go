package day07

import (
	"fmt"
	"sort"
	"strings"
)

type Steps map[string][]string

func (steps Steps) Next(finished []string) (next []string) {
	values := make(map[string]bool, len(finished))
	for _, f := range finished {
		values[f] = true
	}

next:
	for step, deps := range steps {
		if values[step] {
			continue
		}

		for _, d := range deps {
			if !values[d] {
				continue next
			}
		}

		next = append(next, step)
	}
	sort.Strings(next)
	return
}

type Worker struct {
	Step    string
	Seconds int
}

func (w *Worker) forStep(step string) {
	value := 60 + int(step[0]-'A')
	w.Seconds = value
	w.Step = step
}

func Part1(input string) string {
	steps := parseInput(input)
	finished := make([]string, 0, len(steps))
	for len(finished) < len(steps) {
		finished = append(finished, steps.Next(finished)[0])
	}
	return strings.Join(finished, "")
}

func Part2(input string) (seconds int) {
	steps, workers := parseInput(input), [5]Worker{}
	s, w := len(steps), len(workers)
	finished, pending := make([]string, 0, s), make(map[string]bool, s)

	for next := steps.Next(finished); len(finished) < s; seconds++ {
		update := false
		for i := 0; i < w; i++ {
			if workers[i].Seconds > 0 {
				workers[i].Seconds--
				continue
			}

			if workers[i].Step != "" {
				finished = append(finished, workers[i].Step)
				delete(pending, workers[i].Step)
				workers[i].Step = ""
				update = true
			}
		}

		if update {
			next = []string{}
			for _, p := range steps.Next(finished) {
				if !pending[p] {
					next = append(next, p)
				}
			}
		}

		for i := 0; i < w && len(next) > 0; i++ {
			if workers[i].Step == "" {
				workers[i].forStep(next[0])
				pending[next[0]] = true
				next = next[1:]
			}
		}
	}
	return seconds - 1
}

func parseInput(input string) Steps {
	steps := Steps{}
	for _, row := range strings.Split(input, "\n") {
		s := [2]string{}
		fmt.Sscanf(row, "Step %s must be finished before step %s can begin.", &s[0], &s[1])
		for _, step := range s {
			if _, found := steps[step]; !found {
				steps[step] = []string{}
			}
		}
		steps[s[1]] = append(steps[s[1]], s[0])
	}
	return steps
}
