package dto

type Message struct {
	MessageID 	int64 	`gorm:"primaryKey;autoIncrement"`
	UserID		int64 	`gorm:"primaryKey"`
	ToUserID 	int64 	`gorm:"primaryKey"`
	MessageText 	string
	CreateTime	int64
	CreateDate 	string
}
