package day24

import (
	"regexp"
	"strings"
)

type Route struct{ Pos, Steps int }
type Segment struct{ To, From string }
type Distances map[Segment]int
type Maze string

func (m Maze) width() int {
	return strings.Index(string(m), "\n") + 1
}

func (m Maze) POI() map[string]int {
	var (
		pattern = regexp.MustCompile(`[0-9]`)
		poi     = pattern.FindAllStringIndex(string(m), -1)
		res     = make(map[string]int, len(poi))
	)

	for _, p := range poi {
		res[string(m[p[0]])] = p[0]
	}

	return res
}

func (m Maze) Explore(start int) <-chan Route {
	var (
		out     = make(chan Route)
		w       = m.width()
		offsets = [4]int{-1, 1, -w, w}
		seen    = map[int]bool{start: true}
	)

	go func() {
		defer close(out)
		for q := []Route{{start, 0}}; len(q) > 0; q = q[1:] {
			curr := q[0]
			for _, o := range offsets {
				next := Route{curr.Pos + o, curr.Steps + 1}
				what := m[next.Pos]

				switch {
				case what == '.' && !seen[next.Pos]:
					q = append(q, next)
					seen[next.Pos] = true
				case what == '#', seen[next.Pos]:
					// Do nothing
				default:
					out <- next
					q = append(q, next)
					seen[next.Pos] = true
				}
			}
		}
	}()

	return out
}

func (m Maze) Distances() Distances {
	var (
		poi       = m.POI()
		count     = len(poi)
		distances = make(Distances, count*(count-1))
	)

	for from, pos := range poi {
		for to := range m.Explore(pos) {
			seg := Segment{from, string(m[to.Pos])}
			distances[seg] = to.Steps
		}
	}

	return distances
}

func (d Distances) Length(route []string) (length int) {
	for i, n := 1, len(route); i < n; i++ {
		length += d[Segment{route[i-1], route[i]}]
	}
	return
}
