package dto

type Favorite struct {
	FavoriteID 	int64 	`gorm:"primaryKey;autoIncrement"`
	UserID 		int64 	`gorm:"primaryKey"`
	VideoID		int64 	`gorm:"primaryKey"`
	CreateTime	int64
}
