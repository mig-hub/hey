package main

import (
	"fmt"
	"bufio"
	"os"
  "strings"
)

type node struct {
  level int
  tag string
}

type nodeSlice []*node

func (nodes nodeSlice) last() *node {
  return nodes[len(nodes)-1]
}

func closeNodes(nodes nodeSlice, level int) nodeSlice {
  for len(nodes)>0 {
    last := nodes[len(nodes)-1]
    if last.level<level { break }
    fmt.Print("</", last.tag, ">");
    nodes = nodes[:len(nodes)-1]
  }
  return nodes
}

func indentLevel(line string) int {
  i := 0
  for _, runeValue := range line {
    if runeValue==' ' || runeValue=='\t' {
      i++
    } else {
      break
    }
  }
  return i
}

func main() {

  line := ""
  level := 0
  tag := ""
  nodes := make(nodeSlice, 0, 16)

	fileScanner := bufio.NewScanner(os.Stdin)

	for fileScanner.Scan() {
    line = fileScanner.Text()
		level = indentLevel(line)
    line = line[level:]
    if len(line)>0 {
      i := strings.Index(line, " ")
      if i>-1 {
        tag = line[:i]
      } else {
        tag = line[:]
      }
      nodes = closeNodes(nodes, level)
      nodes = append(nodes, &node{level,tag})
      fmt.Print("<",line,">")
    }
	}

  closeNodes(nodes, 0)
  fmt.Println("")
}

