package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	switch {
	case os.Args[0] == "./QuadA":
		QuadA(n1, n2)
	case os.Args[0] == "./QuadB":
		QuadB(n1, n2)
	case os.Args[0] == "./QuadC":
		QuadC(n1, n2)
	case os.Args[0] == "./QuadD":
		QuadD(n1, n2)
	case os.Args[0] == "./QuadE":
		QuadE(n1, n2)
	}
	// quadchecker()
}

func quadchecker() {
	FileNames := [5]string{"QuadA", "QuadB", "QuadC", "QuadD", "QuadE"}
	var x int
	var y int
	val, _ := io.ReadAll(os.Stdin)
	for _, i := range val {
		if i == '\n' {
			y++
		} else {
			x++
		}
	}
	if x == 0 || y == 0 {
		fmt.Println("Not a Raid function")
		return
	}
	x /= y
	twins := []string{}

	for _, i := range FileNames {
		res, _ := exec.Command("./"+i, strconv.Itoa(x), strconv.Itoa(y)).Output()
		if string(val) == string(res) {
			twins = append(twins, i)
		}
	}

	if len(twins) == 0 {
		fmt.Println("Not a Raid function")
		return
	}

	for i, j := range twins {
		if i == 0 {
			fmt.Printf("[%s] [%d] [%d]", j, x, y)
		} else {
			fmt.Printf(" || ")
			fmt.Printf("[%s] [%d] [%d]", j, x, y)
		}
	}
	fmt.Println()
}

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for h := 1; h <= y; h++ {
			for w := 1; w <= x; w++ {
				if (w == 1 && h == 1) || (w == x && h == 1) || (w == 1 && h == y) || (w == x && h == y) {
					z01.PrintRune('o')
				} else if w == 1 || w == x {
					z01.PrintRune('|')
				} else if h == 1 || h == y {
					z01.PrintRune('-')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for h := 1; h <= y; h++ {
			for w := 1; w <= x; w++ {
				if (w == 1 && h == 1) || (w == x && h == y && h != 1 && x != 1) {
					z01.PrintRune('/')
				} else if (w == x && h == 1) || (w == 1 && h == y) {
					z01.PrintRune('\\')
				} else if h == 1 || h == y || w == 1 || w == x {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		for h := 1; h <= y; h++ {
			for w := 1; w <= x; w++ {
				if (w == 1 && h == 1) || (w == x && h == 1) {
					z01.PrintRune('A')
				} else if (w == 1 && h == y) || (w == x && h == y && h != 1 && x != 1) {
					z01.PrintRune('C')
				} else if h == 1 || h == y || w == 1 || w == x {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		for h := 1; h <= y; h++ {
			for w := 1; w <= x; w++ {
				if (w == 1 && h == 1) || (w == 1 && h == y) {
					z01.PrintRune('A')
				} else if (w == x && h == 1) || (w == x && h == y && h != 1 && x != 1) {
					z01.PrintRune('C')
				} else if h == 1 || h == y || w == 1 || w == x {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

func QuadE(x, y int) {
	if x > 0 && y > 0 {
		for h := 1; h <= y; h++ {
			for w := 1; w <= x; w++ {
				if (w == 1 && h == 1) || (w == x && h == y && h != 1 && x != 1) {
					z01.PrintRune('A')
				} else if (w == x && h == 1) || (w == 1 && h == y) {
					z01.PrintRune('C')
				} else if h == 1 || h == y || w == 1 || w == x {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
