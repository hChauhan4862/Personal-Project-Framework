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

func newORGROLECODE(db *gorm.DB, opts ...gen.DOOption) oRGROLECODE {
	_oRGROLECODE := oRGROLECODE{}

	_oRGROLECODE.oRGROLECODEDo.UseDB(db, opts...)
	_oRGROLECODE.oRGROLECODEDo.UseModel(&model.ORGROLECODE{})

	tableName := _oRGROLECODE.oRGROLECODEDo.TableName()
	_oRGROLECODE.ALL = field.NewAsterisk(tableName)
	_oRGROLECODE.ROLESEQ = field.NewInt64(tableName, "ROLE_SEQ")
	_oRGROLECODE.ROLECODE = field.NewString(tableName, "ROLE_CODE")
	_oRGROLECODE.ROLENAME = field.NewString(tableName, "ROLE_NAME")

	_oRGROLECODE.fillFieldMap()

	return _oRGROLECODE
}

type oRGROLECODE struct {
	oRGROLECODEDo

	ALL      field.Asterisk
	ROLESEQ  field.Int64
	ROLECODE field.String
	ROLENAME field.String

	fieldMap map[string]field.Expr
}

func (o oRGROLECODE) Table(newTableName string) *oRGROLECODE {
	o.oRGROLECODEDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o oRGROLECODE) As(alias string) *oRGROLECODE {
	o.oRGROLECODEDo.DO = *(o.oRGROLECODEDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *oRGROLECODE) updateTableName(table string) *oRGROLECODE {
	o.ALL = field.NewAsterisk(table)
	o.ROLESEQ = field.NewInt64(table, "ROLE_SEQ")
	o.ROLECODE = field.NewString(table, "ROLE_CODE")
	o.ROLENAME = field.NewString(table, "ROLE_NAME")

	o.fillFieldMap()

	return o
}

func (o *oRGROLECODE) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *oRGROLECODE) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 3)
	o.fieldMap["ROLE_SEQ"] = o.ROLESEQ
	o.fieldMap["ROLE_CODE"] = o.ROLECODE
	o.fieldMap["ROLE_NAME"] = o.ROLENAME
}

func (o oRGROLECODE) clone(db *gorm.DB) oRGROLECODE {
	o.oRGROLECODEDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o oRGROLECODE) replaceDB(db *gorm.DB) oRGROLECODE {
	o.oRGROLECODEDo.ReplaceDB(db)
	return o
}

type oRGROLECODEDo struct{ gen.DO }

func (o oRGROLECODEDo) Debug() *oRGROLECODEDo {
	return o.withDO(o.DO.Debug())
}

func (o oRGROLECODEDo) WithContext(ctx context.Context) *oRGROLECODEDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o oRGROLECODEDo) ReadDB() *oRGROLECODEDo {
	return o.Clauses(dbresolver.Read)
}

func (o oRGROLECODEDo) WriteDB() *oRGROLECODEDo {
	return o.Clauses(dbresolver.Write)
}

func (o oRGROLECODEDo) Session(config *gorm.Session) *oRGROLECODEDo {
	return o.withDO(o.DO.Session(config))
}

func (o oRGROLECODEDo) Clauses(conds ...clause.Expression) *oRGROLECODEDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o oRGROLECODEDo) Returning(value interface{}, columns ...string) *oRGROLECODEDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o oRGROLECODEDo) Not(conds ...gen.Condition) *oRGROLECODEDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o oRGROLECODEDo) Or(conds ...gen.Condition) *oRGROLECODEDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o oRGROLECODEDo) Select(conds ...field.Expr) *oRGROLECODEDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o oRGROLECODEDo) Where(conds ...gen.Condition) *oRGROLECODEDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o oRGROLECODEDo) Order(conds ...field.Expr) *oRGROLECODEDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o oRGROLECODEDo) Distinct(cols ...field.Expr) *oRGROLECODEDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o oRGROLECODEDo) Omit(cols ...field.Expr) *oRGROLECODEDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o oRGROLECODEDo) Join(table schema.Tabler, on ...field.Expr) *oRGROLECODEDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o oRGROLECODEDo) LeftJoin(table schema.Tabler, on ...field.Expr) *oRGROLECODEDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o oRGROLECODEDo) RightJoin(table schema.Tabler, on ...field.Expr) *oRGROLECODEDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o oRGROLECODEDo) Group(cols ...field.Expr) *oRGROLECODEDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o oRGROLECODEDo) Having(conds ...gen.Condition) *oRGROLECODEDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o oRGROLECODEDo) Limit(limit int) *oRGROLECODEDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o oRGROLECODEDo) Offset(offset int) *oRGROLECODEDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o oRGROLECODEDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *oRGROLECODEDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o oRGROLECODEDo) Unscoped() *oRGROLECODEDo {
	return o.withDO(o.DO.Unscoped())
}

func (o oRGROLECODEDo) Create(values ...*model.ORGROLECODE) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o oRGROLECODEDo) CreateInBatches(values []*model.ORGROLECODE, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o oRGROLECODEDo) Save(values ...*model.ORGROLECODE) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o oRGROLECODEDo) First() (*model.ORGROLECODE, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ORGROLECODE), nil
	}
}

func (o oRGROLECODEDo) Take() (*model.ORGROLECODE, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ORGROLECODE), nil
	}
}

func (o oRGROLECODEDo) Last() (*model.ORGROLECODE, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ORGROLECODE), nil
	}
}

func (o oRGROLECODEDo) Find() ([]*model.ORGROLECODE, error) {
	result, err := o.DO.Find()
	return result.([]*model.ORGROLECODE), err
}

func (o oRGROLECODEDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ORGROLECODE, err error) {
	buf := make([]*model.ORGROLECODE, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o oRGROLECODEDo) FindInBatches(result *[]*model.ORGROLECODE, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o oRGROLECODEDo) Attrs(attrs ...field.AssignExpr) *oRGROLECODEDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o oRGROLECODEDo) Assign(attrs ...field.AssignExpr) *oRGROLECODEDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o oRGROLECODEDo) Joins(fields ...field.RelationField) *oRGROLECODEDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o oRGROLECODEDo) Preload(fields ...field.RelationField) *oRGROLECODEDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o oRGROLECODEDo) FirstOrInit() (*model.ORGROLECODE, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ORGROLECODE), nil
	}
}

func (o oRGROLECODEDo) FirstOrCreate() (*model.ORGROLECODE, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ORGROLECODE), nil
	}
}

func (o oRGROLECODEDo) FindByPage(offset int, limit int) (result []*model.ORGROLECODE, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o oRGROLECODEDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o oRGROLECODEDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o oRGROLECODEDo) Delete(models ...*model.ORGROLECODE) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *oRGROLECODEDo) withDO(do gen.Dao) *oRGROLECODEDo {
	o.DO = *do.(*gen.DO)
	return o
}