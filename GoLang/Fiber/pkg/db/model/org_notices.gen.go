// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameORGNOTICE = "ORG_NOTICES"

// ORGNOTICE mapped from table <ORG_NOTICES>
type ORGNOTICE struct {
	NOTICEID    int64   `gorm:"column:NOTICE_ID;type:int;primaryKey" json:"NOTICE_ID"`
	NOTICEDATE  string  `gorm:"column:NOTICE_DATE;type:char;not null;index:IDX_NOTICE_DATE,priority:1" json:"NOTICE_DATE"`
	NOTICETITLE string  `gorm:"column:NOTICE_TITLE;type:nvarchar;not null;index:IDX_NOTICE_TITLE,priority:1" json:"NOTICE_TITLE"`
	NOTICEBODY  string  `gorm:"column:NOTICE_BODY;type:nvarchar;not null" json:"NOTICE_BODY"`
	NOTICETYPE  string  `gorm:"column:NOTICE_TYPE;type:char;not null" json:"NOTICE_TYPE"`
	ISACTIVE    *bool   `gorm:"column:IS_ACTIVE;type:bit;not null;default:1" json:"IS_ACTIVE"`
	CREATEDAT   string  `gorm:"column:CREATED_AT;type:char;not null" json:"CREATED_AT"`
	UPDATEDAT   *string `gorm:"column:UPDATED_AT;type:char" json:"UPDATED_AT"`
}

// TableName ORGNOTICE's table name
func (*ORGNOTICE) TableName() string {
	return TableNameORGNOTICE
}
