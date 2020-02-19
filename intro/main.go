package main

import (
	"fmt"
	"strconv"
	"github.com/b0nbon1/learn-go/intro/greetings"
  "github.com/b0nbon1/learn-go/intro/concurrency"
)

// type Person struct {
//   name string
//   age int
// }

// func (p Person) sayName() string { // value receiver from the struct
//   return p.name + " is my name!!";
// }

// func (p *Person) setAge(age int) { // pointer receiver from the struct
//   p.age = age
// }

// type Person struct {
//   id int
//   age int
// }

// func mutateArgs(p int, o Person) { // passed by value and not a reference, by copies
//   p++
//   o.age = 30
// }

// func mutateArgs(p *int, o *Person) { // passed by reference, no copies
//   *p++ // dereference and get the content
//   o.age = 30
// }

// func (p *Person/* pointer receiver*/) setAge(age int/* value receiver*/) {
//   p.age = age
// }

type Person struct {
  // private
  name string
  // public when exported
  Age int
}

func (p *Person) Speak() string {
  if p == nil {
    return "I'm nill pointer hommie"
  }
  return "Hi! I'm " + p.name + ", I'm " + strconv.Itoa(p.Age)+ " years old"
}

func main() {
  fmt.Println(greetings.Pi)
  // fmt.Println("Hello world!")
  // var age int = 30 // with type
  // age := 30 // auto a sign acts like const

  // if age > 60 { // no brackets, only curlies
  //   fmt.Println("older than 60")
  // } else if age  >= 30 && age < 50 {
  //   fmt.Println("between 30 and 50")
  // } else {
  //   fmt.Println("between 50 and 60 or younger than 30")
  // }
  // greeting, _, err := greet("Hello ", "Bonvic")
  // if err != nil {
  //   fmt.Println("error encountered")
  // }
  // fmt.Println(greeting)
  // hi := greetings.Greet("Hi", "Bonvic!")
  // fmt.Println(hi, greetings.Pi)
  // c, python, java := true, false, "hi"
  // variables initialized with zero values
  // var i int
  // var f float64
  // var b bool
  // var s string
  // var p *int
  // fmt.Printf("%v %v %v %q %v\n", i,f,b,s,p)


  // normal for loop
  // for i := 0; i < 10; i++ {
  //   fmt.Println(i)
  // }
  // 
  // // inifite loop
  // for {
  //   fmt.Println("we don't end")
  // }

  // while loop like
  // i := 0
  // for i < 10 {
  //   fmt.Println(i)
  //   i+=2
  // }
  // bon := Person{name: "Bonvic", age: 25}
  // bon := Person{"Bonvic", 25}
  // // bon.age = 24

  // // fmt.Printf("%+v", bon)
  // shouting := bon.sayName()
  // fmt.Printf("%+v", shouting)

  // a := 0
  // b := Person{id:0, age:0}

  // // mutateArgs(a,b) // pass a copy of the value
  // mutateArgs(&a,&b) // pass address pointer of the value

  // fmt.Printf("p: %+v\n", a)
  // fmt.Printf("p: %+v\n", b)

  //  b := Person{id:0, age:0}
  // b.setAge(25)
  // fmt.Printf("p: %+v\n", b)
  // arrays
  // primes := [6] int{2,3,5,7,11,13}
  // slice array
  // primes := [] int{2,3,5,7,11,13}
  // slicing arrays
  // s := primes[1:4]
  // primesSlice := primes[:]
  // fmt.Println(addNums(primes))

  // make a slice
//   primes := make([]int, 5)
//   fmt.Println(len(primes))
//   fmt.Println(cap(primes))
//   // primes[5] = 2
//   primes = append(primes, 6)
//    fmt.Println(len(primes))
//   fmt.Println(cap(primes))
//    primes = append(primes, 6)
//  primes = append(primes, 6)
//  primes = append(primes, 6)
//  primes = append(primes, 6)
//  primes = append(primes, 6)
//  primes = append(primes, 6)
// fmt.Println(len(primes))
//   fmt.Println(cap(primes))

  // var p *Person = &Person{"Bonvic", 24}
  //  greeting := p.Speak()
  //   fmt.Println(greeting)

  // learn concurrency
  concurrencyLearn.ConcurrencyLearn()

}

// func addNums(nums []int) int {
//   res := 0
//   for _, n := range nums {
//     // println(i)
//     res += n
//   }
//   return res
// }

// func greet(greeting, person string) (string, int, error)  {
//   res := greeting + person
//   return greeting+person, len(res), nil
// }

/*
* bool
* string
* int int8 int16 int32 int64
* uint uint8 uint16 uint32 uint64 uintptr
* byte // alias for uint8
* rune // alias for int32
* float32 float64
* complex64 complex128
*/
