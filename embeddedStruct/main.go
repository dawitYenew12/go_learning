package main

import "fmt"

type Membership struct {
	Type string
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.MessageCharLimit = 100
	}

	return User{Name: name, Membership: membership}
}

func (u User) sendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= u.MessageCharLimit {
		return message, true
	} else {
		return "", false
	}
}

func main() {
	standardUser := newUser("Alice", "standard")
	premiumUser := newUser("Bob", "premium")
	fmt.Println(standardUser.Name, "has", standardUser.MessageCharLimit, "message character limit")
	fmt.Println(premiumUser.Name, "has", premiumUser.MessageCharLimit, "message character limit")
	fmt.Println("=======================================================================================================")
	message := "lorem ipsum hello world this is a test message for the user to send to the server and check if the message is within the character limit or not and if it is not within the character limit then the message will not be sent lorem ipsum hello world this is a test message for the user to send to the server and check if the message is within the character limit or not and if it is not within the character limit then the message will not be sent"  
	userMessage, messageStatus := standardUser.sendMessage(message, len(message))
	fmt.Printf("Message: %s\nMessage status: %t\n", userMessage, messageStatus)
	fmt.Println("=======================================================================================================")
	userMessage2, messageStatus2 := premiumUser.sendMessage(message, len(message))
	fmt.Printf("Message: %s\nMessage status: %t\n", userMessage2, messageStatus2)
}