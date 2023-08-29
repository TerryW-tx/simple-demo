package model

type Comment struct {
	CommentId 	int32 	`gorm:"primaryKey"`
	VideoId 	int32 	`gorm:"primaryKey"`
	UserId 		int32 	`gorm:"primaryKey"`
	CommentText 	string
	CreateTime	int64
	CreateDate 	string
}
