package day07

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type File struct {
	Size int
	Name string
}

type Directory struct {
	Parent   *Directory
	Children map[string]*Directory
	Files    []File
}

func (d Directory) Size() (size int) {
	for _, f := range d.Files {
		size += f.Size
	}
	for _, c := range d.Children {
		size += c.Size()
	}
	return
}

func Part1(input string) (totalSize int) {
	for _, d := range parseInput(input) {
		if size := d.Size(); size <= 100000 {
			totalSize += size
		}
	}
	return
}

func Part2(input string) (best int) {
	var (
		dirs   = parseInput(input)
		total  = dirs[0].Size()
		target = total - 70000000 + 30000000
	)

	best = total
	for _, d := range dirs[1:] {
		if size := d.Size(); size > target && size < best {
			best = size
		}
	}

	return
}

func parseInput(input string) (dirs []*Directory) {
	var (
		current = new(Directory)
		root    = current
		data    = strings.Split(input, "\n")
		n       = len(data)
	)

	dirs = append(dirs, current)
	current.Children = make(map[string]*Directory)

	for i := 0; i < n; i++ {
		switch data[i] {
		case "$ cd /":
			current = root

		case "$ ls":
			for ; i < n-1 && !strings.HasPrefix(data[i+1], "$ "); i++ {
				switch {
				case strings.HasPrefix(data[i+1], "dir "):
					child := new(Directory)
					child.Children = make(map[string]*Directory)
					child.Parent = current
					current.Children[data[i+1][4:]] = child
					dirs = append(dirs, child)

				default:
					current.Files = append(current.Files, parseFile(data[i+1]))
				}
			}

		case "$ cd ..":
			current = current.Parent

		default:
			current = current.Children[data[i][5:]]
		}
	}
	return
}

func parseFile(input string) File {
	parts := strings.Fields(input)
	return File{utils.ToInt(parts[0]), parts[1]}
}
