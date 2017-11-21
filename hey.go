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
    last := nodes.last()
    if last.level<level { break }
    if last.tag=="|" || last.tag=="/" || last.tag[0]=='<' {
      // No closing
    } else {
      fmt.Print("</", last.tag, ">");
    }
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

func extractTag(line string) string {
  var tag string
  i := strings.Index(line, " ")
  if i>-1 {
    tag = line[:i]
  } else {
    tag = line[:]
  }
  return tag
}

func isCommented(nodes nodeSlice, level int) bool {
  if len(nodes)<1 { return false }
  n := nodes.last()
  return n.tag=="/" && level>n.level
}

func main() {

  nodes := make(nodeSlice, 0, 16)

	fileScanner := bufio.NewScanner(os.Stdin)

	for fileScanner.Scan() {

    line := fileScanner.Text()
    level := indentLevel(line)

    if isCommented(nodes, level) { continue }

    line = line[level:]

    if len(line)>0 {
      
      tag := extractTag(line)

      nodes = closeNodes(nodes, level)
      nodes = append(nodes, &node{level,tag})
      if tag=="|" {
        fmt.Print(line[2:])
      } else if tag=="/" {
        // Comment
      } else if tag[0]=='<' {
        fmt.Print(line)
      } else {
        fmt.Print("<",line,">")
      }
    }
	}

  closeNodes(nodes, 0)
  fmt.Println("")
}

