package main

type Entry struct {
	email    Email
	username string
	password EncryptedPassword
	website  string
	category Category
	note     string
}

type Email struct {
}

type EncryptedPassword struct {
}

type Category struct {
}

func main() {

}
