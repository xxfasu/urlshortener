// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/xxfasu/urlshortener/internal/model"
)

func newURL(db *gorm.DB, opts ...gen.DOOption) uRL {
	_uRL := uRL{}

	_uRL.uRLDo.UseDB(db, opts...)
	_uRL.uRLDo.UseModel(&model.URL{})

	tableName := _uRL.uRLDo.TableName()
	_uRL.ALL = field.NewAsterisk(tableName)
	_uRL.ID = field.NewInt64(tableName, "id")
	_uRL.UserID = field.NewInt64(tableName, "user_id")
	_uRL.OriginalURL = field.NewString(tableName, "original_url")
	_uRL.ShortCode = field.NewString(tableName, "short_code")
	_uRL.IsCustom = field.NewInt64(tableName, "is_custom")
	_uRL.Views = field.NewInt64(tableName, "views")
	_uRL.ExpiredAt = field.NewInt64(tableName, "expired_at")
	_uRL.CreatedAt = field.NewInt64(tableName, "created_at")
	_uRL.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_uRL.DeletedAt = field.NewField(tableName, "deleted_at")

	_uRL.fillFieldMap()

	return _uRL
}

type uRL struct {
	uRLDo uRLDo

	ALL         field.Asterisk
	ID          field.Int64
	UserID      field.Int64
	OriginalURL field.String
	ShortCode   field.String
	IsCustom    field.Int64
	Views       field.Int64
	ExpiredAt   field.Int64
	CreatedAt   field.Int64
	UpdatedAt   field.Int64
	DeletedAt   field.Field

	fieldMap map[string]field.Expr
}

func (u uRL) Table(newTableName string) *uRL {
	u.uRLDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u uRL) As(alias string) *uRL {
	u.uRLDo.DO = *(u.uRLDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *uRL) updateTableName(table string) *uRL {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UserID = field.NewInt64(table, "user_id")
	u.OriginalURL = field.NewString(table, "original_url")
	u.ShortCode = field.NewString(table, "short_code")
	u.IsCustom = field.NewInt64(table, "is_custom")
	u.Views = field.NewInt64(table, "views")
	u.ExpiredAt = field.NewInt64(table, "expired_at")
	u.CreatedAt = field.NewInt64(table, "created_at")
	u.UpdatedAt = field.NewInt64(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *uRL) WithContext(ctx context.Context) IURLDo { return u.uRLDo.WithContext(ctx) }

func (u uRL) TableName() string { return u.uRLDo.TableName() }

func (u uRL) Alias() string { return u.uRLDo.Alias() }

func (u uRL) Columns(cols ...field.Expr) gen.Columns { return u.uRLDo.Columns(cols...) }

func (u *uRL) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *uRL) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 10)
	u.fieldMap["id"] = u.ID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["original_url"] = u.OriginalURL
	u.fieldMap["short_code"] = u.ShortCode
	u.fieldMap["is_custom"] = u.IsCustom
	u.fieldMap["views"] = u.Views
	u.fieldMap["expired_at"] = u.ExpiredAt
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u uRL) clone(db *gorm.DB) uRL {
	u.uRLDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u uRL) replaceDB(db *gorm.DB) uRL {
	u.uRLDo.ReplaceDB(db)
	return u
}

type uRLDo struct{ gen.DO }

type IURLDo interface {
	gen.SubQuery
	Debug() IURLDo
	WithContext(ctx context.Context) IURLDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IURLDo
	WriteDB() IURLDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IURLDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IURLDo
	Not(conds ...gen.Condition) IURLDo
	Or(conds ...gen.Condition) IURLDo
	Select(conds ...field.Expr) IURLDo
	Where(conds ...gen.Condition) IURLDo
	Order(conds ...field.Expr) IURLDo
	Distinct(cols ...field.Expr) IURLDo
	Omit(cols ...field.Expr) IURLDo
	Join(table schema.Tabler, on ...field.Expr) IURLDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IURLDo
	RightJoin(table schema.Tabler, on ...field.Expr) IURLDo
	Group(cols ...field.Expr) IURLDo
	Having(conds ...gen.Condition) IURLDo
	Limit(limit int) IURLDo
	Offset(offset int) IURLDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IURLDo
	Unscoped() IURLDo
	Create(values ...*model.URL) error
	CreateInBatches(values []*model.URL, batchSize int) error
	Save(values ...*model.URL) error
	First() (*model.URL, error)
	Take() (*model.URL, error)
	Last() (*model.URL, error)
	Find() ([]*model.URL, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.URL, err error)
	FindInBatches(result *[]*model.URL, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.URL) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IURLDo
	Assign(attrs ...field.AssignExpr) IURLDo
	Joins(fields ...field.RelationField) IURLDo
	Preload(fields ...field.RelationField) IURLDo
	FirstOrInit() (*model.URL, error)
	FirstOrCreate() (*model.URL, error)
	FindByPage(offset int, limit int) (result []*model.URL, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IURLDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u uRLDo) Debug() IURLDo {
	return u.withDO(u.DO.Debug())
}

func (u uRLDo) WithContext(ctx context.Context) IURLDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u uRLDo) ReadDB() IURLDo {
	return u.Clauses(dbresolver.Read)
}

func (u uRLDo) WriteDB() IURLDo {
	return u.Clauses(dbresolver.Write)
}

func (u uRLDo) Session(config *gorm.Session) IURLDo {
	return u.withDO(u.DO.Session(config))
}

func (u uRLDo) Clauses(conds ...clause.Expression) IURLDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u uRLDo) Returning(value interface{}, columns ...string) IURLDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u uRLDo) Not(conds ...gen.Condition) IURLDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u uRLDo) Or(conds ...gen.Condition) IURLDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u uRLDo) Select(conds ...field.Expr) IURLDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u uRLDo) Where(conds ...gen.Condition) IURLDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u uRLDo) Order(conds ...field.Expr) IURLDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u uRLDo) Distinct(cols ...field.Expr) IURLDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u uRLDo) Omit(cols ...field.Expr) IURLDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u uRLDo) Join(table schema.Tabler, on ...field.Expr) IURLDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u uRLDo) LeftJoin(table schema.Tabler, on ...field.Expr) IURLDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u uRLDo) RightJoin(table schema.Tabler, on ...field.Expr) IURLDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u uRLDo) Group(cols ...field.Expr) IURLDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u uRLDo) Having(conds ...gen.Condition) IURLDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u uRLDo) Limit(limit int) IURLDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u uRLDo) Offset(offset int) IURLDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u uRLDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IURLDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u uRLDo) Unscoped() IURLDo {
	return u.withDO(u.DO.Unscoped())
}

func (u uRLDo) Create(values ...*model.URL) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u uRLDo) CreateInBatches(values []*model.URL, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u uRLDo) Save(values ...*model.URL) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u uRLDo) First() (*model.URL, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.URL), nil
	}
}

func (u uRLDo) Take() (*model.URL, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.URL), nil
	}
}

func (u uRLDo) Last() (*model.URL, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.URL), nil
	}
}

func (u uRLDo) Find() ([]*model.URL, error) {
	result, err := u.DO.Find()
	return result.([]*model.URL), err
}

func (u uRLDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.URL, err error) {
	buf := make([]*model.URL, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u uRLDo) FindInBatches(result *[]*model.URL, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u uRLDo) Attrs(attrs ...field.AssignExpr) IURLDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u uRLDo) Assign(attrs ...field.AssignExpr) IURLDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u uRLDo) Joins(fields ...field.RelationField) IURLDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u uRLDo) Preload(fields ...field.RelationField) IURLDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u uRLDo) FirstOrInit() (*model.URL, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.URL), nil
	}
}

func (u uRLDo) FirstOrCreate() (*model.URL, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.URL), nil
	}
}

func (u uRLDo) FindByPage(offset int, limit int) (result []*model.URL, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u uRLDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u uRLDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u uRLDo) Delete(models ...*model.URL) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *uRLDo) withDO(do gen.Dao) *uRLDo {
	u.DO = *do.(*gen.DO)
	return u
}
