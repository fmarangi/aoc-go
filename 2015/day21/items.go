package day21

type Item struct {
	Cost, Damage, Armor int
}

type Battle []Item

func (b Battle) Cost() (cost int) {
	for _, i := range b {
		cost += i.Cost
	}
	return
}

func (b Battle) Fighter() (f Fighter) {
	f.HitPoints = 100
	for _, i := range b {
		f.Damage += i.Damage
		f.Armor += i.Armor
	}
	return
}

var weapons = []Item{
	{8, 4, 0}, {10, 5, 0}, {25, 6, 0},
	{40, 7, 0}, {74, 8, 0},
}

var armor = []Item{
	{}, {13, 0, 1}, {31, 0, 2},
	{53, 0, 3}, {75, 0, 4}, {102, 0, 5},
}

var rings = []Item{
	{}, {25, 1, 0}, {50, 2, 0}, {100, 3, 0},
	{20, 0, 1}, {40, 0, 2}, {80, 0, 3},
}

func Battles() <-chan Battle {
	out := make(chan Battle)
	go func() {
		defer close(out)
		for _, w := range weapons {
			for _, a := range armor {
				out <- Battle{w, a} // No rings
				for i, n := 0, len(rings); i < n; i++ {
					for j := i + 1; j < n; j++ {
						out <- Battle{w, a, rings[i], rings[j]}
					}
				}
			}
		}
	}()
	return out
}
