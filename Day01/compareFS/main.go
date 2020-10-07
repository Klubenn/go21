package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type Trie struct {
	root	*Node
	name	string
}

type Node struct {
	prefix		rune
	parent		*Node
	children	map[rune]*Node
	length		int
	isWord		bool
	checked		bool
}

func newRoot(pref rune, name string) *Trie {
	return &Trie{
		root: &Node{
			prefix: pref,
			parent: nil,
			children: make(map[rune]*Node),
			length: 0,
			isWord: false,
			checked: false,
		},
		name: name,
	}
}

func addNode(current *Node, char rune) *Node {
	current.children[char] = &Node{
		prefix: char,
		parent: current,
		children: make(map[rune]*Node),
		length: current.length + 1,
		isWord: false,
		checked: false,
	}
	return current.children[char]
}

func (trie *Trie) AddPath(path string) {
	node := trie.root
	for _, char := range path {
		if n, ok := node.children[char]; ok {
			node = n
		} else {
			node = addNode(node, char)
		}
	}
	node.isWord = true
}

func parseFlags() (string, string) {
	oldFS := flag.String("old", "", "old fs dump file")
	newFS := flag.String("new", "", "new fs dump file")
	flag.Parse()
	if *oldFS == "" || *newFS == "" {
		fmt.Println("Both: old and new fs dumps must be present")
		os.Exit(0)
	}
	return *oldFS, *newFS
}

func parseFileIntoPrefixTrie(filePath string) *Trie {
	trie := newRoot(0, filePath)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Can't open the file, the following error occured: %v", err)
	}
	defer f.Close()
	read := bufio.NewScanner(f)
	for read.Scan() {
		trie.AddPath(read.Text())
	}
	return trie
}

func (trie *Trie) Exists(path string) bool {
	node := trie.root
	for _, char := range path {
		if n, ok := node.children[char]; ok {
			node = n
		} else {
			return false
		}
	}
	if node.isWord {
		node.checked = true
		return true
	}
	return false
}

func goThroughTrie(node *Node) {
	for _, child := range node.children {
		if child.isWord && !child.checked {
			printRemoved(child)
		}
		goThroughTrie(child)
	}
}

func printRemoved(node *Node) {
	name := make([]rune, node.length)
	for node.parent != nil {
		name[node.length - 1] = node.prefix
		node = node.parent
	}
	fmt.Println("REMOVED", string(name))
}

func compareFS(oldMap *Trie, newFS string) {
	f, err := os.Open(newFS)
	if err != nil {
		log.Fatalf("Can't open the file, the following error occured: %v", err)
	}
	defer f.Close()
	read := bufio.NewScanner(f)
	for read.Scan() {
		path := read.Text()
		if !oldMap.Exists(path) {
			fmt.Println("ADDED", path)
		}
	}
	goThroughTrie(oldMap.root)
}

func main() {
	oldFS, newFS := parseFlags()
	oldMap := parseFileIntoPrefixTrie(oldFS)
	compareFS(oldMap, newFS)
}
