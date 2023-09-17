// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package zet_query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/anguloc/zet/internal/dao/zet_model"
)

func newRss(db *gorm.DB, opts ...gen.DOOption) rss {
	_rss := rss{}

	_rss.rssDo.UseDB(db, opts...)
	_rss.rssDo.UseModel(&zet_model.Rss{})

	tableName := _rss.rssDo.TableName()
	_rss.ALL = field.NewAsterisk(tableName)
	_rss.ID = field.NewUint64(tableName, "id")
	_rss.Mark = field.NewString(tableName, "mark")
	_rss.Title = field.NewString(tableName, "title")
	_rss.Link = field.NewString(tableName, "link")
	_rss.Author = field.NewString(tableName, "author")
	_rss.PubDate = field.NewTime(tableName, "pub_date")
	_rss.Desctiption = field.NewString(tableName, "desctiption")
	_rss.Bittorrent = field.NewString(tableName, "bittorrent")
	_rss.GUID = field.NewString(tableName, "guid")
	_rss.Uk = field.NewString(tableName, "uk")
	_rss.CreatedAt = field.NewTime(tableName, "created_at")
	_rss.UpdatedAt = field.NewTime(tableName, "updated_at")

	_rss.fillFieldMap()

	return _rss
}

type rss struct {
	rssDo rssDo

	ALL         field.Asterisk
	ID          field.Uint64
	Mark        field.String // 标识
	Title       field.String // 标题
	Link        field.String // 链接
	Author      field.String // 作者
	PubDate     field.Time   // 发布时间
	Desctiption field.String // 描述
	Bittorrent  field.String // 种子路径
	GUID        field.String // 资源标识
	Uk          field.String // 资源标识，业务定义
	CreatedAt   field.Time   // 插入时间
	UpdatedAt   field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (r rss) Table(newTableName string) *rss {
	r.rssDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r rss) As(alias string) *rss {
	r.rssDo.DO = *(r.rssDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *rss) updateTableName(table string) *rss {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewUint64(table, "id")
	r.Mark = field.NewString(table, "mark")
	r.Title = field.NewString(table, "title")
	r.Link = field.NewString(table, "link")
	r.Author = field.NewString(table, "author")
	r.PubDate = field.NewTime(table, "pub_date")
	r.Desctiption = field.NewString(table, "desctiption")
	r.Bittorrent = field.NewString(table, "bittorrent")
	r.GUID = field.NewString(table, "guid")
	r.Uk = field.NewString(table, "uk")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")

	r.fillFieldMap()

	return r
}

func (r *rss) WithContext(ctx context.Context) *rssDo { return r.rssDo.WithContext(ctx) }

func (r rss) TableName() string { return r.rssDo.TableName() }

func (r rss) Alias() string { return r.rssDo.Alias() }

func (r rss) Columns(cols ...field.Expr) gen.Columns { return r.rssDo.Columns(cols...) }

func (r *rss) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *rss) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 12)
	r.fieldMap["id"] = r.ID
	r.fieldMap["mark"] = r.Mark
	r.fieldMap["title"] = r.Title
	r.fieldMap["link"] = r.Link
	r.fieldMap["author"] = r.Author
	r.fieldMap["pub_date"] = r.PubDate
	r.fieldMap["desctiption"] = r.Desctiption
	r.fieldMap["bittorrent"] = r.Bittorrent
	r.fieldMap["guid"] = r.GUID
	r.fieldMap["uk"] = r.Uk
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
}

func (r rss) clone(db *gorm.DB) rss {
	r.rssDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r rss) replaceDB(db *gorm.DB) rss {
	r.rssDo.ReplaceDB(db)
	return r
}

type rssDo struct{ gen.DO }

func (r rssDo) Debug() *rssDo {
	return r.withDO(r.DO.Debug())
}

func (r rssDo) WithContext(ctx context.Context) *rssDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r rssDo) ReadDB() *rssDo {
	return r.Clauses(dbresolver.Read)
}

func (r rssDo) WriteDB() *rssDo {
	return r.Clauses(dbresolver.Write)
}

func (r rssDo) Session(config *gorm.Session) *rssDo {
	return r.withDO(r.DO.Session(config))
}

func (r rssDo) Clauses(conds ...clause.Expression) *rssDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r rssDo) Returning(value interface{}, columns ...string) *rssDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r rssDo) Not(conds ...gen.Condition) *rssDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r rssDo) Or(conds ...gen.Condition) *rssDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r rssDo) Select(conds ...field.Expr) *rssDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r rssDo) Where(conds ...gen.Condition) *rssDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r rssDo) Order(conds ...field.Expr) *rssDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r rssDo) Distinct(cols ...field.Expr) *rssDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r rssDo) Omit(cols ...field.Expr) *rssDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r rssDo) Join(table schema.Tabler, on ...field.Expr) *rssDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r rssDo) LeftJoin(table schema.Tabler, on ...field.Expr) *rssDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r rssDo) RightJoin(table schema.Tabler, on ...field.Expr) *rssDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r rssDo) Group(cols ...field.Expr) *rssDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r rssDo) Having(conds ...gen.Condition) *rssDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r rssDo) Limit(limit int) *rssDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r rssDo) Offset(offset int) *rssDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r rssDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *rssDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r rssDo) Unscoped() *rssDo {
	return r.withDO(r.DO.Unscoped())
}

func (r rssDo) Create(values ...*zet_model.Rss) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r rssDo) CreateInBatches(values []*zet_model.Rss, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r rssDo) Save(values ...*zet_model.Rss) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r rssDo) First() (*zet_model.Rss, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*zet_model.Rss), nil
	}
}

func (r rssDo) Take() (*zet_model.Rss, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*zet_model.Rss), nil
	}
}

func (r rssDo) Last() (*zet_model.Rss, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*zet_model.Rss), nil
	}
}

func (r rssDo) Find() ([]*zet_model.Rss, error) {
	result, err := r.DO.Find()
	return result.([]*zet_model.Rss), err
}

func (r rssDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*zet_model.Rss, err error) {
	buf := make([]*zet_model.Rss, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r rssDo) FindInBatches(result *[]*zet_model.Rss, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r rssDo) Attrs(attrs ...field.AssignExpr) *rssDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r rssDo) Assign(attrs ...field.AssignExpr) *rssDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r rssDo) Joins(fields ...field.RelationField) *rssDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r rssDo) Preload(fields ...field.RelationField) *rssDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r rssDo) FirstOrInit() (*zet_model.Rss, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*zet_model.Rss), nil
	}
}

func (r rssDo) FirstOrCreate() (*zet_model.Rss, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*zet_model.Rss), nil
	}
}

func (r rssDo) FindByPage(offset int, limit int) (result []*zet_model.Rss, count int64, err error) {
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

func (r rssDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r rssDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r rssDo) Delete(models ...*zet_model.Rss) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *rssDo) withDO(do gen.Dao) *rssDo {
	r.DO = *do.(*gen.DO)
	return r
}
