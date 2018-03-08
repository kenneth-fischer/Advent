package main

import (
	"fmt"
	"github.com/kenneth-fischer/Advent/files"
	"log"
	"strconv"
	"strings"
)

func main() {
	var err error
	verify("b inc 5 if a > 1")
	verify("a inc 1 if b < 5")
	verify("c dec -10 if a >= 1")
	verify("c inc -20 if c == 10")

	instructions := read("b inc 5 if a > 1", "a inc 1 if b < 5", "c dec -10 if a >= 1", "c inc -20 if c == 10")
	m := execute(true, instructions...)
	register, value := m.Max()
	fmt.Printf("Max register %s = %d\n", register, value)
	fmt.Printf("Max register ever %s = %d", m.MaxRegister, m.MaxValue)
	fmt.Println()

	instructions, err  = readFile("instructions.txt")
	if err != nil {
		log.Fatal(err)
	}
	m = execute(false, instructions...)
	m.PrintRegisters()
	register, value = m.Max()
	fmt.Printf("Max register %s = %d\n", register, value)
	fmt.Printf("Max register ever %s = %d\n", m.MaxRegister, m.MaxValue)
}

func execute(debug bool, instructions ...instruction) machine {
	m := machine{}
	m.Execute(debug, instructions...)
	return m
}

type machine struct {
	Registers 	map[string]int
	MaxRegister	string
	MaxValue	int
}

func (m *machine) Execute(debug bool, instructions ... instruction) {
	m.MaxRegister = ""
	m.MaxValue = 0

	for _, instruction := range instructions {
		if m.Registers == nil {
			m.Registers = map[string]int{}
		}

		if instruction.Condition.isTrue(*m) {
			val := m.getRegister(instruction.Register)
			change := instruction.Change * instruction.Increment
			m.setRegister(instruction.Register, val + change)
			maxRegister, maxValue := m.Max()
			if m.MaxRegister == "" || maxValue > m.MaxValue {
				m.MaxValue = maxValue
				m.MaxRegister = maxRegister
			} 
		}
		if debug {
			m.PrintRegisters()
			fmt.Println("===================")
		}
	}
}

func (m machine) PrintRegisters() {
	for name, value := range m.Registers {
		fmt.Printf("%s = %d\n", name, value)
	}
}	 


func (m *machine) Max() (string, int) {
	maxKey := ""
	max := 0

	for key, value := range m.Registers {
		if maxKey == "" {
			max = value
			maxKey = key
		} else {
			if value > max {
				max = value
				maxKey = key 
			}
		}
	}
	return maxKey, max
}
func (m *machine) getRegister(name string) int {
	if value, ok := m.Registers[name]; ok {
		return value
	}
	m.Registers[name] = 0
	return 0
}

func (m *machine) setRegister(name string, value int) {
	m.Registers[name] = value
}

type register struct {
	Name string
	Value int
}

type condition struct {
	Register string
	Operator string
	Value	 int
}

func (cond condition) isTrue(m machine) bool {
	regValue := m.getRegister(cond.Register)
	switch cond.Operator {
	case "<":
		return regValue < cond.Value
	case ">":
		return regValue > cond.Value
	case "<=":
		return regValue <= cond.Value
	case ">=":
		return regValue >= cond.Value
	case "!=":
		return regValue != cond.Value
	case "==":
		return regValue == cond.Value
	}
	log.Panicf("Comparison %q contains an invalid operator %q", cond.String(), cond.Operator)
	return false
}

func (cond condition) String() string {
	return fmt.Sprintf("%s %s %d", cond.Register, cond.Operator, cond.Value)
}

type instruction struct {
	Register 	string
	Increment	int
	Change	 	int
	Condition	condition
}

func (instr instruction) String() string {
	increment := "inc"
	if instr.Increment < 0 {
		increment = "dec"
	}
	return fmt.Sprintf("%s %s %d if %s", instr.Register, increment, instr.Change, instr.Condition.String())
}

func readFile(path string) ([]instruction, error) {
	lines, err := files.ReadLines(path)

	if err != nil {
		return []instruction{}, err
	}
	return read(lines...), nil
} 
 	
func read(instructions ... string) []instruction {
	results := []instruction{}

	for _, text := range instructions {
		var register1, incr, register2, operator string
		var change, value int
		parts := strings.Split(text, " ")

		if len(parts) > 0 {
			register1 = parts[0]
		}
		if len(parts) > 1 {
			incr = parts[1]
		}
		if len(parts) > 2 {
			change, _ = strconv.Atoi(parts[2])
		}
		if len(parts) > 4 {
			register2 = parts[4]
		}
		if len(parts) > 5 {
			operator = parts[5]
		}
		if len(parts) > 6 {
			value, _ = strconv.Atoi(parts[6])
		}
		instr := instruction{}
		instr.Register = register1
		instr.Change = change
		if incr == "inc" {
			instr.Increment = 1
		} else {
			instr.Increment = -1
		}

		instr.Condition = condition{}
		instr.Condition.Register = register2
		instr.Condition.Operator = operator
		instr.Condition.Value = value
		results = append(results, instr)
	}
	return results
}

func verify(text string) {
	instructions := read(text)
	if instructions[0].String() != text {
		log.Fatalf("Expected %q. Got %q", text, instructions[0].String())
	}
}
