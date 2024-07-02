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

func newAPPRESPONSECODESCAT(db *gorm.DB, opts ...gen.DOOption) aPPRESPONSECODESCAT {
	_aPPRESPONSECODESCAT := aPPRESPONSECODESCAT{}

	_aPPRESPONSECODESCAT.aPPRESPONSECODESCATDo.UseDB(db, opts...)
	_aPPRESPONSECODESCAT.aPPRESPONSECODESCATDo.UseModel(&model.APPRESPONSECODESCAT{})

	tableName := _aPPRESPONSECODESCAT.aPPRESPONSECODESCATDo.TableName()
	_aPPRESPONSECODESCAT.ALL = field.NewAsterisk(tableName)
	_aPPRESPONSECODESCAT.SEQID = field.NewInt64(tableName, "SEQ_ID")
	_aPPRESPONSECODESCAT.NAME = field.NewString(tableName, "NAME")

	_aPPRESPONSECODESCAT.fillFieldMap()

	return _aPPRESPONSECODESCAT
}

type aPPRESPONSECODESCAT struct {
	aPPRESPONSECODESCATDo

	ALL   field.Asterisk
	SEQID field.Int64
	NAME  field.String

	fieldMap map[string]field.Expr
}

func (a aPPRESPONSECODESCAT) Table(newTableName string) *aPPRESPONSECODESCAT {
	a.aPPRESPONSECODESCATDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a aPPRESPONSECODESCAT) As(alias string) *aPPRESPONSECODESCAT {
	a.aPPRESPONSECODESCATDo.DO = *(a.aPPRESPONSECODESCATDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *aPPRESPONSECODESCAT) updateTableName(table string) *aPPRESPONSECODESCAT {
	a.ALL = field.NewAsterisk(table)
	a.SEQID = field.NewInt64(table, "SEQ_ID")
	a.NAME = field.NewString(table, "NAME")

	a.fillFieldMap()

	return a
}

func (a *aPPRESPONSECODESCAT) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *aPPRESPONSECODESCAT) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 2)
	a.fieldMap["SEQ_ID"] = a.SEQID
	a.fieldMap["NAME"] = a.NAME
}

func (a aPPRESPONSECODESCAT) clone(db *gorm.DB) aPPRESPONSECODESCAT {
	a.aPPRESPONSECODESCATDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a aPPRESPONSECODESCAT) replaceDB(db *gorm.DB) aPPRESPONSECODESCAT {
	a.aPPRESPONSECODESCATDo.ReplaceDB(db)
	return a
}

type aPPRESPONSECODESCATDo struct{ gen.DO }

func (a aPPRESPONSECODESCATDo) Debug() *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Debug())
}

func (a aPPRESPONSECODESCATDo) WithContext(ctx context.Context) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a aPPRESPONSECODESCATDo) ReadDB() *aPPRESPONSECODESCATDo {
	return a.Clauses(dbresolver.Read)
}

func (a aPPRESPONSECODESCATDo) WriteDB() *aPPRESPONSECODESCATDo {
	return a.Clauses(dbresolver.Write)
}

func (a aPPRESPONSECODESCATDo) Session(config *gorm.Session) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Session(config))
}

func (a aPPRESPONSECODESCATDo) Clauses(conds ...clause.Expression) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a aPPRESPONSECODESCATDo) Returning(value interface{}, columns ...string) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a aPPRESPONSECODESCATDo) Not(conds ...gen.Condition) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a aPPRESPONSECODESCATDo) Or(conds ...gen.Condition) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a aPPRESPONSECODESCATDo) Select(conds ...field.Expr) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a aPPRESPONSECODESCATDo) Where(conds ...gen.Condition) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a aPPRESPONSECODESCATDo) Order(conds ...field.Expr) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a aPPRESPONSECODESCATDo) Distinct(cols ...field.Expr) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a aPPRESPONSECODESCATDo) Omit(cols ...field.Expr) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a aPPRESPONSECODESCATDo) Join(table schema.Tabler, on ...field.Expr) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a aPPRESPONSECODESCATDo) LeftJoin(table schema.Tabler, on ...field.Expr) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a aPPRESPONSECODESCATDo) RightJoin(table schema.Tabler, on ...field.Expr) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a aPPRESPONSECODESCATDo) Group(cols ...field.Expr) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a aPPRESPONSECODESCATDo) Having(conds ...gen.Condition) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a aPPRESPONSECODESCATDo) Limit(limit int) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a aPPRESPONSECODESCATDo) Offset(offset int) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a aPPRESPONSECODESCATDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a aPPRESPONSECODESCATDo) Unscoped() *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Unscoped())
}

func (a aPPRESPONSECODESCATDo) Create(values ...*model.APPRESPONSECODESCAT) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a aPPRESPONSECODESCATDo) CreateInBatches(values []*model.APPRESPONSECODESCAT, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a aPPRESPONSECODESCATDo) Save(values ...*model.APPRESPONSECODESCAT) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a aPPRESPONSECODESCATDo) First() (*model.APPRESPONSECODESCAT, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPRESPONSECODESCAT), nil
	}
}

func (a aPPRESPONSECODESCATDo) Take() (*model.APPRESPONSECODESCAT, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPRESPONSECODESCAT), nil
	}
}

func (a aPPRESPONSECODESCATDo) Last() (*model.APPRESPONSECODESCAT, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPRESPONSECODESCAT), nil
	}
}

func (a aPPRESPONSECODESCATDo) Find() ([]*model.APPRESPONSECODESCAT, error) {
	result, err := a.DO.Find()
	return result.([]*model.APPRESPONSECODESCAT), err
}

func (a aPPRESPONSECODESCATDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.APPRESPONSECODESCAT, err error) {
	buf := make([]*model.APPRESPONSECODESCAT, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a aPPRESPONSECODESCATDo) FindInBatches(result *[]*model.APPRESPONSECODESCAT, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a aPPRESPONSECODESCATDo) Attrs(attrs ...field.AssignExpr) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a aPPRESPONSECODESCATDo) Assign(attrs ...field.AssignExpr) *aPPRESPONSECODESCATDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a aPPRESPONSECODESCATDo) Joins(fields ...field.RelationField) *aPPRESPONSECODESCATDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a aPPRESPONSECODESCATDo) Preload(fields ...field.RelationField) *aPPRESPONSECODESCATDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a aPPRESPONSECODESCATDo) FirstOrInit() (*model.APPRESPONSECODESCAT, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPRESPONSECODESCAT), nil
	}
}

func (a aPPRESPONSECODESCATDo) FirstOrCreate() (*model.APPRESPONSECODESCAT, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPRESPONSECODESCAT), nil
	}
}

func (a aPPRESPONSECODESCATDo) FindByPage(offset int, limit int) (result []*model.APPRESPONSECODESCAT, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a aPPRESPONSECODESCATDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a aPPRESPONSECODESCATDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a aPPRESPONSECODESCATDo) Delete(models ...*model.APPRESPONSECODESCAT) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *aPPRESPONSECODESCATDo) withDO(do gen.Dao) *aPPRESPONSECODESCATDo {
	a.DO = *do.(*gen.DO)
	return a
}