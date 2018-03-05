package main

import (
	"fmt"
	"github.com/kenneth-fischer/Advent/files"
	"log"
	"strconv"
	"strings"
)

var (
	towers []*tower
	towersByName map[string]*tower
)

type tower struct {
	Name string
	Weight int
	totalWeight *int
	Subtowers []string
	Parent string
}

func (t tower) String() string {
	tStr := fmt.Sprintf("%s (%d)", t.Name, t.Weight)
	if len(t.Subtowers) > 0 {
		tStr = fmt.Sprintf("%s -> %s", tStr, t.Subtowers[0])
		for i := 1; i < len(t.Subtowers); i++ {
			tStr = fmt.Sprintf("%s, %s", tStr, t.Subtowers[i])
		}
	}
	return tStr
}

func (t tower) GetSubtowers() []*tower {
	subtowers := []*tower{}
	for _, name := range t.Subtowers {
		if subtower, ok := towersByName[name]; ok {
			subtowers = append(subtowers, subtower)
		}
	}
	return subtowers
}

func (t tower) TotalWeight() int {
	if t.totalWeight != nil {
		return *t.totalWeight
	}
	total := t.Weight
	for _, subtower := range t.GetSubtowers() {
		total += subtower.TotalWeight()
	}
	t.totalWeight = &total
	return *t.totalWeight
}
		

func main() {
	resetTowers()
	loadTower("pbga (66)")
	loadTower("xhth (57)")
	loadTower("ebii (61)")
	loadTower("havc (66)")
	loadTower("ktlj (57)")
	loadTower("fwft (72) -> ktlj, cntj, xhth")
	loadTower("qoyq (66)")
	loadTower("padx (45) -> pbga, havc, qoyq")
	loadTower("tknk (41) -> ugml, padx, fwft")
	loadTower("jptl (61)")
	loadTower("ugml (68) -> gyxo, ebii, jptl")
	loadTower("gyxo (61)")
	loadTower("cntj (57)")
	linkTowers()
	root := findRoot()
	fmt.Printf("root: %s\n", root.Name)
	isBalanced(*root)
	fmt.Println()
	
	resetTowers()
	loadTowers("towers.txt")
	linkTowers()
	root = findRoot()
	fmt.Printf("root: %s\n", root.Name)
	isBalanced(*root)
}

func isBalanced(t tower) bool {
	for _, subtower := range t.GetSubtowers() {
		if !isBalanced(*subtower) {
			return false
		}
	}

	balanced := []*tower{}
	unbalanced := []*tower{}
	
	for _, subtower := range t.GetSubtowers() {
		// Check weight against expected weight, if there is one.
		if len(balanced) > 0 {
			if subtower.TotalWeight() == balanced[0].TotalWeight() {
				balanced = append(balanced, subtower)
			} else {
				unbalanced = append(unbalanced, subtower)
			}
			continue
		}

		// Check weight against weight of unmatched subtowers
		matched := false
		for i := 0; i < len(unbalanced); i++ {
			unmatched := unbalanced[i]
			if unmatched.TotalWeight() == subtower.TotalWeight() {
				balanced = append(balanced, subtower, unmatched)
				unbalanced = append(unbalanced[:i], unbalanced[i+1:]...)
				matched = true
				break
			}
		}

		if !matched {
			unbalanced = append(unbalanced, subtower)
		}
	}
	if len(unbalanced) > 1 {
		fmt.Printf("%s is unbalanced. %s\n", t.Name, weightsDoNotMatch(unbalanced, balanced))
		return false
	}
	if len(unbalanced) > 0 && len(balanced) > 0 {
		fmt.Printf("%s is unbalanced. %s\n", t.Name, weightsDoNotMatch(unbalanced, balanced))
		return false
	}
	return true
}

func weightsDoNotMatch(wrong []*tower, right []*tower) string {
	if len(right) > 0 {
		msg := fmt.Sprintf("Weight of %s (%d/%d) does not match %s (%d)", wrong[0].Name, wrong[0].Weight, wrong[0].TotalWeight(), right[0].Name, right[0].TotalWeight())
		for i := 1; i < len(right); i++ {
			msg = fmt.Sprintf("%s, %s (%d)", msg, right[i].Name, right[i].TotalWeight())
		}
		return msg
	}
	if len(wrong) > 1 {
		msg := fmt.Sprintf("Weights do not match: %s (%d/%d)", wrong[0].Name, wrong[0].Weight, wrong[0].TotalWeight())
		for i := 1; i < len(right); i++ {
			msg = fmt.Sprintf("%s, %s (%d/%d)", msg, wrong[i].Name, wrong[i].Weight, wrong[i].TotalWeight())
		}
		return msg
	}
	return ""
}
	
func findRoot() *tower {
	for _, tower := range towers {
		if tower.Parent == "" {
			return tower
		}
	}
	return nil
}

func resetTowers() {
	towers = []*tower{}
	towersByName = map[string]*tower{}
}

func loadTowers(path string) {
	lines, err := files.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		loadTower(line)
	}
	linkTowers()
}

func loadTower(towerInfo string) {
	tower := tower{}
	parts := strings.Split(towerInfo, " -> ")
	nameAndWeight := strings.Split(parts[0], " ")
	tower.Name = nameAndWeight[0]
	weightStr := nameAndWeight[1]
	weightStr = string(weightStr[1:len(weightStr) - 1])
	tower.Weight, _ = strconv.Atoi(weightStr)
	if len(parts) > 1 {
		tower.Subtowers = strings.Split(parts[1], ", ")
	} else {
		tower.Subtowers = []string{}
	}
	towers = append(towers, &tower)
	towersByName[tower.Name] = &tower
} 

func linkTowers() {
	for _, tower := range towers {
		for _, name := range tower.Subtowers {
			towersByName[name].Parent = tower.Name
		}
	}
}
