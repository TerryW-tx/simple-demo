package model

type Follow struct {
	FollowId 	int32 	`gorm:"primaryKey"`
	FollowbyId 	int32 	`gorm:"primaryKey"`
	FollowerId 	int32 	`gorm:"primaryKey"`
	CreateTime	int64
}
