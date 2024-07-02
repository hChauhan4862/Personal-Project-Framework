// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCOMMUNICATIONWEBPUSH = "COMMUNICATION_WEBPUSH"

// COMMUNICATIONWEBPUSH mapped from table <COMMUNICATION_WEBPUSH>
type COMMUNICATIONWEBPUSH struct {
	SEQID          int64   `gorm:"column:SEQ_ID;type:int;primaryKey" json:"SEQ_ID"`
	MESSAGEID      string  `gorm:"column:MESSAGE_ID;type:nvarchar;not null" json:"MESSAGE_ID"`
	PUSHDATETIME   string  `gorm:"column:PUSH_DATE_TIME;type:char;not null" json:"PUSH_DATE_TIME"`
	USERID         string  `gorm:"column:USER_ID;type:nvarchar;not null" json:"USER_ID"`
	MESSAGETYPE    string  `gorm:"column:MESSAGE_TYPE;type:char;not null" json:"MESSAGE_TYPE"`
	MESSAGETITLEEN string  `gorm:"column:MESSAGE_TITLE_EN;type:nvarchar;not null" json:"MESSAGE_TITLE_EN"`
	MESSAGEBODYEN  string  `gorm:"column:MESSAGE_BODY_EN;type:nvarchar;not null" json:"MESSAGE_BODY_EN"`
	MESSAGETITLEKO string  `gorm:"column:MESSAGE_TITLE_KO;type:nvarchar;not null" json:"MESSAGE_TITLE_KO"`
	MESSAGEBODYKO  string  `gorm:"column:MESSAGE_BODY_KO;type:nvarchar;not null" json:"MESSAGE_BODY_KO"`
	SENTSTATUS     *bool   `gorm:"column:SENT_STATUS;type:bit;not null;default:0" json:"SENT_STATUS"`
	SENTDATETIME   *string `gorm:"column:SENT_DATE_TIME;type:char" json:"SENT_DATE_TIME"`
}

// TableName COMMUNICATIONWEBPUSH's table name
func (*COMMUNICATIONWEBPUSH) TableName() string {
	return TableNameCOMMUNICATIONWEBPUSH
}
