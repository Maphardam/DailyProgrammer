package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Trie struct {
	Children map[byte]Trie
}

func main() {

	// read words file and build trie
	words, err := os.Open("/usr/share/dict/words")
	if err != nil {
		log.Fatal("File not found.")
		return
	}
	defer words.Close()

	scanner := bufio.NewScanner(words)
	root := Trie{make(map[byte]Trie)}
	for scanner.Scan() {
		buildTrie(root, strings.ToLower(scanner.Text()))
	}

	// read input file and check contained words
	input, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("File not found.")
		return
	}
	defer input.Close()

	scanner = bufio.NewScanner(input)
	for scanner.Scan() {
		checkWord(root, strings.ToLower(scanner.Text()))
	}
}

func buildTrie(t Trie, w string) {
	if w == "" {
		return
	}
	if _, ok := t.Children[w[0]]; !ok {
		t.Children[w[0]] = Trie{make(map[byte]Trie)}
	}
	buildTrie(t.Children[w[0]], w[1:])
}

func checkWord(t Trie, w string) {
	if w == "" {
		fmt.Print("\n")
	}
	if _, ok := t.Children[w[0]]; !ok {
		fmt.Printf("%v<%v\n", string(w[0]), w[1:])
	} else {
		fmt.Print(string(w[0]))
		checkWord(t.Children[w[0]], w[1:])
	}
}
