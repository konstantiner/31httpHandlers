package services

import (
	"fmt"
	"encoding/json"
	"31httpHandlers/internal/entities"
	"31httpHandlers/internal/storage"
)

func GetAllUsers() (b []byte){
	var response []entities.User
	response = append(response, storage.AllUsers()...)
	b, _ = json.Marshal(response)
	return b
}

func CreateUser(u entities.User) (userId int) {
	userId = storage.CreateUser(u)
	return
}

func NewFriends(SourceId int, TargetId int) (b []byte){
	username1 := storage.UserName(SourceId) 
    username2 := storage.UserName(TargetId) 
	
	//проверяем: если уже дружат, то не будем ещё раз добавлять в список друзей
    var allFriendsUser []int = storage.UserFriends(SourceId)
    for _, x := range allFriendsUser {
        if x == TargetId {
            b = []byte(fmt.Sprintf("%s и %s уже дружат", username1, username2))
            return
        }
    }

    storage.MakeFriends(SourceId, TargetId)
	b = []byte(fmt.Sprintf("%s и %s теперь друзья", username1, username2))
	return
}

func DeleteUser(TargetId int) (b []byte) {
	//вытащим список друзей пользователя, зайдем к каждому другу и удалим у него удаленного пользователя из списка друзей
	var allFriends []int = storage.UserFriends(TargetId)
	for _, x := range allFriends {
		var friendsUser []int = storage.UserFriends(x)
		for _, z := range friendsUser {
			if TargetId == z {
				storage.DeleteFriend(x, z)				
			}
		}
	}
	nameDeleteUser :=  storage.UserName(TargetId)
	storage.DeleteUser(TargetId)

	b = []byte(fmt.Sprintf("Пользователь %s удален.", nameDeleteUser))
	return
}

func UserFriends(userID int) (b []byte){
	var allFriends []int = storage.UserFriends(userID)
	friendsName := ""
	for _, x := range allFriends {
		friendsName += fmt.Sprintf("ID: %d, Имя: %s\n",x, storage.UserName(x))
	}

	b = []byte(fmt.Sprintf("Друзья пользователя: %s\n", friendsName))
	return
}

func UpdateUserAge(userId int, age int) (b []byte) {
	storage.UpdateUserAge(userId, age)
	b = []byte("Возраст пользователя успешно обновлён")
	return
}