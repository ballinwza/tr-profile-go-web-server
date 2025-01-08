package models_lotto

import "time"

type Lotto struct {
	Date                 time.Time `json:"date" bson:"date" example:"2025-01-01T00:00:00Z"`
	FirstPrize           string    `json:"first_prize" bson:"first_prize" example:"730209"`
	FrontThreeDigitPrize []string  `json:"front_three_digit_prize" bson:"front_three_digit_prize" example:"['446', '065']"`
	BackThreeDigitPrize  []string  `json:"back_three_digit_prize" bson:"back_three_digit_prize" example:"['376', '297']"`
	BackTwoDigitPrize    string    `json:"back_two_digit_prize" bson:"back_two_digit_prize" example:"51"`
}
