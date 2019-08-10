// The instructor of the linear algebra course decides that calculus is now a prerequisite. Extend the topoSort
// function to report cycles.
package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	// "linear algebra":        {"calculus"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	var route []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				for _, i := range route {
					if item == i {
						fmt.Println("circular reference found")
					}
				}
				route = append(route, item)
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	for key := range m {
		route = append(route, key)
		visitAll([]string{key})
		route = nil
	}
	return order
}
