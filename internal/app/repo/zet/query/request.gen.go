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

	"github.com/anguloc/zet/internal/app/repo/zet/model"
)

func newRequest(db *gorm.DB, opts ...gen.DOOption) request {
	_request := request{}

	_request.requestDo.UseDB(db, opts...)
	_request.requestDo.UseModel(&model.Request{})

	tableName := _request.requestDo.TableName()
	_request.ALL = field.NewAsterisk(tableName)
	_request.ID = field.NewInt64(tableName, "id")
	_request.URL = field.NewString(tableName, "url")
	_request.Mark = field.NewString(tableName, "mark")
	_request.RequestNum = field.NewInt32(tableName, "request_num")
	_request.Content = field.NewString(tableName, "content")
	_request.ParseStatus = field.NewInt32(tableName, "parse_status")
	_request.CreatedAt = field.NewTime(tableName, "created_at")
	_request.UpdatedAt = field.NewTime(tableName, "updated_at")

	_request.fillFieldMap()

	return _request
}

type request struct {
	requestDo requestDo

	ALL         field.Asterisk
	ID          field.Int64  // id
	URL         field.String // 资源地址
	Mark        field.String // 标识
	RequestNum  field.Int32  // 请求次数
	Content     field.String // 请求原始结果
	ParseStatus field.Int32  // 请求内容解析状态
	CreatedAt   field.Time   // 插入时间
	UpdatedAt   field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (r request) Table(newTableName string) *request {
	r.requestDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r request) As(alias string) *request {
	r.requestDo.DO = *(r.requestDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *request) updateTableName(table string) *request {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.URL = field.NewString(table, "url")
	r.Mark = field.NewString(table, "mark")
	r.RequestNum = field.NewInt32(table, "request_num")
	r.Content = field.NewString(table, "content")
	r.ParseStatus = field.NewInt32(table, "parse_status")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")

	r.fillFieldMap()

	return r
}

func (r *request) WithContext(ctx context.Context) *requestDo { return r.requestDo.WithContext(ctx) }

func (r request) TableName() string { return r.requestDo.TableName() }

func (r request) Alias() string { return r.requestDo.Alias() }

func (r request) Columns(cols ...field.Expr) gen.Columns { return r.requestDo.Columns(cols...) }

func (r *request) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *request) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 8)
	r.fieldMap["id"] = r.ID
	r.fieldMap["url"] = r.URL
	r.fieldMap["mark"] = r.Mark
	r.fieldMap["request_num"] = r.RequestNum
	r.fieldMap["content"] = r.Content
	r.fieldMap["parse_status"] = r.ParseStatus
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
}

func (r request) clone(db *gorm.DB) request {
	r.requestDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r request) replaceDB(db *gorm.DB) request {
	r.requestDo.ReplaceDB(db)
	return r
}

type requestDo struct{ gen.DO }

func (r requestDo) Debug() *requestDo {
	return r.withDO(r.DO.Debug())
}

func (r requestDo) WithContext(ctx context.Context) *requestDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r requestDo) ReadDB() *requestDo {
	return r.Clauses(dbresolver.Read)
}

func (r requestDo) WriteDB() *requestDo {
	return r.Clauses(dbresolver.Write)
}

func (r requestDo) Session(config *gorm.Session) *requestDo {
	return r.withDO(r.DO.Session(config))
}

func (r requestDo) Clauses(conds ...clause.Expression) *requestDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r requestDo) Returning(value interface{}, columns ...string) *requestDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r requestDo) Not(conds ...gen.Condition) *requestDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r requestDo) Or(conds ...gen.Condition) *requestDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r requestDo) Select(conds ...field.Expr) *requestDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r requestDo) Where(conds ...gen.Condition) *requestDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r requestDo) Order(conds ...field.Expr) *requestDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r requestDo) Distinct(cols ...field.Expr) *requestDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r requestDo) Omit(cols ...field.Expr) *requestDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r requestDo) Join(table schema.Tabler, on ...field.Expr) *requestDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r requestDo) LeftJoin(table schema.Tabler, on ...field.Expr) *requestDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r requestDo) RightJoin(table schema.Tabler, on ...field.Expr) *requestDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r requestDo) Group(cols ...field.Expr) *requestDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r requestDo) Having(conds ...gen.Condition) *requestDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r requestDo) Limit(limit int) *requestDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r requestDo) Offset(offset int) *requestDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r requestDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *requestDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r requestDo) Unscoped() *requestDo {
	return r.withDO(r.DO.Unscoped())
}

func (r requestDo) Create(values ...*model.Request) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r requestDo) CreateInBatches(values []*model.Request, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r requestDo) Save(values ...*model.Request) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r requestDo) First() (*model.Request, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Request), nil
	}
}

func (r requestDo) Take() (*model.Request, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Request), nil
	}
}

func (r requestDo) Last() (*model.Request, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Request), nil
	}
}

func (r requestDo) Find() ([]*model.Request, error) {
	result, err := r.DO.Find()
	return result.([]*model.Request), err
}

func (r requestDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Request, err error) {
	buf := make([]*model.Request, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r requestDo) FindInBatches(result *[]*model.Request, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r requestDo) Attrs(attrs ...field.AssignExpr) *requestDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r requestDo) Assign(attrs ...field.AssignExpr) *requestDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r requestDo) Joins(fields ...field.RelationField) *requestDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r requestDo) Preload(fields ...field.RelationField) *requestDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r requestDo) FirstOrInit() (*model.Request, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Request), nil
	}
}

func (r requestDo) FirstOrCreate() (*model.Request, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Request), nil
	}
}

func (r requestDo) FindByPage(offset int, limit int) (result []*model.Request, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r requestDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r requestDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r requestDo) Delete(models ...*model.Request) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *requestDo) withDO(do gen.Dao) *requestDo {
	r.DO = *do.(*gen.DO)
	return r
}
