package day19

const Overlapping = 12

type Scanner []Beacon
type Distances []int

func (s Scanner) Distances() []Distances {
	n := len(s)
	dist := make([]Distances, 0, n)
	for i := 0; i < n; i++ {
		dist = append(dist, make(Distances, n))
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist[i][j] = s[i].Distance(s[j])
			dist[j][i] = dist[i][j]
		}
	}
	return dist
}

func (a Scanner) Intersect(b Scanner) (Beacon, Rotate) {
	for x, j := range a.Distances() {
		for y, k := range b.Distances() {
			if m := j.Intersect(k); len(m) >= Overlapping {
				fromA, fromB := a[x], b[y]
				for k, v := range m {
					offA, offB := fromA.Offset(a[k]), fromB.Offset(b[v])
					rotate, err := findRotation(offA, offB)
					if err != nil {
						continue
					}

					offset := rotate(fromB).Offset(fromA)
					if len(a)+len(b)-count(a, b.Translate(offset, rotate)) >= Overlapping {
						return offset, rotate
					}
				}
			}
		}
	}
	return Beacon{}, nil
}

func (s Scanner) Translate(offset Beacon, rotate Rotate) Scanner {
	to := make(Scanner, 0, len(s))
	for _, beacon := range s {
		to = append(to, rotate(beacon).Move(offset))
	}
	return to
}

func (d Distances) toMap() map[int]int {
	values := map[int]int{}
	for k, v := range d {
		// Duplicates found, try with another node
		if _, found := values[v]; found {
			return map[int]int{}
		}
		values[v] = k
	}
	return values
}

func (a Distances) Intersect(b Distances) map[int]int {
	values, matches := a.toMap(), map[int]int{}
	for ref, v := range b {
		if id, found := values[v]; found {
			matches[id] = ref
		}
	}
	return matches
}
