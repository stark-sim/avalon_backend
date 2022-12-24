// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type GameRequest struct {
	ID string `json:"id"`
}

type JoinRoomInput struct {
	ShortCode string `json:"shortCode"`
	UserID    string `json:"userID"`
}

type RoomRequest struct {
	ID string `json:"id"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (User) IsEntity() {}
