package struct_lotto

import "time"

type Lotto struct {
	Date                 time.Time `json:"date" bson:"date"`
	FirstPrize           string    `json:"first_prize" bson:"first_prize"`
	FrontThreeDigitPrize []string  `json:"front_three_digit_prize" bson:"front_three_digit_prize"`
	BackThreeDigitPrize  []string  `json:"back_three_digit_prize" bson:"back_three_digit_prize"`
	BackTwoDigitPrize    string    `json:"back_two_digit_prize" bson:"back_two_digit_prize"`
}
