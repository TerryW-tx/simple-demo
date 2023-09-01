package dto

type Follow struct {
	FollowID 	int64 	`gorm:"primaryKey;autoIncrement"`
	FollowbyID 	int64 	`gorm:"primaryKey"`
	FollowerID 	int64 	`gorm:"primaryKey"`
	CreateTime	int64
}
