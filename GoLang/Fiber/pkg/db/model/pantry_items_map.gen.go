// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePANTRYITEMSMAP = "PANTRY_ITEMS_MAP"

// PANTRYITEMSMAP mapped from table <PANTRY_ITEMS_MAP>
type PANTRYITEMSMAP struct {
	PANTRYID int64 `gorm:"column:PANTRY_ID;type:int;primaryKey;index:IDX_PANTRY_ITEM_MAPPING,priority:1" json:"PANTRY_ID"`
	ITEMID   int64 `gorm:"column:ITEM_ID;type:int;primaryKey;index:IDX_PANTRY_ITEM_MAPPING,priority:2" json:"ITEM_ID"`
}

// TableName PANTRYITEMSMAP's table name
func (*PANTRYITEMSMAP) TableName() string {
	return TableNamePANTRYITEMSMAP
}
