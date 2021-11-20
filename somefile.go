package mygreeting

import "fmt"

# Hello function
func Hello(name string) string {
  message := fmt.Sprintf("Hi, %v.  Welcome to go!", name)
  return message
}
