package entities

import (
	"fmt"
)

type User struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
	Friends []int `json:"friends" db:"friends"`
}

type MakeFriends struct {
	SourceId int `json:"sourceId"`
	TargetId int `json:"targetId"`
}

func ToString(id int, u User) string {
	return fmt.Sprintf("Id: %d, Name: %s, age: %d, friends: %d \n",id, u.Name, u.Age, u.Friends)
}