package concurrencyLearn

import (
	"fmt"
	"math/rand"
	"time"
)

func ConcurrencyLearn() {
  start := time.Now()
  defer func() {
    fmt.Println(time.Since(start))
  }()
  // evilNinjas := []string{"Tommy", "Johnny", "Bobby", "Andy"}
  
  // without the goroutines
  // for _, evilNinja := range evilNinjas {
  //   attack(evilNinja)
  // }

   // with routines
  //  for _, evilNinja := range evilNinjas {
  //   go attack(evilNinja)
  // }
  // smokeSignal := make(chan bool, 1)
  // evilNinja := "Bonvic"
  // go attack(evilNinja, smokeSignal)
  // smokeSignal <- false // gives out a deadlock but if you add a buffer on the channel it will work
  // // smokeSignal <- false
  // fmt.Println(<-smokeSignal)

  // time.Sleep(time.Second*2) // use channels insteads to synchronize
  // message := make(chan string, 2)
  // message <- "First message"
  // message <- "Second message"
  // fmt.Println(<-message)
  // fmt.Println(<-message)

  // loops and channels, we know how many loops
  // channel := make(chan string)
  // numRounds := 3
  // go throwingNinjaStar(channel, numRounds)
  // for i := 0; i < numRounds; i++ {
  //   fmt.Println(<-channel)
  // }

  // channel := make(chan string)
  // go throwingNinjaStar(channel)
  // for message := range channel {
  //   fmt.Println(message)
  // }

  channel := make(chan string)
  go throwingNinjaStar(channel)
  for {
    message, open := <- channel
    if !open {
      break
    }
    fmt.Println(message)
  }
}

func attack(target string, attacked chan bool) {
  time.Sleep(time.Second)
  fmt.Println("Thowing ninja stars at ", target)
  attacked <- true
}

func throwingNinjaStar(channel chan string) {
  rand.Seed(time.Now().UnixNano())
  numRounds := 3
  for i := 0; i < numRounds; i++ {
    score := rand.Intn(10)
    channel <- fmt.Sprint("You scored: ", score)
  }
  // once channel done will close to avoid a deadlock
  close(channel)
}

