package model

type Book struct {
	Title    string
	Author   string
	Url      string
	Desc     string
	Category string
}

type Category struct {
	Id   int
	Name string
}
