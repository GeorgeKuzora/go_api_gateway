package api

import "time"

type request struct {
	id int
	card_id int
	merchant string
	amount int
	timestamp time.Time
}

type response struct {
	id int
	request_id int
	result string
}