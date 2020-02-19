package greetings

const Pi = 3.14

func Greet(greeting, person string) string  {
  res := greeting + " " + person
  return res
}
