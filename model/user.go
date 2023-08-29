package model

type User struct {
	UserId 		int32 	`gorm:"primaryKey"`
	Username 	string
	Password 	string
	CreateTime	int64
	Token 		string
	TokenUpdateTime int64
	Avatar 		string
	BackgroundImage string
	FollowCount 	int
	FollowerCount 	int
	FavoriteCount 	int
	WorkCount 	int
}
