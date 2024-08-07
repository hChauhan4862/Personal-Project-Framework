// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"SMART_OFFICE_APP/pkg/db/model"
)

func newVMSDISPLAY(db *gorm.DB, opts ...gen.DOOption) vMSDISPLAY {
	_vMSDISPLAY := vMSDISPLAY{}

	_vMSDISPLAY.vMSDISPLAYDo.UseDB(db, opts...)
	_vMSDISPLAY.vMSDISPLAYDo.UseModel(&model.VMSDISPLAY{})

	tableName := _vMSDISPLAY.vMSDISPLAYDo.TableName()
	_vMSDISPLAY.ALL = field.NewAsterisk(tableName)
	_vMSDISPLAY.DOORDISPLAYID = field.NewInt64(tableName, "DOOR_DISPLAY_ID")
	_vMSDISPLAY.DOORDISPLAYNO = field.NewString(tableName, "DOOR_DISPLAY_NO")
	_vMSDISPLAY.DOORDISPLAYNAME = field.NewString(tableName, "DOOR_DISPLAY_NAME")
	_vMSDISPLAY.DOORDISPLAYLOCATION = field.NewString(tableName, "DOOR_DISPLAY_LOCATION")
	_vMSDISPLAY.OFFICEID = field.NewInt64(tableName, "OFFICE_ID")
	_vMSDISPLAY.FLOORID = field.NewInt64(tableName, "FLOOR_ID")
	_vMSDISPLAY.ROOMID = field.NewInt64(tableName, "ROOM_ID")
	_vMSDISPLAY.PANTRYID = field.NewInt64(tableName, "PANTRY_ID")
	_vMSDISPLAY.MACHINEID = field.NewString(tableName, "MACHINE_ID")
	_vMSDISPLAY.IPADDRESS = field.NewString(tableName, "IP_ADDRESS")
	_vMSDISPLAY.SUBNETMASK = field.NewString(tableName, "SUBNET_MASK")
	_vMSDISPLAY.MACADDRESS = field.NewString(tableName, "MAC_ADDRESS")
	_vMSDISPLAY.ISALLOWNOTICE = field.NewBool(tableName, "IS_ALLOW_NOTICE")
	_vMSDISPLAY.ISALLOWBOOKING = field.NewBool(tableName, "IS_ALLOW_BOOKING")
	_vMSDISPLAY.ISALLOWAMENITIES = field.NewBool(tableName, "IS_ALLOW_AMENITIES")
	_vMSDISPLAY.ISALLOWHOSTINFO = field.NewBool(tableName, "IS_ALLOW_HOST_INFO")
	_vMSDISPLAY.ISALLOWNEXTMEETING = field.NewBool(tableName, "IS_ALLOW_NEXT_MEETING")
	_vMSDISPLAY.ISALLOWADDINTERNALUSER = field.NewBool(tableName, "IS_ALLOW_ADD_INTERNAL_USER")
	_vMSDISPLAY.ISALLOWADDEXTERNALUSER = field.NewBool(tableName, "IS_ALLOW_ADD_EXTERNAL_USER")
	_vMSDISPLAY.ISALLOWBOOKINGCONFIRMATION = field.NewBool(tableName, "IS_ALLOW_BOOKING_CONFIRMATION")
	_vMSDISPLAY.LICENCEKEY = field.NewString(tableName, "LICENCE_KEY")
	_vMSDISPLAY.LICENCEACTIVATIONDATE = field.NewString(tableName, "LICENCE_ACTIVATION_DATE")
	_vMSDISPLAY.LICENCEEXPIRYDATE = field.NewString(tableName, "LICENCE_EXPIRY_DATE")
	_vMSDISPLAY.ISACTIVE = field.NewBool(tableName, "IS_ACTIVE")
	_vMSDISPLAY.CREATEDAT = field.NewString(tableName, "CREATED_AT")
	_vMSDISPLAY.UPDATEDAT = field.NewString(tableName, "UPDATED_AT")

	_vMSDISPLAY.fillFieldMap()

	return _vMSDISPLAY
}

type vMSDISPLAY struct {
	vMSDISPLAYDo

	ALL                        field.Asterisk
	DOORDISPLAYID              field.Int64
	DOORDISPLAYNO              field.String
	DOORDISPLAYNAME            field.String
	DOORDISPLAYLOCATION        field.String
	OFFICEID                   field.Int64
	FLOORID                    field.Int64
	ROOMID                     field.Int64
	PANTRYID                   field.Int64
	MACHINEID                  field.String
	IPADDRESS                  field.String
	SUBNETMASK                 field.String
	MACADDRESS                 field.String
	ISALLOWNOTICE              field.Bool
	ISALLOWBOOKING             field.Bool
	ISALLOWAMENITIES           field.Bool
	ISALLOWHOSTINFO            field.Bool
	ISALLOWNEXTMEETING         field.Bool
	ISALLOWADDINTERNALUSER     field.Bool
	ISALLOWADDEXTERNALUSER     field.Bool
	ISALLOWBOOKINGCONFIRMATION field.Bool
	LICENCEKEY                 field.String
	LICENCEACTIVATIONDATE      field.String
	LICENCEEXPIRYDATE          field.String
	ISACTIVE                   field.Bool
	CREATEDAT                  field.String
	UPDATEDAT                  field.String

	fieldMap map[string]field.Expr
}

func (v vMSDISPLAY) Table(newTableName string) *vMSDISPLAY {
	v.vMSDISPLAYDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v vMSDISPLAY) As(alias string) *vMSDISPLAY {
	v.vMSDISPLAYDo.DO = *(v.vMSDISPLAYDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *vMSDISPLAY) updateTableName(table string) *vMSDISPLAY {
	v.ALL = field.NewAsterisk(table)
	v.DOORDISPLAYID = field.NewInt64(table, "DOOR_DISPLAY_ID")
	v.DOORDISPLAYNO = field.NewString(table, "DOOR_DISPLAY_NO")
	v.DOORDISPLAYNAME = field.NewString(table, "DOOR_DISPLAY_NAME")
	v.DOORDISPLAYLOCATION = field.NewString(table, "DOOR_DISPLAY_LOCATION")
	v.OFFICEID = field.NewInt64(table, "OFFICE_ID")
	v.FLOORID = field.NewInt64(table, "FLOOR_ID")
	v.ROOMID = field.NewInt64(table, "ROOM_ID")
	v.PANTRYID = field.NewInt64(table, "PANTRY_ID")
	v.MACHINEID = field.NewString(table, "MACHINE_ID")
	v.IPADDRESS = field.NewString(table, "IP_ADDRESS")
	v.SUBNETMASK = field.NewString(table, "SUBNET_MASK")
	v.MACADDRESS = field.NewString(table, "MAC_ADDRESS")
	v.ISALLOWNOTICE = field.NewBool(table, "IS_ALLOW_NOTICE")
	v.ISALLOWBOOKING = field.NewBool(table, "IS_ALLOW_BOOKING")
	v.ISALLOWAMENITIES = field.NewBool(table, "IS_ALLOW_AMENITIES")
	v.ISALLOWHOSTINFO = field.NewBool(table, "IS_ALLOW_HOST_INFO")
	v.ISALLOWNEXTMEETING = field.NewBool(table, "IS_ALLOW_NEXT_MEETING")
	v.ISALLOWADDINTERNALUSER = field.NewBool(table, "IS_ALLOW_ADD_INTERNAL_USER")
	v.ISALLOWADDEXTERNALUSER = field.NewBool(table, "IS_ALLOW_ADD_EXTERNAL_USER")
	v.ISALLOWBOOKINGCONFIRMATION = field.NewBool(table, "IS_ALLOW_BOOKING_CONFIRMATION")
	v.LICENCEKEY = field.NewString(table, "LICENCE_KEY")
	v.LICENCEACTIVATIONDATE = field.NewString(table, "LICENCE_ACTIVATION_DATE")
	v.LICENCEEXPIRYDATE = field.NewString(table, "LICENCE_EXPIRY_DATE")
	v.ISACTIVE = field.NewBool(table, "IS_ACTIVE")
	v.CREATEDAT = field.NewString(table, "CREATED_AT")
	v.UPDATEDAT = field.NewString(table, "UPDATED_AT")

	v.fillFieldMap()

	return v
}

func (v *vMSDISPLAY) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *vMSDISPLAY) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 26)
	v.fieldMap["DOOR_DISPLAY_ID"] = v.DOORDISPLAYID
	v.fieldMap["DOOR_DISPLAY_NO"] = v.DOORDISPLAYNO
	v.fieldMap["DOOR_DISPLAY_NAME"] = v.DOORDISPLAYNAME
	v.fieldMap["DOOR_DISPLAY_LOCATION"] = v.DOORDISPLAYLOCATION
	v.fieldMap["OFFICE_ID"] = v.OFFICEID
	v.fieldMap["FLOOR_ID"] = v.FLOORID
	v.fieldMap["ROOM_ID"] = v.ROOMID
	v.fieldMap["PANTRY_ID"] = v.PANTRYID
	v.fieldMap["MACHINE_ID"] = v.MACHINEID
	v.fieldMap["IP_ADDRESS"] = v.IPADDRESS
	v.fieldMap["SUBNET_MASK"] = v.SUBNETMASK
	v.fieldMap["MAC_ADDRESS"] = v.MACADDRESS
	v.fieldMap["IS_ALLOW_NOTICE"] = v.ISALLOWNOTICE
	v.fieldMap["IS_ALLOW_BOOKING"] = v.ISALLOWBOOKING
	v.fieldMap["IS_ALLOW_AMENITIES"] = v.ISALLOWAMENITIES
	v.fieldMap["IS_ALLOW_HOST_INFO"] = v.ISALLOWHOSTINFO
	v.fieldMap["IS_ALLOW_NEXT_MEETING"] = v.ISALLOWNEXTMEETING
	v.fieldMap["IS_ALLOW_ADD_INTERNAL_USER"] = v.ISALLOWADDINTERNALUSER
	v.fieldMap["IS_ALLOW_ADD_EXTERNAL_USER"] = v.ISALLOWADDEXTERNALUSER
	v.fieldMap["IS_ALLOW_BOOKING_CONFIRMATION"] = v.ISALLOWBOOKINGCONFIRMATION
	v.fieldMap["LICENCE_KEY"] = v.LICENCEKEY
	v.fieldMap["LICENCE_ACTIVATION_DATE"] = v.LICENCEACTIVATIONDATE
	v.fieldMap["LICENCE_EXPIRY_DATE"] = v.LICENCEEXPIRYDATE
	v.fieldMap["IS_ACTIVE"] = v.ISACTIVE
	v.fieldMap["CREATED_AT"] = v.CREATEDAT
	v.fieldMap["UPDATED_AT"] = v.UPDATEDAT
}

func (v vMSDISPLAY) clone(db *gorm.DB) vMSDISPLAY {
	v.vMSDISPLAYDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v vMSDISPLAY) replaceDB(db *gorm.DB) vMSDISPLAY {
	v.vMSDISPLAYDo.ReplaceDB(db)
	return v
}

type vMSDISPLAYDo struct{ gen.DO }

func (v vMSDISPLAYDo) Debug() *vMSDISPLAYDo {
	return v.withDO(v.DO.Debug())
}

func (v vMSDISPLAYDo) WithContext(ctx context.Context) *vMSDISPLAYDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v vMSDISPLAYDo) ReadDB() *vMSDISPLAYDo {
	return v.Clauses(dbresolver.Read)
}

func (v vMSDISPLAYDo) WriteDB() *vMSDISPLAYDo {
	return v.Clauses(dbresolver.Write)
}

func (v vMSDISPLAYDo) Session(config *gorm.Session) *vMSDISPLAYDo {
	return v.withDO(v.DO.Session(config))
}

func (v vMSDISPLAYDo) Clauses(conds ...clause.Expression) *vMSDISPLAYDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v vMSDISPLAYDo) Returning(value interface{}, columns ...string) *vMSDISPLAYDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v vMSDISPLAYDo) Not(conds ...gen.Condition) *vMSDISPLAYDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v vMSDISPLAYDo) Or(conds ...gen.Condition) *vMSDISPLAYDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v vMSDISPLAYDo) Select(conds ...field.Expr) *vMSDISPLAYDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v vMSDISPLAYDo) Where(conds ...gen.Condition) *vMSDISPLAYDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v vMSDISPLAYDo) Order(conds ...field.Expr) *vMSDISPLAYDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v vMSDISPLAYDo) Distinct(cols ...field.Expr) *vMSDISPLAYDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v vMSDISPLAYDo) Omit(cols ...field.Expr) *vMSDISPLAYDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v vMSDISPLAYDo) Join(table schema.Tabler, on ...field.Expr) *vMSDISPLAYDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v vMSDISPLAYDo) LeftJoin(table schema.Tabler, on ...field.Expr) *vMSDISPLAYDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v vMSDISPLAYDo) RightJoin(table schema.Tabler, on ...field.Expr) *vMSDISPLAYDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v vMSDISPLAYDo) Group(cols ...field.Expr) *vMSDISPLAYDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v vMSDISPLAYDo) Having(conds ...gen.Condition) *vMSDISPLAYDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v vMSDISPLAYDo) Limit(limit int) *vMSDISPLAYDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v vMSDISPLAYDo) Offset(offset int) *vMSDISPLAYDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v vMSDISPLAYDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *vMSDISPLAYDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v vMSDISPLAYDo) Unscoped() *vMSDISPLAYDo {
	return v.withDO(v.DO.Unscoped())
}

func (v vMSDISPLAYDo) Create(values ...*model.VMSDISPLAY) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v vMSDISPLAYDo) CreateInBatches(values []*model.VMSDISPLAY, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v vMSDISPLAYDo) Save(values ...*model.VMSDISPLAY) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v vMSDISPLAYDo) First() (*model.VMSDISPLAY, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VMSDISPLAY), nil
	}
}

func (v vMSDISPLAYDo) Take() (*model.VMSDISPLAY, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VMSDISPLAY), nil
	}
}

func (v vMSDISPLAYDo) Last() (*model.VMSDISPLAY, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VMSDISPLAY), nil
	}
}

func (v vMSDISPLAYDo) Find() ([]*model.VMSDISPLAY, error) {
	result, err := v.DO.Find()
	return result.([]*model.VMSDISPLAY), err
}

func (v vMSDISPLAYDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VMSDISPLAY, err error) {
	buf := make([]*model.VMSDISPLAY, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v vMSDISPLAYDo) FindInBatches(result *[]*model.VMSDISPLAY, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v vMSDISPLAYDo) Attrs(attrs ...field.AssignExpr) *vMSDISPLAYDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v vMSDISPLAYDo) Assign(attrs ...field.AssignExpr) *vMSDISPLAYDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v vMSDISPLAYDo) Joins(fields ...field.RelationField) *vMSDISPLAYDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v vMSDISPLAYDo) Preload(fields ...field.RelationField) *vMSDISPLAYDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v vMSDISPLAYDo) FirstOrInit() (*model.VMSDISPLAY, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VMSDISPLAY), nil
	}
}

func (v vMSDISPLAYDo) FirstOrCreate() (*model.VMSDISPLAY, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VMSDISPLAY), nil
	}
}

func (v vMSDISPLAYDo) FindByPage(offset int, limit int) (result []*model.VMSDISPLAY, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v vMSDISPLAYDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v vMSDISPLAYDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v vMSDISPLAYDo) Delete(models ...*model.VMSDISPLAY) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *vMSDISPLAYDo) withDO(do gen.Dao) *vMSDISPLAYDo {
	v.DO = *do.(*gen.DO)
	return v
}
