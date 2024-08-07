// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameORGSPECIALCLOSE = "ORG_SPECIAL_CLOSE"

// ORGSPECIALCLOSE mapped from table <ORG_SPECIAL_CLOSE>
type ORGSPECIALCLOSE struct {
	CLOSEID     int64   `gorm:"column:CLOSE_ID;type:int;primaryKey" json:"CLOSE_ID"`
	CLOSESTART  string  `gorm:"column:CLOSE_START;type:char;not null" json:"CLOSE_START"`
	CLOSEEND    string  `gorm:"column:CLOSE_END;type:char;not null" json:"CLOSE_END"`
	CLOSENAME   string  `gorm:"column:CLOSE_NAME;type:nvarchar;not null" json:"CLOSE_NAME"`
	CLOSEREASON *string `gorm:"column:CLOSE_REASON;type:nvarchar" json:"CLOSE_REASON"`
	OFFICEID    int64   `gorm:"column:OFFICE_ID;type:int;not null;index:IDX_OFFICE_ID,priority:1" json:"OFFICE_ID"`
	FLOORID     *int64  `gorm:"column:FLOOR_ID;type:int;index:IDX_FLOOR_ID,priority:1" json:"FLOOR_ID"`
	ROOMID      *int64  `gorm:"column:ROOM_ID;type:int;index:IDX_ROOM_ID,priority:1" json:"ROOM_ID"`
	CREATEDAT   string  `gorm:"column:CREATED_AT;type:char;not null" json:"CREATED_AT"`
	UPDATEDAT   *string `gorm:"column:UPDATED_AT;type:char" json:"UPDATED_AT"`
}

// TableName ORGSPECIALCLOSE's table name
func (*ORGSPECIALCLOSE) TableName() string {
	return TableNameORGSPECIALCLOSE
}
