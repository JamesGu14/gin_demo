package main

type User struct {
	Id uint
}

type Student struct {
	User
	name string
}

func main() {
	s := Student{}
	s.Id = 1

}
