package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

// User participating in Secret Santa
type User struct {
	Name  string
	Email string
}

// Message to participants in Secret Santa
type Message struct {
	SecretSanta User
	Recipient   User
}

func main() {
	users, err := FetchUsers(os.Args[1])

	if err != nil {
		log.Fatalf("Unable to fetch Users: %v", err)
	}

	messages := BuildMessages(users)
	fmt.Printf("SECRET SANTA: %+v", messages)
}

//FetchUsers get the users from input csv
func FetchUsers(fileName string) ([]User, error) {
	csvFile, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var users []User

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		users = append(users, User{
			Name:  record[0],
			Email: record[1],
		})
	}

	return users, nil
}

// BuildMessages pairs up users and gets messages ready
func BuildMessages(users []User) []Message {
	rand.Seed(time.Now().UnixNano())
	var messages []Message

	santa := users[0]
	for len(users) > 0 {
		users = users[1:]

		rand.Shuffle(
			len(users),
			func(i, j int) { users[i], users[j] = users[j], users[i] },
		)

		recipient := users[0]
		messages = append(messages, Message{SecretSanta: santa, Recipient: recipient})
		santa = recipient
	}
	return messages
}
