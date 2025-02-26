package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	// ?
	return msg.getMessage(), len(msg.getMessage())
}

type message interface {
	// ?
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main() {
	bm := birthdayMessage{
		birthdayTime:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		recipientName: "Alice",
	}
	sr := sendingReport{
		reportName:    "Monthly Report",
		numberOfSends: 100,
	}

	message1, message1Len := sendMessage(bm)
	message2, message2Len := sendMessage(sr)
	fmt.Printf("Message: %s\nMessage length: %d\n", message1, message1Len)
	fmt.Printf("Message: %s\nMessage length: %d\n", message2, message2Len)
}