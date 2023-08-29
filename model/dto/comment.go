package dto

type Comment struct {
	CommentID 	int64 	`gorm:"primaryKey"`
	VideoID 	int64 	`gorm:"primaryKey"`
	UserID 		int64 	`gorm:"primaryKey"`
	CommentText 	string
	CreateTime	int64
	CreateDate 	string
}
