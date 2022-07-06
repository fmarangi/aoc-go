package day16

type Opcode func(registers *[4]int, values [4]int)

var opcodes = [16]Opcode{
	// add
	func(registers *[4]int, values [4]int) {
		(*registers)[values[3]] = (*registers)[values[1]] + (*registers)[values[2]]
	},
	func(registers *[4]int, values [4]int) {
		(*registers)[values[3]] = (*registers)[values[1]] + values[2]
	},

	// mul
	func(registers *[4]int, values [4]int) {
		(*registers)[values[3]] = (*registers)[values[1]] * (*registers)[values[2]]
	},
	func(registers *[4]int, values [4]int) {
		(*registers)[values[3]] = (*registers)[values[1]] * values[2]
	},

	// and
	func(registers *[4]int, values [4]int) {
		(*registers)[values[3]] = (*registers)[values[1]] & (*registers)[values[2]]
	},
	func(registers *[4]int, values [4]int) {
		(*registers)[values[3]] = (*registers)[values[1]] & values[2]
	},

	// or
	func(registers *[4]int, values [4]int) {
		(*registers)[values[3]] = (*registers)[values[1]] | (*registers)[values[2]]
	},
	func(registers *[4]int, values [4]int) {
		(*registers)[values[3]] = (*registers)[values[1]] | values[2]
	},

	// assign
	func(registers *[4]int, values [4]int) {
		(*registers)[values[3]] = (*registers)[values[1]]
	},
	func(registers *[4]int, values [4]int) {
		(*registers)[values[3]] = values[1]
	},

	// greater than
	func(registers *[4]int, values [4]int) {
		if values[1] > (*registers)[values[2]] {
			(*registers)[values[3]] = 1
		} else {
			(*registers)[values[3]] = 0
		}
	},
	func(registers *[4]int, values [4]int) {
		if (*registers)[values[1]] > values[2] {
			(*registers)[values[3]] = 1
		} else {
			(*registers)[values[3]] = 0
		}
	},
	func(registers *[4]int, values [4]int) {
		if (*registers)[values[1]] > (*registers)[values[2]] {
			(*registers)[values[3]] = 1
		} else {
			(*registers)[values[3]] = 0
		}
	},

	// equality
	func(registers *[4]int, values [4]int) {
		if values[1] == (*registers)[values[2]] {
			(*registers)[values[3]] = 1
		} else {
			(*registers)[values[3]] = 0
		}
	},
	func(registers *[4]int, values [4]int) {
		if (*registers)[values[1]] == values[2] {
			(*registers)[values[3]] = 1
		} else {
			(*registers)[values[3]] = 0
		}
	},
	func(registers *[4]int, values [4]int) {
		if (*registers)[values[1]] == (*registers)[values[2]] {
			(*registers)[values[3]] = 1
		} else {
			(*registers)[values[3]] = 0
		}
	},
}

func OpcodeNumbers(samples []Sample) map[int]int {
	var (
		test     [4]int
		matches  [16][16]bool
		numbers  [16][]int
		assigned = make(map[int]int, 16)
	)

	for _, s := range samples {
		for i, op := range opcodes {
			copy(test[:], s.before[:])
			op(&test, s.values)
			if test == s.after {
				matches[i][s.values[0]] = true
			}
		}
	}

	for i := 0; i < 256; i++ {
		if j := i / 16; matches[j][i%16] {
			numbers[j] = append(numbers[j], i%16)
		}
	}

	for len(assigned) < 16 {
		for i, num := range numbers {
			open := []int{}
			for _, n := range num {
				if _, found := assigned[n]; !found {
					open = append(open, n)
				}
			}
			if len(open) == 1 {
				assigned[open[0]] = i
			}
		}
	}

	return assigned
}
