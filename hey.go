package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
  level int
  tag string
  previous *node
}

func closeNodes(lastNode *node, level int) *node {
  for lastNode!=nil && level<=lastNode.level {
    fmt.Print("</", lastNode.tag, ">");
    lastNode = lastNode.previous
  }
  return lastNode;
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

	fileScanner := bufio.NewScanner(os.Stdin)

  line := ""
  level := 0
  tag := ""
  var lastNode *node

	for fileScanner.Scan() {
    line = fileScanner.Text()
		level = indentLevel(line)
    tag = line[level:]
    if len(tag)>0 {
      lastNode = closeNodes(lastNode, level)
      lastNode = &node{level,tag,lastNode}
      fmt.Print("<",tag,">")
    }
	}

  closeNodes(lastNode, 0)
  fmt.Println("")
}

