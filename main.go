package main

import (
	lab "lab/labs"
	"sync"
)

func main() {

	dataset := lab.ReadDataset("DS3.txt")

	wait := sync.WaitGroup{}

	wait.Add(3)

	go func() {
		lab.Lab2("result_lab2.png", dataset)
		wait.Done()
	}()

	go func() {
		lab.Lab3("result_lab3.png", dataset)
		wait.Done()
	}()

	go func() {
		lab.Lab4("result_lab4.png", dataset)
		wait.Done()
	}()

	wait.Wait()
}