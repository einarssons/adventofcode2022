package main

import (
	"fmt"
	"strings"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

func task1(lines []string) int {
	m := parse(lines)
	sum := 0
	for _, v := range m {
		size := v.Size(m)
		if size < 100000 {
			sum += size
		}
	}
	return sum
}

func task2(lines []string) int {
	m := parse(lines)
	var bigDirs []int
	totalSize := m["/"].Size(m)
	maxSize := 40000000
	toRemove := totalSize - maxSize
	for _, v := range m {
		size := v.Size(m)
		if size > toRemove {
			bigDirs = append(bigDirs, size)
		}
	}
	return u.Min(bigDirs)
}

type node struct {
	subDirs []string
	leafs   []leaf
}

type leaf struct {
	name string
	size int
}

func (n *node) Size(m map[string]*node) int {
	sum := 0
	for _, dir := range n.subDirs {
		sub, ok := m[dir]
		if !ok {
			fmt.Printf("dir %s does not exist\n", dir)
		}
		sum += sub.Size(m)
	}
	for _, leaf := range n.leafs {
		sum += leaf.size
	}
	return sum
}

func parse(lines []string) map[string]*node {
	path := ""
	m := make(map[string]*node)
	var n *node
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "$ cd /"):
			path = "/"
			n := &node{}
			m[path] = n
		case strings.HasPrefix(line, "$ cd"):
			parts := strings.Split(line, " ")
			d := parts[len(parts)-1]
			switch d {
			case "..":
				pathParts := strings.Split(path, "/")
				path = strings.Join(pathParts[:len(pathParts)-2], "/") + "/"
				n = m[path]
			default:
				path += d + "/"
				m[path] = &node{}
			}
		case strings.HasPrefix(line, "$ ls"):
			n = m[path]
		default:
			parts := strings.Split(line, " ")
			if parts[0] == "dir" {
				p := path + parts[1] + "/"
				n.subDirs = append(n.subDirs, p)
			} else {
				l := leaf{parts[1], u.Atoi(parts[0])}
				n.leafs = append(n.leafs, l)
			}
		}
	}
	return m
}
