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

func newPANTRYINVENTORYREGISTER(db *gorm.DB, opts ...gen.DOOption) pANTRYINVENTORYREGISTER {
	_pANTRYINVENTORYREGISTER := pANTRYINVENTORYREGISTER{}

	_pANTRYINVENTORYREGISTER.pANTRYINVENTORYREGISTERDo.UseDB(db, opts...)
	_pANTRYINVENTORYREGISTER.pANTRYINVENTORYREGISTERDo.UseModel(&model.PANTRYINVENTORYREGISTER{})

	tableName := _pANTRYINVENTORYREGISTER.pANTRYINVENTORYREGISTERDo.TableName()
	_pANTRYINVENTORYREGISTER.ALL = field.NewAsterisk(tableName)
	_pANTRYINVENTORYREGISTER.SEQID = field.NewInt64(tableName, "SEQ_ID")
	_pANTRYINVENTORYREGISTER.INVENTORYDATETIME = field.NewString(tableName, "INVENTORY_DATE_TIME")
	_pANTRYINVENTORYREGISTER.ITEMID = field.NewInt64(tableName, "ITEM_ID")
	_pANTRYINVENTORYREGISTER.ITEMNAME = field.NewString(tableName, "ITEM_NAME")
	_pANTRYINVENTORYREGISTER.OPENINGQTY = field.NewFloat64(tableName, "OPENING_QTY")
	_pANTRYINVENTORYREGISTER.SPENTQTY = field.NewFloat64(tableName, "SPENT_QTY")
	_pANTRYINVENTORYREGISTER.TOTALQTY = field.NewFloat64(tableName, "TOTAL_QTY")
	_pANTRYINVENTORYREGISTER.MINIMUMQTYSET = field.NewFloat64(tableName, "MINIMUM_QTY_SET")
	_pANTRYINVENTORYREGISTER.UNITID = field.NewInt64(tableName, "UNIT_ID")
	_pANTRYINVENTORYREGISTER.UNITNAME = field.NewString(tableName, "UNIT_NAME")
	_pANTRYINVENTORYREGISTER.PANTRYID = field.NewInt64(tableName, "PANTRY_ID")

	_pANTRYINVENTORYREGISTER.fillFieldMap()

	return _pANTRYINVENTORYREGISTER
}

type pANTRYINVENTORYREGISTER struct {
	pANTRYINVENTORYREGISTERDo

	ALL               field.Asterisk
	SEQID             field.Int64
	INVENTORYDATETIME field.String
	ITEMID            field.Int64
	ITEMNAME          field.String
	OPENINGQTY        field.Float64
	SPENTQTY          field.Float64
	TOTALQTY          field.Float64
	MINIMUMQTYSET     field.Float64
	UNITID            field.Int64
	UNITNAME          field.String
	PANTRYID          field.Int64

	fieldMap map[string]field.Expr
}

func (p pANTRYINVENTORYREGISTER) Table(newTableName string) *pANTRYINVENTORYREGISTER {
	p.pANTRYINVENTORYREGISTERDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pANTRYINVENTORYREGISTER) As(alias string) *pANTRYINVENTORYREGISTER {
	p.pANTRYINVENTORYREGISTERDo.DO = *(p.pANTRYINVENTORYREGISTERDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pANTRYINVENTORYREGISTER) updateTableName(table string) *pANTRYINVENTORYREGISTER {
	p.ALL = field.NewAsterisk(table)
	p.SEQID = field.NewInt64(table, "SEQ_ID")
	p.INVENTORYDATETIME = field.NewString(table, "INVENTORY_DATE_TIME")
	p.ITEMID = field.NewInt64(table, "ITEM_ID")
	p.ITEMNAME = field.NewString(table, "ITEM_NAME")
	p.OPENINGQTY = field.NewFloat64(table, "OPENING_QTY")
	p.SPENTQTY = field.NewFloat64(table, "SPENT_QTY")
	p.TOTALQTY = field.NewFloat64(table, "TOTAL_QTY")
	p.MINIMUMQTYSET = field.NewFloat64(table, "MINIMUM_QTY_SET")
	p.UNITID = field.NewInt64(table, "UNIT_ID")
	p.UNITNAME = field.NewString(table, "UNIT_NAME")
	p.PANTRYID = field.NewInt64(table, "PANTRY_ID")

	p.fillFieldMap()

	return p
}

func (p *pANTRYINVENTORYREGISTER) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pANTRYINVENTORYREGISTER) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 11)
	p.fieldMap["SEQ_ID"] = p.SEQID
	p.fieldMap["INVENTORY_DATE_TIME"] = p.INVENTORYDATETIME
	p.fieldMap["ITEM_ID"] = p.ITEMID
	p.fieldMap["ITEM_NAME"] = p.ITEMNAME
	p.fieldMap["OPENING_QTY"] = p.OPENINGQTY
	p.fieldMap["SPENT_QTY"] = p.SPENTQTY
	p.fieldMap["TOTAL_QTY"] = p.TOTALQTY
	p.fieldMap["MINIMUM_QTY_SET"] = p.MINIMUMQTYSET
	p.fieldMap["UNIT_ID"] = p.UNITID
	p.fieldMap["UNIT_NAME"] = p.UNITNAME
	p.fieldMap["PANTRY_ID"] = p.PANTRYID
}

func (p pANTRYINVENTORYREGISTER) clone(db *gorm.DB) pANTRYINVENTORYREGISTER {
	p.pANTRYINVENTORYREGISTERDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pANTRYINVENTORYREGISTER) replaceDB(db *gorm.DB) pANTRYINVENTORYREGISTER {
	p.pANTRYINVENTORYREGISTERDo.ReplaceDB(db)
	return p
}

type pANTRYINVENTORYREGISTERDo struct{ gen.DO }

func (p pANTRYINVENTORYREGISTERDo) Debug() *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Debug())
}

func (p pANTRYINVENTORYREGISTERDo) WithContext(ctx context.Context) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pANTRYINVENTORYREGISTERDo) ReadDB() *pANTRYINVENTORYREGISTERDo {
	return p.Clauses(dbresolver.Read)
}

func (p pANTRYINVENTORYREGISTERDo) WriteDB() *pANTRYINVENTORYREGISTERDo {
	return p.Clauses(dbresolver.Write)
}

func (p pANTRYINVENTORYREGISTERDo) Session(config *gorm.Session) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Session(config))
}

func (p pANTRYINVENTORYREGISTERDo) Clauses(conds ...clause.Expression) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pANTRYINVENTORYREGISTERDo) Returning(value interface{}, columns ...string) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pANTRYINVENTORYREGISTERDo) Not(conds ...gen.Condition) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pANTRYINVENTORYREGISTERDo) Or(conds ...gen.Condition) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pANTRYINVENTORYREGISTERDo) Select(conds ...field.Expr) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pANTRYINVENTORYREGISTERDo) Where(conds ...gen.Condition) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pANTRYINVENTORYREGISTERDo) Order(conds ...field.Expr) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pANTRYINVENTORYREGISTERDo) Distinct(cols ...field.Expr) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pANTRYINVENTORYREGISTERDo) Omit(cols ...field.Expr) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pANTRYINVENTORYREGISTERDo) Join(table schema.Tabler, on ...field.Expr) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pANTRYINVENTORYREGISTERDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pANTRYINVENTORYREGISTERDo) RightJoin(table schema.Tabler, on ...field.Expr) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pANTRYINVENTORYREGISTERDo) Group(cols ...field.Expr) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pANTRYINVENTORYREGISTERDo) Having(conds ...gen.Condition) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pANTRYINVENTORYREGISTERDo) Limit(limit int) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pANTRYINVENTORYREGISTERDo) Offset(offset int) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pANTRYINVENTORYREGISTERDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pANTRYINVENTORYREGISTERDo) Unscoped() *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pANTRYINVENTORYREGISTERDo) Create(values ...*model.PANTRYINVENTORYREGISTER) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pANTRYINVENTORYREGISTERDo) CreateInBatches(values []*model.PANTRYINVENTORYREGISTER, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pANTRYINVENTORYREGISTERDo) Save(values ...*model.PANTRYINVENTORYREGISTER) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pANTRYINVENTORYREGISTERDo) First() (*model.PANTRYINVENTORYREGISTER, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYINVENTORYREGISTER), nil
	}
}

func (p pANTRYINVENTORYREGISTERDo) Take() (*model.PANTRYINVENTORYREGISTER, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYINVENTORYREGISTER), nil
	}
}

func (p pANTRYINVENTORYREGISTERDo) Last() (*model.PANTRYINVENTORYREGISTER, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYINVENTORYREGISTER), nil
	}
}

func (p pANTRYINVENTORYREGISTERDo) Find() ([]*model.PANTRYINVENTORYREGISTER, error) {
	result, err := p.DO.Find()
	return result.([]*model.PANTRYINVENTORYREGISTER), err
}

func (p pANTRYINVENTORYREGISTERDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PANTRYINVENTORYREGISTER, err error) {
	buf := make([]*model.PANTRYINVENTORYREGISTER, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pANTRYINVENTORYREGISTERDo) FindInBatches(result *[]*model.PANTRYINVENTORYREGISTER, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pANTRYINVENTORYREGISTERDo) Attrs(attrs ...field.AssignExpr) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pANTRYINVENTORYREGISTERDo) Assign(attrs ...field.AssignExpr) *pANTRYINVENTORYREGISTERDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pANTRYINVENTORYREGISTERDo) Joins(fields ...field.RelationField) *pANTRYINVENTORYREGISTERDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pANTRYINVENTORYREGISTERDo) Preload(fields ...field.RelationField) *pANTRYINVENTORYREGISTERDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pANTRYINVENTORYREGISTERDo) FirstOrInit() (*model.PANTRYINVENTORYREGISTER, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYINVENTORYREGISTER), nil
	}
}

func (p pANTRYINVENTORYREGISTERDo) FirstOrCreate() (*model.PANTRYINVENTORYREGISTER, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYINVENTORYREGISTER), nil
	}
}

func (p pANTRYINVENTORYREGISTERDo) FindByPage(offset int, limit int) (result []*model.PANTRYINVENTORYREGISTER, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pANTRYINVENTORYREGISTERDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pANTRYINVENTORYREGISTERDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pANTRYINVENTORYREGISTERDo) Delete(models ...*model.PANTRYINVENTORYREGISTER) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pANTRYINVENTORYREGISTERDo) withDO(do gen.Dao) *pANTRYINVENTORYREGISTERDo {
	p.DO = *do.(*gen.DO)
	return p
}
