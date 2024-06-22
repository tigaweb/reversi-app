package model

type Square struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	TurnId uint `json:"turn_id"`
	Turn   Turn `json:"turn" gorm:"foreignkey:turn_id"`
	X      int  `json:"x"`
	Y      int  `json:"y"`
	Disc   int  `json:"disc"`
}
