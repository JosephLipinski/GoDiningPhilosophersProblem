package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

//Philosopher is a struct used to define all the attributes of a philosopher
type Philosopher struct {
	name                        string
	leftNeighbor, rightNeighbor *Philosopher
	leftFork, rightFork         chan bool
}

var wg sync.WaitGroup
var forks = []chan bool{}
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

//MakeFork produces a channel of type boolean and adds it to the forks array
func MakeFork(index int) {
	forkChannel := make(chan bool, 1)
	forkChannel <- false
	forks[index] = forkChannel
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the number (maximum of 10) of philisophers that you would like to simulate?")
	number, _ := inputReader.ReadString('\n')
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

	//time.Sleep(time.Second)
	wg.Wait()
	fmt.Printf("hello, world\n")
}
