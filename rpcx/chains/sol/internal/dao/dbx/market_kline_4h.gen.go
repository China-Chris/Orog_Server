// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dbx

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/simance-ai/smdx/rpcx/chains/sol/internal/dao/model"
)

func newMarketKline4h(db *gorm.DB, opts ...gen.DOOption) marketKline4h {
	_marketKline4h := marketKline4h{}

	_marketKline4h.marketKline4hDo.UseDB(db, opts...)
	_marketKline4h.marketKline4hDo.UseModel(&model.MarketKline4h{})

	tableName := _marketKline4h.marketKline4hDo.TableName()
	_marketKline4h.ALL = field.NewAsterisk(tableName)
	_marketKline4h.MarketAddress = field.NewString(tableName, "market_address")
	_marketKline4h.O = field.NewFloat64(tableName, "o")
	_marketKline4h.H = field.NewFloat64(tableName, "h")
	_marketKline4h.L = field.NewFloat64(tableName, "l")
	_marketKline4h.C = field.NewFloat64(tableName, "c")
	_marketKline4h.V = field.NewFloat64(tableName, "v")
	_marketKline4h.Timestamp = field.NewTime(tableName, "timestamp")
	_marketKline4h.UpdatedAt = field.NewTime(tableName, "updated_at")
	_marketKline4h.CreatedAt = field.NewTime(tableName, "created_at")

	_marketKline4h.fillFieldMap()

	return _marketKline4h
}

type marketKline4h struct {
	marketKline4hDo marketKline4hDo

	ALL           field.Asterisk
	MarketAddress field.String
	O             field.Float64
	H             field.Float64
	L             field.Float64
	C             field.Float64
	V             field.Float64
	Timestamp     field.Time
	UpdatedAt     field.Time
	CreatedAt     field.Time

	fieldMap map[string]field.Expr
}

func (m marketKline4h) Table(newTableName string) *marketKline4h {
	m.marketKline4hDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m marketKline4h) As(alias string) *marketKline4h {
	m.marketKline4hDo.DO = *(m.marketKline4hDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *marketKline4h) updateTableName(table string) *marketKline4h {
	m.ALL = field.NewAsterisk(table)
	m.MarketAddress = field.NewString(table, "market_address")
	m.O = field.NewFloat64(table, "o")
	m.H = field.NewFloat64(table, "h")
	m.L = field.NewFloat64(table, "l")
	m.C = field.NewFloat64(table, "c")
	m.V = field.NewFloat64(table, "v")
	m.Timestamp = field.NewTime(table, "timestamp")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.CreatedAt = field.NewTime(table, "created_at")

	m.fillFieldMap()

	return m
}

func (m *marketKline4h) WithContext(ctx context.Context) *marketKline4hDo {
	return m.marketKline4hDo.WithContext(ctx)
}

func (m marketKline4h) TableName() string { return m.marketKline4hDo.TableName() }

func (m marketKline4h) Alias() string { return m.marketKline4hDo.Alias() }

func (m marketKline4h) Columns(cols ...field.Expr) gen.Columns {
	return m.marketKline4hDo.Columns(cols...)
}

func (m *marketKline4h) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *marketKline4h) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 9)
	m.fieldMap["market_address"] = m.MarketAddress
	m.fieldMap["o"] = m.O
	m.fieldMap["h"] = m.H
	m.fieldMap["l"] = m.L
	m.fieldMap["c"] = m.C
	m.fieldMap["v"] = m.V
	m.fieldMap["timestamp"] = m.Timestamp
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["created_at"] = m.CreatedAt
}

func (m marketKline4h) clone(db *gorm.DB) marketKline4h {
	m.marketKline4hDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m marketKline4h) replaceDB(db *gorm.DB) marketKline4h {
	m.marketKline4hDo.ReplaceDB(db)
	return m
}

type marketKline4hDo struct{ gen.DO }

func (m marketKline4hDo) Debug() *marketKline4hDo {
	return m.withDO(m.DO.Debug())
}

func (m marketKline4hDo) WithContext(ctx context.Context) *marketKline4hDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m marketKline4hDo) ReadDB() *marketKline4hDo {
	return m.Clauses(dbresolver.Read)
}

func (m marketKline4hDo) WriteDB() *marketKline4hDo {
	return m.Clauses(dbresolver.Write)
}

func (m marketKline4hDo) Session(config *gorm.Session) *marketKline4hDo {
	return m.withDO(m.DO.Session(config))
}

func (m marketKline4hDo) Clauses(conds ...clause.Expression) *marketKline4hDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m marketKline4hDo) Returning(value interface{}, columns ...string) *marketKline4hDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m marketKline4hDo) Not(conds ...gen.Condition) *marketKline4hDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m marketKline4hDo) Or(conds ...gen.Condition) *marketKline4hDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m marketKline4hDo) Select(conds ...field.Expr) *marketKline4hDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m marketKline4hDo) Where(conds ...gen.Condition) *marketKline4hDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m marketKline4hDo) Order(conds ...field.Expr) *marketKline4hDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m marketKline4hDo) Distinct(cols ...field.Expr) *marketKline4hDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m marketKline4hDo) Omit(cols ...field.Expr) *marketKline4hDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m marketKline4hDo) Join(table schema.Tabler, on ...field.Expr) *marketKline4hDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m marketKline4hDo) LeftJoin(table schema.Tabler, on ...field.Expr) *marketKline4hDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m marketKline4hDo) RightJoin(table schema.Tabler, on ...field.Expr) *marketKline4hDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m marketKline4hDo) Group(cols ...field.Expr) *marketKline4hDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m marketKline4hDo) Having(conds ...gen.Condition) *marketKline4hDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m marketKline4hDo) Limit(limit int) *marketKline4hDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m marketKline4hDo) Offset(offset int) *marketKline4hDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m marketKline4hDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *marketKline4hDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m marketKline4hDo) Unscoped() *marketKline4hDo {
	return m.withDO(m.DO.Unscoped())
}

func (m marketKline4hDo) Create(values ...*model.MarketKline4h) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m marketKline4hDo) CreateInBatches(values []*model.MarketKline4h, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m marketKline4hDo) Save(values ...*model.MarketKline4h) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m marketKline4hDo) First() (*model.MarketKline4h, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketKline4h), nil
	}
}

func (m marketKline4hDo) Take() (*model.MarketKline4h, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketKline4h), nil
	}
}

func (m marketKline4hDo) Last() (*model.MarketKline4h, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketKline4h), nil
	}
}

func (m marketKline4hDo) Find() ([]*model.MarketKline4h, error) {
	result, err := m.DO.Find()
	return result.([]*model.MarketKline4h), err
}

func (m marketKline4hDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MarketKline4h, err error) {
	buf := make([]*model.MarketKline4h, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m marketKline4hDo) FindInBatches(result *[]*model.MarketKline4h, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m marketKline4hDo) Attrs(attrs ...field.AssignExpr) *marketKline4hDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m marketKline4hDo) Assign(attrs ...field.AssignExpr) *marketKline4hDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m marketKline4hDo) Joins(fields ...field.RelationField) *marketKline4hDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m marketKline4hDo) Preload(fields ...field.RelationField) *marketKline4hDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m marketKline4hDo) FirstOrInit() (*model.MarketKline4h, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketKline4h), nil
	}
}

func (m marketKline4hDo) FirstOrCreate() (*model.MarketKline4h, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketKline4h), nil
	}
}

func (m marketKline4hDo) FindByPage(offset int, limit int) (result []*model.MarketKline4h, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m marketKline4hDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m marketKline4hDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m marketKline4hDo) Delete(models ...*model.MarketKline4h) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *marketKline4hDo) withDO(do gen.Dao) *marketKline4hDo {
	m.DO = *do.(*gen.DO)
	return m
}
