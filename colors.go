package colors

import "fmt"

const (
  cyan         string = "\x1b[36m"
  gray         string = "\x1b[90m"
  green        string = "\x1b[32m"
  red          string = "\x1b[91m"
  yellow       string = "\x1b[33m"
  lightGray    string = "\x1b[37m"
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

func Gray(output string) string {
  return colorize(gray, output) 
}

func LightGray(output string) string {
  return colorize(lightGray, output) 
}

func colorize(color string, output string) string {
  return fmt.Sprintf("%s%s%s", color, output, defaultStyle)
}