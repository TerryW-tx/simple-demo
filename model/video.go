package model

type Video struct {
	VideoId		int32 	`gorm:"primaryKey"`
	UserId		int32 	`gorm:"primaryKey"`
	Token		string
	CreateAt	int64
	PlayUrl		string
	CoverUrl	string
	FavoriteCount	int
	CommentCount	int
	Title		string
}
