// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameMessage = "messages"

// Message mapped from table <messages>
type Message struct {
	MessageID   int64  `gorm:"column:message_id;primaryKey;autoIncrement:true" json:"message_id"`
	UserID      int64  `gorm:"column:user_id;primaryKey" json:"user_id"`
	ToUserID    int64  `gorm:"column:to_user_id;primaryKey" json:"to_user_id"`
	MessageText string `gorm:"column:message_text" json:"message_text"`
	CreateTime  int64  `gorm:"column:create_time" json:"create_time"`
	CreateDate  string `gorm:"column:create_date" json:"create_date"`
}

// TableName Message's table name
func (*Message) TableName() string {
	return TableNameMessage
}
