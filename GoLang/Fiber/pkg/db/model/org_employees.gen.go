// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameORGEMPLOYEE = "ORG_EMPLOYEES"

// ORGEMPLOYEE mapped from table <ORG_EMPLOYEES>
type ORGEMPLOYEE struct {
	SEQID               int64   `gorm:"column:SEQ_ID;type:int;primaryKey" json:"SEQ_ID"`
	EMPUNQCD            string  `gorm:"column:EMP_UNQCD;type:nvarchar;not null;uniqueIndex:WN_INTERNAL_USERS_USER_UNQCD_key,priority:1;index:IDX_USER_UNQCD,priority:1" json:"EMP_UNQCD"`
	EMPEMAIL            string  `gorm:"column:EMP_EMAIL;type:nvarchar;not null;uniqueIndex:WN_INTERNAL_USERS_USER_EMAIL_key,priority:1;index:IDX_USER_EMAIL,priority:1" json:"EMP_EMAIL"`
	EMPNAME             string  `gorm:"column:EMP_NAME;type:nvarchar;not null;index:IDX_USER_NAME,priority:1" json:"EMP_NAME"`
	EMPCONTACTNO        *string `gorm:"column:EMP_CONTACT_NO;type:nvarchar;index:IDX_USER_CONTACT_NO,priority:1" json:"EMP_CONTACT_NO"`
	ALTERNATECONTACTNO  *string `gorm:"column:ALTERNATE_CONTACT_NO;type:nvarchar" json:"ALTERNATE_CONTACT_NO"`
	GENDER              *string `gorm:"column:GENDER;type:char" json:"GENDER"`
	IDNTCODE            *string `gorm:"column:IDNT_CODE;type:char;index:IDX_IDENTITY_CODE,priority:1;index:IDX_DEPARTMENT_CODE_IDENTITY_CODE,priority:2" json:"IDNT_CODE"`
	IDNTVALUE           *string `gorm:"column:IDNT_VALUE;type:nvarchar" json:"IDNT_VALUE"`
	DEPARTMENTCODE      *string `gorm:"column:DEPARTMENT_CODE;type:char;index:IDX_DEPARTMENT_CODE,priority:1;index:IDX_DEPARTMENT_CODE_IDENTITY_CODE,priority:1" json:"DEPARTMENT_CODE"`
	DEPARTMENTCODEVALUE *string `gorm:"column:DEPARTMENT_CODE_VALUE;type:nvarchar" json:"DEPARTMENT_CODE_VALUE"`
	LOGINID             string  `gorm:"column:LOGIN_ID;type:nvarchar;not null;uniqueIndex:WN_INTERNAL_USERS_USER_ID_key,priority:1;index:IDX_USER_ID,priority:1" json:"LOGIN_ID"`
	PASSWORD            string  `gorm:"column:PASSWORD;type:nvarchar;not null" json:"PASSWORD"`
	PROFILEPIC          *string `gorm:"column:PROFILE_PIC;type:nvarchar" json:"PROFILE_PIC"`
	ISACTIVE            *bool   `gorm:"column:IS_ACTIVE;type:bit;not null;default:1" json:"IS_ACTIVE"`
	CREATEDAT           string  `gorm:"column:CREATED_AT;type:char;not null" json:"CREATED_AT"`
	UPDATEDAT           *string `gorm:"column:UPDATED_AT;type:char" json:"UPDATED_AT"`
}

// TableName ORGEMPLOYEE's table name
func (*ORGEMPLOYEE) TableName() string {
	return TableNameORGEMPLOYEE
}