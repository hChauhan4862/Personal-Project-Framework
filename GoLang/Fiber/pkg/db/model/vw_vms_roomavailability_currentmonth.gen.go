// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameVwVMSRoomAvailabilityCurrentMonth = "vw_VMS_RoomAvailability_CurrentMonth"

// VwVMSRoomAvailabilityCurrentMonth mapped from table <vw_VMS_RoomAvailability_CurrentMonth>
type VwVMSRoomAvailabilityCurrentMonth struct {
	ROOMID    int64      `gorm:"column:ROOM_ID;type:int;not null" json:"ROOM_ID"`
	ROOMNAME  string     `gorm:"column:ROOM_NAME;type:nvarchar;not null" json:"ROOM_NAME"`
	SlotStart *time.Time `gorm:"column:SlotStart;type:datetime" json:"SlotStart"`
	SlotEnd   *time.Time `gorm:"column:SlotEnd;type:datetime" json:"SlotEnd"`
	Status    string     `gorm:"column:Status;type:varchar;not null" json:"Status"`
}

// TableName VwVMSRoomAvailabilityCurrentMonth's table name
func (*VwVMSRoomAvailabilityCurrentMonth) TableName() string {
	return TableNameVwVMSRoomAvailabilityCurrentMonth
}
