package main

type User struct {
	ID       int
	Username string
	Email    string
}

type Storage interface {
	FindUserByEmail(string) User
	FindUsers() []User
}

type InMemoryStorage struct {
	users []User
}

func NewInMemoryStorage() InMemoryStorage {
	return InMemoryStorage{
		users: []User{
			{
				ID:       1,
				Username: "MilesDavis",
				Email:    "miles.davis@gmail.com",
			},
			{
				ID:       2,
				Username: "JohnColtrane",
				Email:    "john.coltrane@gmail.com",
			},
		},
	}
}

func (m InMemoryStorage) FindUserByEmail(email string) User {
	var user User
	for i := range m.users {
		if m.users[i].Email == email {
			return m.users[i]
		}
	}
	return user
}

func (m InMemoryStorage) FindUsers() []User {
	return m.users
}
