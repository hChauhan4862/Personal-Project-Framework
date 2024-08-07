// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameVMSRESTRICTIONSETTING = "VMS_RESTRICTION_SETTINGS"

// VMSRESTRICTIONSETTING mapped from table <VMS_RESTRICTION_SETTINGS>
type VMSRESTRICTIONSETTING struct {
	ID                     int64  `gorm:"column:ID;type:int;primaryKey" json:"ID"`
	LOCATION               string `gorm:"column:LOCATION;type:nvarchar;not null;index:IDX_LOCATION,priority:1" json:"LOCATION"`
	LOCATIONID             int64  `gorm:"column:LOCATION_ID;type:int;not null;index:IDX_LOCATION,priority:2" json:"LOCATION_ID"`
	ISCONFIRMATIONREQUIRED bool   `gorm:"column:IS_CONFIRMATION_REQUIRED;type:bit;not null" json:"IS_CONFIRMATION_REQUIRED"`
	CONFIRMATIONWAITTIME   int64  `gorm:"column:CONFIRMATION_WAIT_TIME;type:int;not null" json:"CONFIRMATION_WAIT_TIME"`
	ISAPPROVALREQUIRED     bool   `gorm:"column:IS_APPROVAL_REQUIRED;type:bit;not null" json:"IS_APPROVAL_REQUIRED"`
	APPROVALWAITTIME       int64  `gorm:"column:APPROVAL_WAIT_TIME;type:int;not null" json:"APPROVAL_WAIT_TIME"`
	ISAUTOCANCELALLOWED    bool   `gorm:"column:IS_AUTO_CANCEL_ALLOWED;type:bit;not null" json:"IS_AUTO_CANCEL_ALLOWED"`
	CANCELLATIONWAITTIME   int64  `gorm:"column:CANCELLATION_WAIT_TIME;type:int;not null" json:"CANCELLATION_WAIT_TIME"`
	ISMESSAGEALLOWED       bool   `gorm:"column:IS_MESSAGE_ALLOWED;type:bit;not null" json:"IS_MESSAGE_ALLOWED"`
	IDENTITYRESTRICTION    bool   `gorm:"column:IDENTITY_RESTRICTION;type:bit;not null" json:"IDENTITY_RESTRICTION"`
	DEPARTMENTRESTRICTION  bool   `gorm:"column:DEPARTMENT_RESTRICTION;type:bit;not null" json:"DEPARTMENT_RESTRICTION"`
}

// TableName VMSRESTRICTIONSETTING's table name
func (*VMSRESTRICTIONSETTING) TableName() string {
	return TableNameVMSRESTRICTIONSETTING
}
