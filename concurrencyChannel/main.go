package main

import (
	"fmt"
	"time"
)

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan // Consume one token per database
	}
}


func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}

func main() {
	numDBs := 5 // Number of databases to wait for
	dbChan, count := getDBsChannel(numDBs)

	// Wait for all databases
	for i := 0; i < numDBs; i++ {
		waitForDBs(numDBs, dbChan)
	}

	// Give some time for goroutine to finish (since it prints messages asynchronously)
	time.Sleep(time.Millisecond * 100)

	fmt.Printf("Total databases online: %d\n", *count)
}