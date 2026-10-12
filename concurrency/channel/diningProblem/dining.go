package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Five philosophers numbered 0 through 4 live in a house where the table is laid for them; each philosopher has their own place at the table. Their only difficulty besides those of philosophy is that the dish served is a very difficult kind of spaghetti which has to be eaten with two forks. There are two forks net to each plate, so that presents no difficulty. As a consequence, however this means that no two neighbours may be eating simultaneously since there are five philosophers and five forks

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var hunger = 3
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

// To write the philosopher's eating order
// var ph chan Philosopher

var orderMutex sync.Mutex
var orderFinished []string

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()

	// eat three times
	for i := hunger; i > 0; i-- {
		//get a lock on both forks

		// To avoid the logical race condition, where all picks up their left fork or their right fork
		// So we will take the lock of lower number first - as out of all philosphers, one of them will have distinct set as it is round robin like Plate(first philospher in here)
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s take the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s take the right fork.\n", philosopher.name)
		} else {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s take the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s take the left fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s has both forks and is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)

	}

	// ph <- philosopher

	fmt.Println(philosopher.name, " is satisfied.")
	fmt.Println(philosopher.name, " left the table")

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}

func dine() {
	// wait group for dinner done
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// wait group for everyone's seated
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	//start the meal
	for i := 0; i < len(philosophers); i++ {
		//fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()

}

func main() {
	//ph = make(chan Philosopher, len(philosophers))
	// print out a welcome message
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is empty")

    time.Sleep(sleepTime)

	// start the meal
	dine()
    
	// i := 1
    // for p := range(ph){
	// 	if i >= len(philosophers){
	// 		close(ph)
	// 	}
	// 	i++
	// 	fmt.Printf("%s finished the meal\n", p.name)
	// }

	// print out finished message
	fmt.Println("The table is empty")

	// order of eating
	fmt.Printf("Order finished: %s.\n", strings.Join(orderFinished, ", "))
}
