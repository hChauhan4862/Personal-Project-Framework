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

func newPANTRYITEMSCAT(db *gorm.DB, opts ...gen.DOOption) pANTRYITEMSCAT {
	_pANTRYITEMSCAT := pANTRYITEMSCAT{}

	_pANTRYITEMSCAT.pANTRYITEMSCATDo.UseDB(db, opts...)
	_pANTRYITEMSCAT.pANTRYITEMSCATDo.UseModel(&model.PANTRYITEMSCAT{})

	tableName := _pANTRYITEMSCAT.pANTRYITEMSCATDo.TableName()
	_pANTRYITEMSCAT.ALL = field.NewAsterisk(tableName)
	_pANTRYITEMSCAT.CATSEQ = field.NewInt64(tableName, "CAT_SEQ")
	_pANTRYITEMSCAT.CATNAME = field.NewString(tableName, "CAT_NAME")
	_pANTRYITEMSCAT.CATACTIVE = field.NewBool(tableName, "CAT_ACTIVE")

	_pANTRYITEMSCAT.fillFieldMap()

	return _pANTRYITEMSCAT
}

type pANTRYITEMSCAT struct {
	pANTRYITEMSCATDo

	ALL       field.Asterisk
	CATSEQ    field.Int64
	CATNAME   field.String
	CATACTIVE field.Bool

	fieldMap map[string]field.Expr
}

func (p pANTRYITEMSCAT) Table(newTableName string) *pANTRYITEMSCAT {
	p.pANTRYITEMSCATDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pANTRYITEMSCAT) As(alias string) *pANTRYITEMSCAT {
	p.pANTRYITEMSCATDo.DO = *(p.pANTRYITEMSCATDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pANTRYITEMSCAT) updateTableName(table string) *pANTRYITEMSCAT {
	p.ALL = field.NewAsterisk(table)
	p.CATSEQ = field.NewInt64(table, "CAT_SEQ")
	p.CATNAME = field.NewString(table, "CAT_NAME")
	p.CATACTIVE = field.NewBool(table, "CAT_ACTIVE")

	p.fillFieldMap()

	return p
}

func (p *pANTRYITEMSCAT) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pANTRYITEMSCAT) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 3)
	p.fieldMap["CAT_SEQ"] = p.CATSEQ
	p.fieldMap["CAT_NAME"] = p.CATNAME
	p.fieldMap["CAT_ACTIVE"] = p.CATACTIVE
}

func (p pANTRYITEMSCAT) clone(db *gorm.DB) pANTRYITEMSCAT {
	p.pANTRYITEMSCATDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pANTRYITEMSCAT) replaceDB(db *gorm.DB) pANTRYITEMSCAT {
	p.pANTRYITEMSCATDo.ReplaceDB(db)
	return p
}

type pANTRYITEMSCATDo struct{ gen.DO }

func (p pANTRYITEMSCATDo) Debug() *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Debug())
}

func (p pANTRYITEMSCATDo) WithContext(ctx context.Context) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pANTRYITEMSCATDo) ReadDB() *pANTRYITEMSCATDo {
	return p.Clauses(dbresolver.Read)
}

func (p pANTRYITEMSCATDo) WriteDB() *pANTRYITEMSCATDo {
	return p.Clauses(dbresolver.Write)
}

func (p pANTRYITEMSCATDo) Session(config *gorm.Session) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Session(config))
}

func (p pANTRYITEMSCATDo) Clauses(conds ...clause.Expression) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pANTRYITEMSCATDo) Returning(value interface{}, columns ...string) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pANTRYITEMSCATDo) Not(conds ...gen.Condition) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pANTRYITEMSCATDo) Or(conds ...gen.Condition) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pANTRYITEMSCATDo) Select(conds ...field.Expr) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pANTRYITEMSCATDo) Where(conds ...gen.Condition) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pANTRYITEMSCATDo) Order(conds ...field.Expr) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pANTRYITEMSCATDo) Distinct(cols ...field.Expr) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pANTRYITEMSCATDo) Omit(cols ...field.Expr) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pANTRYITEMSCATDo) Join(table schema.Tabler, on ...field.Expr) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pANTRYITEMSCATDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pANTRYITEMSCATDo) RightJoin(table schema.Tabler, on ...field.Expr) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pANTRYITEMSCATDo) Group(cols ...field.Expr) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pANTRYITEMSCATDo) Having(conds ...gen.Condition) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pANTRYITEMSCATDo) Limit(limit int) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pANTRYITEMSCATDo) Offset(offset int) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pANTRYITEMSCATDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pANTRYITEMSCATDo) Unscoped() *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pANTRYITEMSCATDo) Create(values ...*model.PANTRYITEMSCAT) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pANTRYITEMSCATDo) CreateInBatches(values []*model.PANTRYITEMSCAT, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pANTRYITEMSCATDo) Save(values ...*model.PANTRYITEMSCAT) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pANTRYITEMSCATDo) First() (*model.PANTRYITEMSCAT, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYITEMSCAT), nil
	}
}

func (p pANTRYITEMSCATDo) Take() (*model.PANTRYITEMSCAT, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYITEMSCAT), nil
	}
}

func (p pANTRYITEMSCATDo) Last() (*model.PANTRYITEMSCAT, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYITEMSCAT), nil
	}
}

func (p pANTRYITEMSCATDo) Find() ([]*model.PANTRYITEMSCAT, error) {
	result, err := p.DO.Find()
	return result.([]*model.PANTRYITEMSCAT), err
}

func (p pANTRYITEMSCATDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PANTRYITEMSCAT, err error) {
	buf := make([]*model.PANTRYITEMSCAT, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pANTRYITEMSCATDo) FindInBatches(result *[]*model.PANTRYITEMSCAT, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pANTRYITEMSCATDo) Attrs(attrs ...field.AssignExpr) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pANTRYITEMSCATDo) Assign(attrs ...field.AssignExpr) *pANTRYITEMSCATDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pANTRYITEMSCATDo) Joins(fields ...field.RelationField) *pANTRYITEMSCATDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pANTRYITEMSCATDo) Preload(fields ...field.RelationField) *pANTRYITEMSCATDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pANTRYITEMSCATDo) FirstOrInit() (*model.PANTRYITEMSCAT, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYITEMSCAT), nil
	}
}

func (p pANTRYITEMSCATDo) FirstOrCreate() (*model.PANTRYITEMSCAT, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYITEMSCAT), nil
	}
}

func (p pANTRYITEMSCATDo) FindByPage(offset int, limit int) (result []*model.PANTRYITEMSCAT, count int64, err error) {
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

func (p pANTRYITEMSCATDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pANTRYITEMSCATDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pANTRYITEMSCATDo) Delete(models ...*model.PANTRYITEMSCAT) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pANTRYITEMSCATDo) withDO(do gen.Dao) *pANTRYITEMSCATDo {
	p.DO = *do.(*gen.DO)
	return p
}
