package main

import (
	"path/filepath"
	"testing"
)

func TestFetchUsers(t *testing.T) {
	fileName := "test_users.csv"
	path := filepath.Join("testdata", fileName)
	results, err := fetchUsers(path)

	if err != nil {
		t.Fatalf("Fatal error in parsing csv: %s", err)
	}

	for i, expected := range expectedUsers() {
		if expected.Name != results[i].Name {
			t.Errorf("Name: Expected: %s, Got %s", expected.Name, results[i].Name)
		}
		if expected.Email != results[i].Email {
			t.Errorf("Email: Expected: %s, Got %s", expected.Email, results[i].Email)
		}
	}
}

func TestBuildMessagesTwoUsers(t *testing.T) {
	users := []User{
		User{"Star Fox", "sf@sfox.com"},
		User{"Falco Lombardi", "fl@sfox.com"},
	}

	expectedMessages := []Message{
		Message{SecretSanta: users[0], Recipient: users[1]},
		Message{SecretSanta: users[1], Recipient: users[0]},
	}

	results := buildMessages(users)

	for i, expected := range expectedMessages {
		if expected.SecretSanta != results[i].SecretSanta {
			t.Errorf("SecretSanta: Expected: %s, Got %s", expected.SecretSanta, results[i].SecretSanta)
		}
		if expected.Recipient != results[i].Recipient {
			t.Errorf("Recipient: Expected: %s, Got %s", expected.Recipient, results[i].Recipient)
		}
	}
}

func TestBuildMessagesAllUsersAccountedFor(t *testing.T) {
	users := expectedUsers()
	results := buildMessages(users)

	secretSantaUserCount := make(map[string]int)
	recipientUserCount := make(map[string]int)

	// count the users as secret santa and recipient
	for _, message := range results {
		secretSantaUserCount[message.SecretSanta.Email]++
		recipientUserCount[message.Recipient.Email]++
	}

	// check that a user is a secret santa and recipient only once
	for _, user := range users {
		if secretSantaUserCount[user.Email] != 1 {
			t.Errorf("%s was not made a secret santa", user.Email)
		}

		if recipientUserCount[user.Email] != 1 {
			t.Errorf("%s was not made a recipient", user.Email)
		}

		delete(secretSantaUserCount, user.Email)
		delete(recipientUserCount, user.Email)
	}

	if len(secretSantaUserCount) != 0 {
		t.Errorf("Additional secret santas were generated: %v", secretSantaUserCount)
	}
	if len(recipientUserCount) != 0 {
		t.Errorf("Additional recipients were generated: %v", recipientUserCount)
	}
}

func expectedUsers() []User {
	return []User{
		User{"Philip J Fry", "pfry@planetexpress.com"},
		User{"Turanga Leela", "tleela@planetexpress.com"},
		User{"Dr. Amy Wong", "awong@planetexpress.com"},
		User{"Professor Hubert J. Farnsworth", "hfarnsworth@planetexpress.com"},
		User{"Bender Bending Rodriguez", "bender@planetexpress.com"},
	}
}
