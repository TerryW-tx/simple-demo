package dto

type Video struct {
	VideoID		int64 	`gorm:"primaryKey"`
	UserID		int64 	`gorm:"primaryKey"`
	Token		string
	CreateTime	int64
	PlayUrl		string
	CoverUrl	string
	FavoriteCount	int
	CommentCount	int
	Title		string
}
