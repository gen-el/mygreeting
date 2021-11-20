package mygreeting

import "fmt"

func Hello(name string) string {
  message := fmt.Sprintf("Hi, %v.  Welcome to go!", name)
  return message
}
