// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameORGCOMPANY = "ORG_COMPANY"

// ORGCOMPANY mapped from table <ORG_COMPANY>
type ORGCOMPANY struct {
	COMPANYCODE    string  `gorm:"column:COMPANY_CODE;type:char;primaryKey" json:"COMPANY_CODE"`
	COMPANYNAME    string  `gorm:"column:COMPANY_NAME;type:nvarchar;not null;index:IDX_COMPANY_NAME,priority:1" json:"COMPANY_NAME"`
	COMPANYADDRESS *string `gorm:"column:COMPANY_ADDRESS;type:nvarchar" json:"COMPANY_ADDRESS"`
	COMPANYCONTACT *string `gorm:"column:COMPANY_CONTACT;type:nvarchar" json:"COMPANY_CONTACT"`
	COMPANYEMAIL   *string `gorm:"column:COMPANY_EMAIL;type:nvarchar" json:"COMPANY_EMAIL"`
	COMPANYWEBSITE *string `gorm:"column:COMPANY_WEBSITE;type:nvarchar" json:"COMPANY_WEBSITE"`
	COMPANYLOGO    *string `gorm:"column:COMPANY_LOGO;type:nvarchar" json:"COMPANY_LOGO"`
}

// TableName ORGCOMPANY's table name
func (*ORGCOMPANY) TableName() string {
	return TableNameORGCOMPANY
}
