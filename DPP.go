package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

//Philosopher is a struct used to define all the attributes of a philosopher
type Philosopher struct {
	name                        string
	leftNeighbor, rightNeighbor *Philosopher
	leftFork, rightFork         chan bool
}

var wg sync.WaitGroup
var forks = [10]chan bool{}
var philosopherNames = []string{
	"Aristotle",
	"Immanuel Kant",
	"Confucius",
	"Rene Descarte",
	"John Locke",
	"Voltaire",
	"Baron de Montesquieu",
	"Sun Tzu",
	"Socrates",
	"Siddhartha Gautama Buddha"}
var philosophers = [10]Philosopher{}

//MakeFork produces a channel of type boolean and adds it to the forks array
func MakeFork(index int) {
	forkChannel := make(chan bool, 1)
	forkChannel <- true
	forks[index] = forkChannel
}

func SayHello(_philosopher *Philosopher) {
	fmt.Printf("Hello my name is %s\n", _philosopher.name)
}

func Dine(_philosopher *Philosopher) {

	//Think
	//GetForks
	//Dine
	//wg.Done()
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the number (maximum of 10) of philisophers that you would like to simulate?")
	number, _ := inputReader.ReadString('\n')
	number = strings.TrimSuffix(number, "\n")
	numberOfPhilosophers, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("An error has occurred based on the number that you have entered. The defualt value of two will be used")
		numberOfPhilosophers = 2
	}
	if numberOfPhilosophers < 1 {
		fmt.Printf("You have selected an insufficent number to simulate. The defualt number of philosophers will be used instead")
		numberOfPhilosophers = 2
	} else if numberOfPhilosophers > 10 {
		fmt.Printf("You have selected too many philosophers to simulate. The maximum number of philosophers will be used instead")
		numberOfPhilosophers = 10
	}

	for i := 0; i < numberOfPhilosophers; i++ {
		MakeFork(i)
	}

	for i := 0; i < numberOfPhilosophers; i++ {
		if i > 0 {
			_philosopher := &Philosopher{philosopherNames[i], &philosophers[i-1], nil, forks[numberOfPhilosophers-1], nil}
			philosophers[i] = *_philosopher
		} else {
			_philosopher := &Philosopher{philosopherNames[i], nil, nil, forks[i], nil}
			philosophers[i] = *_philosopher
		}
	}

	philosophers[0].leftNeighbor = &philosophers[numberOfPhilosophers-1]

	for i := numberOfPhilosophers - 1; i > 0; i-- {
		if i == numberOfPhilosophers-1 {
			philosophers[i].rightNeighbor = &philosophers[0]
			philosophers[i].rightFork = forks[0]
		} else {
			philosophers[i].rightNeighbor = &philosophers[i+1]
			philosophers[i].rightFork = forks[i+1]
		}
	}
	for i := 0; i < numberOfPhilosophers; i++ {
		SayHello(&philosophers[i])
		Dine(&philosophers[i])
		//wg.Add(1)
	}
	/**/
	//time.Sleep(time.Second)
	wg.Wait()

}
