package dto

type Video struct {
	VideoID		int64 	`gorm:"primaryKey;autoIncrement"`
	UserID		int64 	`gorm:"primaryKey"`
	Token		string
	CreateTime	int64
	PlayURL		string
	CoverURL	string
	FavoriteCount	int
	CommentCount	int
	Title		string
}
