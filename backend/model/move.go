package model

type Move struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	TurnId uint `json:"turn_id"`
	Turn   Turn `json:"turn" gorm:"foreignkey:turn_id"`
	Disc   int  `json:"disc"`
	X      int  `json:"x"`
	Y      int  `json:"y"`
}

