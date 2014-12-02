package colors

import "fmt"

const (
  red          string = "\x1b[91m"
  cyan         string = "\x1b[36m"
  green        string = "\x1b[32m"
  yellow       string = "\x1b[33m"
  defaultStyle string = "\x1b[0m"
)

func Red(output string) string {
  return colorize(red, output)
}

func Green(output string) string {
  return colorize(green, output)
}

func Cyan(output string) string {
  return colorize(cyan, output)
}

func Yellow(output string) string {
  return colorize(yellow, output) 
}

func colorize(color string, output string) string {
  return fmt.Sprintf("%s%s%s", color, output, defaultStyle)
}