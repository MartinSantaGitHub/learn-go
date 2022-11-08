package user

import "time"

type User struct {
	Id      int
	Name    string
	NewDate time.Time
	Status  bool
}

func (this *User) NewUser(id int, name string, newDate time.Time, status bool) {
	this.Id = id
	this.Name = name
	this.NewDate = newDate
	this.Status = status
}
