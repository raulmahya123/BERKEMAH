package belajarmahya

import (
	"time"
)

type Payload struct {
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Nomor    string    `json:"nomor"`
	Exp      time.Time `json:"exp"`
	Iat      time.Time `json:"iat"`
	Nbf      time.Time `json:"nbf"`
}

type User struct {
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
	Nomor    string `json:"nomor" bson:"nomor"`
}

type CredentialUser struct {
	Status bool `json:"status" bson:"status"`
	Data   struct {
		Name     string `json:"name" bson:"name"`
		Username string `json:"username" bson:"username"`
		Role     string `json:"role" bson:"role"`
		Nomor    string `json:"nomor" bson:"nomor"`
	} `json:"data" bson:"data"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Pesan struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data,omitempty" bson:"data,omitempty"`
	Role    string      `json:"role,omitempty" bson:"role,omitempty"`
	Token   string      `json:"token,omitempty" bson:"token,omitempty"`
	Nomor   string      `json:"nomor,omitempty" bson:"nomor,omitempty"`
}

type Response struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type ProductOnlineCourse struct {
	ID            string `json:"id" bson:"id"`
	Name          string `json:"name" bson:"name"`
	Description   string `json:"description" bson:"description"`
	List          string `json:"list" bson:"list"`
	Price         int    `json:"price" bson:"price"`
	Authorization string `json:"authorization" bson:"authorization"`
}
