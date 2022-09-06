package day21

type Fighter struct {
	HitPoints, Damage, Armor int
}

func (a Fighter) Wins(b Fighter) bool {
	var (
		attackA = max(a.Damage-b.Armor, 1)
		attackB = max(b.Damage-a.Armor, 1)
	)
	return ceil(b.HitPoints, attackA) <= ceil(a.HitPoints, attackB)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ceil(num, div int) int {
	return (num + div - 1) / div
}
