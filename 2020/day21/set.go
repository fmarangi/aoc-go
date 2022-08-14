package day21

type Set map[string]bool

func (set *Set) Add(value string) {
	(*set)[value] = true
}

func (set *Set) AddAll(values []string) {
	for _, value := range values {
		set.Add(value)
	}
}

func Intersect(a, b Set) Set {
	intersection := make(Set)
	if len(a) > len(b) {
		a, b = b, a
	}

	for c := range b {
		if a[c] {
			intersection.Add(c)
		}
	}
	return intersection
}
