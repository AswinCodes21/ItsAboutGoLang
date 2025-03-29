// // scenarios likePreventing data corruption in concurrent access  when we use sync.Muitex it can be achieved.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Counter struct {
// 	mu    sync.Mutex
// 	value int
// }

// func (c *Counter) Increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.value++
// }

// func main() {
// 	var wg sync.WaitGroup
// 	counter := Counter{}

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			counter.Increment()
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("Final Counter:", counter.value)
// }
//--------------------------------------------------------------------------------------------------------------------------------

// it is mainly used for Singleton pattern, meant for  Ensuring a function executes only once (exmpl: initializing config).
//sync.Once ensures initialize is called only once, even if multiple goroutines attempt to invoke it.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var once sync.Once

// func initialize() {
// 	fmt.Println("Initializing...")
// }

// func main() {
// 	for i := 0; i < 3; i++ {
// 		once.Do(initialize)
// 	}
// }

//-----------------------------------------------------------------------------------------------------------------------------------

// it ensures all the go routines functions completed or not  before the program exists
// we can use sync.waitGroup for task synchronizatio sync.WaitGroup waits for all goroutines to finish before proceeding.
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d started\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("All workers completed")
// }
//----------------------------------------------------------------------------------------------------------------------------------------------------------------------

// FoR  tIME HANDLING we can use time.after time.ticker Executing functions at intervals.
//time.Ticker triggers events periodically.
//Useful for repeated tasks like cleanup jobs.

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop()

// 	for i := 0; i < 3; i++ {
// 		<-ticker.C
// 		fmt.Println("Task executed at", time.Now())
// 	}
// }

//-----------------------------------------------------------------------------------------------------------------------------------------------------

// we can prevent crashes and logging errors by using panic and recover keyword. recover() prevents application crashes by handling panic.

// package main

// import "fmt"

// func riskyFunction() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered from:", r)
// 		}
// 	}()
// 	panic("Something went wrong")
// }

// func main() {
// 	riskyFunction()
// 	fmt.Println("Program continues after panic recovery")
// }
//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------

//using generics we can Write reusable functions for different data types. --its similiar to ovrloading and ovrriding convpt in c#
// package main

// import "fmt"

// func Sum[T int | float64](a, b T) T {
//     return a + b
// }

// func main() {
//     fmt.Println(Sum(3, 7))
//     fmt.Println(Sum(4.5, 2.5))
// }
//------------------------------------------------------------------------------------------------------------------------------------------------------------

// We caan generate random nums using math.rand
//crypto/rand generates cryptographically secure random numbers

// package main

// import (
// 	"crypto/rand"
// 	"fmt"
// 	"math/big"
// )

// func main() {
// 	num, _ := rand.Int(rand.Reader, big.NewInt(100))
// 	fmt.Println("Random Number:", num)
// }
//---------------------------------------------------------------------------------------------------------------------------------

// Scenario like Reading and writing files.
//os.WriteFile writes data to a file.
//os.ReadFile reads data from a file.

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	err := os.WriteFile("sample.txt", []byte("Hello, mapla"), 0644)
// 	if err != nil {
// 		fmt.Println("Error writing file:", err)
// 		return
// 	}

// 	data, err := os.ReadFile("sample.txt")
// 	if err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}

// 	fmt.Println("File Content:", string(data))
// }
//------------------------------------------------------------------------------------------------------------------

//Scenario: Parsing arguments from CLI.

//... oit Uses flag package to parse CLI arguments.

package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Aswin", "Ur name")
	flag.Parse()
	fmt.Println("Hlo,", *name)
}
