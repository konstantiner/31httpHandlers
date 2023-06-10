package storage

import (
	"31httpHandlers/internal/entities"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const addrBD string = "userskill:paswordskill@tcp(127.0.0.1:3306)/skillbox31"

func AllUsers() []entities.User{

	db, err := sql.Open("mysql", addrBD)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `users` ")
	if err != nil {
		panic(err)
	}
	
	var allUsers = []entities.User{}

	for res.Next() {
		var user entities.User
		err = res.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		allUsers = append(allUsers, user)
	}
	
	return allUsers
}

func CreateUser(u entities.User) int {

	db, err := sql.Open("mysql", addrBD)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (name, age) VALUES(?, ?)", u.Name, u.Age)
	if err != nil {
		panic(err)
	}

	res, err := db.Query("SELECT `id` FROM `users` where `name` = ?", u.Name)
	if err != nil {
		panic(err)
	}

	var newId int
	for res.Next() {
		err = res.Scan(&newId)
		if err != nil {
			panic(err)
		}
	}

	return newId
}

func MakeFriends(SourceId int, TargetId int) {
	db, err := sql.Open("mysql", addrBD)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO friends (user_id1, user_id2) VALUES(?, ?)", SourceId, TargetId)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO friends (user_id1, user_id2) VALUES(?, ?)", TargetId, SourceId)
	if err != nil {
		panic(err)
	}
}

func DeleteUser(userId int) {
	db, err := sql.Open("mysql", addrBD)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	_, err = db.Exec("delete from users where id = ?", userId)
	if err != nil {
		panic(err)
	}
}

func DeleteFriend(userId int, friendId int) {
	db, err := sql.Open("mysql", addrBD)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	_, err = db.Exec("delete from friends where user_id1 = ? and user_id2 = ?", userId, friendId)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("delete from friends where user_id2 = ? and user_id1 = ?", userId, friendId)
	if err != nil {
		panic(err)
	}
}

func UserFriends(userId int) []int {
	db, err := sql.Open("mysql", addrBD)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT `user_id2` FROM `friends` where `user_id1` = ?", userId)
	if err != nil {
		panic(err)
	}

	var friends []int
	for res.Next() {
		var idFriends int
		err = res.Scan(&idFriends)
		if err != nil {
			panic(err)
		}
		friends = append(friends, idFriends)
	}
	return friends
}

func UserName(userId int) string {
	db, err := sql.Open("mysql", addrBD)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT `name` FROM `users` where `id` = ?", userId)
	if err != nil {
		panic(err)
	}

	var name string
	for res.Next() {
		err = res.Scan(&name)
		if err != nil {
			panic(err)
		}
	}
	
	return name
}

func UpdateUserAge(userId int, age int) {
	db, err := sql.Open("mysql", addrBD)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE users set age=? WHERE id=?", age, userId)
	if err != nil {
		panic(err)
	}
}