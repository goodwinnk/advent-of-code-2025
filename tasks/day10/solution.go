package day10

import (
	"AdventOfCode2025/internal/util"
	"AdventOfCode2025/internal/util/coll"
	"bufio"
	"fmt"
	"strings"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize/convex/lp"
)

const day = 10

type TPart1 = int
type TPart2 = TPart1

func Part1() (TPart1, error) {
	return Part1Text(util.Input(day))
}

func Part2() (TPart2, error) {
	return Part2Text(util.Input(day))
}

func (m *Machine) solve() (int, error) {
	known := make(map[int]int)
	known[0] = 0
	current := []int{0}

	for len(current) > 0 {
		next := make([]int, 0)
		for _, state := range current {
			for _, b := range m.buttons {
				updated := state ^ b
				_, ok := known[updated]
				if !ok {
					known[updated] = known[state] + 1
					if m.on == updated {
						return known[updated], nil
					}

					next = append(next, updated)
				}
			}
		}
		current = next
	}

	return -1, fmt.Errorf("couldn't needed buttons")
}

func Part1Text(input string) (TPart1, error) {
	machines, err := parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	r := 0
	for _, m := range machines {
		machine := NewMachine(&m)

		minPress, err := machine.solve()
		if err != nil {
			return 0, fmt.Errorf("couldn't solve machine: %v %w", machine, err)
		}

		r += minPress
	}

	return r, nil
}

func Part2Text(input string) (TPart2, error) {
	machines, err := parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	for _, m := range machines {
		fmt.Println(m)

		c := coll.NewSlice[float64](len(m.buttons), 1)

		problem := make([][]int, len(m.joltage))
		for j := 0; j < len(m.joltage); j++ {
			problem[j] = make([]int, len(m.buttons)+1)
			problem[j][len(m.buttons)] = m.joltage[j]
		}
		for i, b := range m.buttons {
			for _, br := range b {
				problem[br][i] = 1
			}
		}
		for _, r := range problem {
			fmt.Println(r)
		}

		AData := make([]float64, len(m.buttons)*len(m.joltage))
		for i := 0; i < len(m.joltage); i++ {
			for b := 0; b < len(m.buttons); b++ {
				AData[i*len(m.buttons)+b] = float64(problem[i][b])
			}
		}
		A := mat.NewDense(len(m.joltage), len(m.buttons), AData)
		fmt.Println("A", A)

		b := make([]float64, len(m.joltage))
		for i, jol := range m.joltage {
			b[i] = float64(jol)
		}
		fmt.Println("b", b)

		opt, x, err := lp.Simplex(c, A, b, 0, nil)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("opt: %v\n", opt)
			fmt.Printf("x: %v\n", x)
		}
	}

	return 0, nil
}

type Machine struct {
	total   int
	on      int
	buttons []int
	joltage []int
}

func (m Machine) String() string {
	on := fmt.Sprintf("%0*b", m.total, m.on)
	binButtons := make([]string, len(m.buttons))
	for i, b := range m.buttons {
		binButtons[i] = fmt.Sprintf("%0*b", m.total, b)
	}

	return fmt.Sprintf("[%v] %v %v", on, binButtons, m.joltage)
}

type MachineRaw struct {
	on      string
	buttons [][]int
	joltage []int
}

func NewMachine(raw *MachineRaw) Machine {
	onN := 0
	for i := 0; i < len(raw.on); i++ {
		onN = onN << 1
		if raw.on[i] == '#' {
			onN++
		}
	}

	buttons := make([]int, len(raw.buttons))
	for b := range raw.buttons {
		bN := 0
		for _, bri := range raw.buttons[b] {
			bN = bN | (1 << (len(raw.on) - 1 - bri))
		}
		buttons[b] = bN
	}

	return Machine{total: len(raw.on), on: onN, buttons: buttons, joltage: raw.joltage}
}

func (m MachineRaw) String() string {
	return fmt.Sprintf("[%s] %v %v", m.on, m.buttons, m.joltage)
}

func parse(input string) ([]MachineRaw, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var result []MachineRaw

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// on section
		iL := strings.IndexByte(line, '[')
		iR := strings.IndexByte(line, ']')
		if iL == -1 || iR == -1 || iL > iR {
			return nil, fmt.Errorf("invalid line (on): %q", line)
		}
		on := strings.TrimSpace(line[iL+1 : iR])

		// joltage section
		jL := strings.LastIndexByte(line, '{')
		jR := strings.LastIndexByte(line, '}')
		if jL == -1 || jR == -1 || jL > jR || jL < iR {
			return nil, fmt.Errorf("invalid line (joltage): %q", line)
		}
		joltage, err := util.ParseInts(line[jL+1 : jR])
		if err != nil {
			return nil, fmt.Errorf("invalid joltage: %w", err)
		}

		// buttons section(s) exist in the middle between ']' and '{'
		middle := line[iR+1 : jL]
		var buttons [][]int
		for {
			bL := strings.IndexByte(middle, '(')
			if bL == -1 {
				break
			}
			bR := strings.IndexByte(middle[bL+1:], ')')
			if bR == -1 {
				return nil, fmt.Errorf("unclosed buttons group: %q", line)
			}
			bR += bL + 1

			group := middle[bL+1 : bR]
			ints, err := util.ParseInts(group)
			if err != nil {
				return nil, fmt.Errorf("invalid buttons group %q: %w", group, err)
			}
			buttons = append(buttons, ints)

			middle = middle[bR+1:]
		}

		result = append(result, MachineRaw{on: on, buttons: buttons, joltage: joltage})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
