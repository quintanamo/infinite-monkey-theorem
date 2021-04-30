package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
	"math/rand"
	"time"
	"context"
	"sync"
)

// decalre globals
var _madeClone = false
var _waitGroup = sync.WaitGroup{}
var _totalCount = 0
var _context context.Context
var _cancel context.CancelFunc

func generateClone(content []byte) {
	// defer done call until method finishes
	defer _waitGroup.Done()

	// create a seed for the random number generator
	rand.Seed(time.Now().UnixNano())

	// declare function variables
	var clonedContent = make([]byte,len(content))
	var isClone bool = false

	// while the original file content has not been cloned, loop
	for (isClone == false) {
		select {
			// check if a monkey has successfully cloned the original file's content
			case <-_context.Done():
				return
			// keep on trying to clone it
			default:
				for i, _ := range content {
					// generate random bytes in the array
					clonedContent[i] = byte(rand.Intn(93) + 33)
				}
				// check bytes in array to see if it matches the original
				for i, v := range clonedContent {
					if (content[i] != v) {
						isClone = false
						break
					}
					isClone = true
				}
		}
	}
	// end the other routines
	_cancel()
	_madeClone = true
	if (isClone)  {
		fmt.Println("The monkeys have generated a clone of the original file!")
		fmt.Println("Generated content: ", string(clonedContent))
		fmt.Println("Original  content: ", string(content))
	}
}

func main() {
	// declare variables
	var fileName string
	var monkeys int

	// initialize _context and _cancel globals
	_context, _cancel = context.WithCancel(context.Background())

	// read command line arguments
	if (len(os.Args) > 1) {
		fileName = os.Args[1]
		if (len(os.Args) > 2) {
			monkeys, _ = strconv.Atoi(os.Args[2])
		}
	} else {
		fmt.Println("No file name provided.")
		return
	}

	// if no monkeys specified, default to one
	if (monkeys == 0) {
		monkeys = 1
	}

	fmt.Println("\nReplicating the contents of", fileName, "with", monkeys, "monkeys...")

	// read contents of file
	fileContent, err := ioutil.ReadFile(fileName)
	if (err != nil) {
		fmt.Println("Error reading from file " + fileName)
	}

	// get start time
	start := time.Now()

	// run a go routine for each monkey
	for i := 0; i < monkeys; i++ {
		_waitGroup.Add(1)
		go generateClone(fileContent)
	}

	// wait until routines finish executing
	_waitGroup.Wait()

	// get end time and subtract the start time from it
	end := time.Now()
	elapsed := end.Sub(start)
	
	fmt.Println("Generated a clone in", elapsed, "!")
	fmt.Println("\nPress enter to exit.")
	fmt.Scanf("Exit")
}