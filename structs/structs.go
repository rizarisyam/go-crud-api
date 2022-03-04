package structs

import "time"

// Posts is a representation of a post
type Posts struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Category     string    `json:"category"`
	Created_date time.Time `json:"created_date"`
	Updated_date time.Time `json:"updated_date"`
	Status       string    `json:"status"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type RiskProfile struct {
	ID     int     `json:"id"`
	Userid int     `json:"userid"`
	MM     float32 `json:"mm"`
	Bond   float32 `json:"bond"`
	Stock  float32 `json:"stock"`
}

// Result is an array of post
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
