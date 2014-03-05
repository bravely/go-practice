package main

import (
  "fmt"
  "strings"
)

func WordCount(s string) map[string]int {
  s_array := strings.Fields(s)
  counts := make(map[string]int)

  for i := 0; i < len(s_array); i++ {
    counts[s_array[i]]++
  }

  return counts
}

func main() {
  fmt.Println(WordCount("Hi how are how"))
}
