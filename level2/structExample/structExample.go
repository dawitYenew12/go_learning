package structExample

import "fmt"

type messageToSend struct {
    phoneNumber int
    message string
}

func test(m messageToSend) {
    fmt.Println("============================================================")
    fmt.Printf("message to send: '%s' to: '%v'\n", m.message, m.phoneNumber)
    fmt.Println("============================================================")
}

func StructExample() {
    test(messageToSend{1234567890, "Hello, World!"})
}
