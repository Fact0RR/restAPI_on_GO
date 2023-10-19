package rest

import "time"

type Duty struct {
	Id        int       `json:"duty_id"`
	User_id   string    `json:"user_id"`
	Dish      bool      `json:"dish"`
	Trash     bool      `json:"trash"`
	TimeDish  time.Time `json:"time_dish"`
	TimeTrash time.Time `json:"time_d"`
}
