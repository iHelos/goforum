package database

import ()

type User struct {
	id          int
	name        string
	username    string
	email       string
	about       string
	isanonymous bool
}

type Forum struct {
	id          int
	name        string
	short_name  string
	email       string
	about       string
	isAnonymous bool
}

type Post struct {
	id          int
	date        string
	dislikes    int
	isApproved  bool
	about       string
	isAnonymous bool
}