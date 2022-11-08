package main

import (
	"fmt"
	"time"

	u "user"
)

type something struct {
	u.User
}

func main() {
	user := new(something)

	user.NewUser(1, "Santa", time.Now(), true)

	fmt.Println(user.User)
}
