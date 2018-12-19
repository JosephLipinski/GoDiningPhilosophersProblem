package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

//Philosopher is a struct used to define all the attributes of a philosopher
type Philosopher struct {
	name                        string
	leftNeighbor, rightNeighbor *Philosopher
	leftFork, rightFork         chan bool
}

var wg sync.WaitGroup
var forks = [10]chan bool{}
var philosophers = [10]Philosopher{}

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

//SayHello is a simple function that makes philosophers state their name
func SayHello(_philosopher *Philosopher) {
	fmt.Printf("Hello my name is %s\n", _philosopher.name)
}

//Think is a function that causes philosophers to think for some amount of time
func Think(_philosopher *Philosopher) {
	fmt.Printf("I, %s, am thinking\n", _philosopher.name)
	time.Sleep(time.Duration(1+rand.Intn(10)) * time.Second)
	fmt.Printf("I, %s, am done thinking\n", _philosopher.name)
}

//GetLeftFork is a function that causes philosopher's to obtain a fork
func GetLeftFork(_philosopher *Philosopher) {
	fmt.Printf("I, %s, need a fork\n", _philosopher.name)
	forkAcquired := false
	for {
		if !forkAcquired {
			select {
			case status := <-_philosopher.leftFork:
				if status {
					_philosopher.leftFork <- false
					forkAcquired = true
					break
				} else {
					_philosopher.leftFork <- false
					fmt.Printf("I, %s, still need a fork\n", _philosopher.name)
					break
				}
			default:
				time.Sleep(time.Second * 2)
				fmt.Printf("I, %s, still need a fork\n", _philosopher.name)
				break
			}
		} else {
			break
		}
		time.Sleep(time.Second)
	}
}

//GetRightFork is a function that causes philosopher's to obtain a fork
func GetRightFork(_philosopher *Philosopher) {
	fmt.Printf("I, %s, need a fork\n", _philosopher.name)
	forkAcquired := false
	for {
		if !forkAcquired {
			select {
			case status := <-_philosopher.rightFork:
				if status {
					_philosopher.rightFork <- false
					forkAcquired = true
					break
				} else {
					_philosopher.rightFork <- false
					fmt.Printf("I, %s, still need a fork\n", _philosopher.name)
					break
				}
			default:
				time.Sleep(time.Second * 2)
				fmt.Printf("I, %s, still need a fork\n", _philosopher.name)
				break
			}
		} else {
			break
		}
		time.Sleep(time.Second)
	}
}

//Eat is a function that causes philosopher's to eat
func Eat(_philosopher *Philosopher) {
	fmt.Printf("I, %s, have two forks and think that theses noodles are delicious\n", _philosopher.name)
}

//PutForksDown is a functions that causes philosophers to put their forks back
func PutForksDown(_philosopher *Philosopher) {
	fork1, fork2 := false, false
	for {
		if !fork1 {
			select {
			case status := <-_philosopher.leftFork:
				if !status {
					_philosopher.leftFork <- true
					fork1 = true
					break
				} else {
					_philosopher.leftFork <- false
					break
				}
			default:
				break
			}
		} else {
			break
		}
		time.Sleep(time.Second)
	}
	for {
		if !fork2 {
			select {
			case status := <-_philosopher.rightFork:
				if !status {
					_philosopher.rightFork <- true
					fork2 = true
				} else {
					_philosopher.rightFork <- false
					break
				}
			default:
				break
			}
		} else {
			break
		}
		time.Sleep(time.Second)
	}
}

//Dine is a function that handles all the behaviors of philosophers
func Dine(_philosopher *Philosopher) {
	SayHello(_philosopher)
	Think(_philosopher)
	GetLeftFork(_philosopher)
	fmt.Println("I got a fork")
	GetRightFork(_philosopher)
	fmt.Println("I got a fork")
	Eat(_philosopher)
	PutForksDown(_philosopher)
	fmt.Println("ALL DONE")
	wg.Done()
}

//MakeFork produces a channel of type boolean and adds it to the forks array
func MakeFork(index int) {
	forkChannel := make(chan bool, 1)
	forkChannel <- true
	forks[index] = forkChannel
}

func MakePhilosopher(_numberOfPhilosophers int) {
	for i := 0; i < _numberOfPhilosophers; i++ {
		_philosopher := &Philosopher{philosopherNames[i], nil, nil, nil, nil}
		philosophers[i] = *_philosopher
	}
}

func SetupLeftSide(_numberOfPhilosophers int) {
	for i := 0; i < _numberOfPhilosophers-2; i++ {
		philosophers[i].leftFork = forks[i]
	}
	philosophers[_numberOfPhilosophers-1].leftFork = forks[_numberOfPhilosophers-1]
}

func SetUpRightSide(_numberOfPhilosophers int) {
	for i := 0; i < _numberOfPhilosophers-2; i++ {
		philosophers[i].rightFork = forks[i+1]
	}
	philosophers[_numberOfPhilosophers-1].rightFork = forks[_numberOfPhilosophers-2]
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
	if numberOfPhilosophers < 2 {
		fmt.Printf("You have selected an insufficent number to simulate. The defualt number of philosophers will be used instead")
		numberOfPhilosophers = 2
	} else if numberOfPhilosophers > 10 {
		fmt.Printf("You have selected too many philosophers to simulate. The maximum number of philosophers will be used instead")
		numberOfPhilosophers = 10
	}

	for i := 0; i < numberOfPhilosophers; i++ {
		MakeFork(i)
	}
	/*
		MakePhilosopher(numberOfPhilosophers)
		SetupLeftSide(numberOfPhilosophers)
		SetUpRightSide(numberOfPhilosophers)
	*/

	_philosopher := &Philosopher{philosopherNames[0], nil, nil, nil, nil}
	philosophers[0] = *_philosopher
	_philosopher = &Philosopher{philosopherNames[1], nil, nil, nil, nil}
	philosophers[1] = *_philosopher
	_philosopher = &Philosopher{philosopherNames[2], nil, nil, nil, nil}
	philosophers[2] = *_philosopher

	philosophers[0].leftFork = forks[0]
	philosophers[0].rightFork = forks[1]
	philosophers[1].leftFork = forks[0]
	philosophers[1].rightFork = forks[1]

	philosophers[2].leftFork = forks[0]
	philosophers[2].rightFork = forks[2]

	/*
		0
		1

		1
		0

		0
		1

		1
		2

		0
		2
	*/

	for i := 0; i < numberOfPhilosophers; i++ {
		go Dine(&philosophers[i])
		wg.Add(1)
	}
	/**/
	//time.Sleep(time.Second)
	wg.Wait()

}
