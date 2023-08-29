package model

type Message struct {
	MessageId 	int32 	`gorm:"primaryKey"`
	UserId 		int32 	`gorm:"primaryKey"`
	ToUserId 	int32 	`gorm:"primaryKey"`
	MessageText 	string
	CreateTime	int64
	CreateDate 	string
}
