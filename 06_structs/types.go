package main

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo -> 丁寧な書き方
	contactInfo
}
