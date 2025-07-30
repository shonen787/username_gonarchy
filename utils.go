package main

import (
	"os"
	"strings"
)

func upcaseIt(s string) string{
  if s == "" {
    return ""
  }
  return strings.ToUpper(string(s[0])) + strings.ToLower(string(s[1:]))
}

func saveOutput(usernames chan string) {
  file, err := os.Create("./output.txt")
  check(err)
  defer file.Close()

  for user := range usernames{
    file.WriteString(user + "\n")
  }
file.Sync()
  
}

