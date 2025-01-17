// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/RaymondCode/simple-demo/model/entity"
)

func newFavorite(db *gorm.DB, opts ...gen.DOOption) favorite {
	_favorite := favorite{}

	_favorite.favoriteDo.UseDB(db, opts...)
	_favorite.favoriteDo.UseModel(&entity.Favorite{})

	tableName := _favorite.favoriteDo.TableName()
	_favorite.ALL = field.NewAsterisk(tableName)
	_favorite.FavoriteID = field.NewInt64(tableName, "favorite_id")
	_favorite.UserID = field.NewInt64(tableName, "user_id")
	_favorite.VideoID = field.NewInt64(tableName, "video_id")
	_favorite.CreateTime = field.NewInt64(tableName, "create_time")

	_favorite.fillFieldMap()

	return _favorite
}

type favorite struct {
	favoriteDo favoriteDo

	ALL        field.Asterisk
	FavoriteID field.Int64
	UserID     field.Int64
	VideoID    field.Int64
	CreateTime field.Int64

	fieldMap map[string]field.Expr
}

func (f favorite) Table(newTableName string) *favorite {
	f.favoriteDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f favorite) As(alias string) *favorite {
	f.favoriteDo.DO = *(f.favoriteDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *favorite) updateTableName(table string) *favorite {
	f.ALL = field.NewAsterisk(table)
	f.FavoriteID = field.NewInt64(table, "favorite_id")
	f.UserID = field.NewInt64(table, "user_id")
	f.VideoID = field.NewInt64(table, "video_id")
	f.CreateTime = field.NewInt64(table, "create_time")

	f.fillFieldMap()

	return f
}

func (f *favorite) WithContext(ctx context.Context) *favoriteDo { return f.favoriteDo.WithContext(ctx) }

func (f favorite) TableName() string { return f.favoriteDo.TableName() }

func (f favorite) Alias() string { return f.favoriteDo.Alias() }

func (f favorite) Columns(cols ...field.Expr) gen.Columns { return f.favoriteDo.Columns(cols...) }

func (f *favorite) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *favorite) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 4)
	f.fieldMap["favorite_id"] = f.FavoriteID
	f.fieldMap["user_id"] = f.UserID
	f.fieldMap["video_id"] = f.VideoID
	f.fieldMap["create_time"] = f.CreateTime
}

func (f favorite) clone(db *gorm.DB) favorite {
	f.favoriteDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f favorite) replaceDB(db *gorm.DB) favorite {
	f.favoriteDo.ReplaceDB(db)
	return f
}

type favoriteDo struct{ gen.DO }

func (f favoriteDo) Debug() *favoriteDo {
	return f.withDO(f.DO.Debug())
}

func (f favoriteDo) WithContext(ctx context.Context) *favoriteDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f favoriteDo) ReadDB() *favoriteDo {
	return f.Clauses(dbresolver.Read)
}

func (f favoriteDo) WriteDB() *favoriteDo {
	return f.Clauses(dbresolver.Write)
}

func (f favoriteDo) Session(config *gorm.Session) *favoriteDo {
	return f.withDO(f.DO.Session(config))
}

func (f favoriteDo) Clauses(conds ...clause.Expression) *favoriteDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f favoriteDo) Returning(value interface{}, columns ...string) *favoriteDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f favoriteDo) Not(conds ...gen.Condition) *favoriteDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f favoriteDo) Or(conds ...gen.Condition) *favoriteDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f favoriteDo) Select(conds ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f favoriteDo) Where(conds ...gen.Condition) *favoriteDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f favoriteDo) Order(conds ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f favoriteDo) Distinct(cols ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f favoriteDo) Omit(cols ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f favoriteDo) Join(table schema.Tabler, on ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f favoriteDo) LeftJoin(table schema.Tabler, on ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f favoriteDo) RightJoin(table schema.Tabler, on ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f favoriteDo) Group(cols ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f favoriteDo) Having(conds ...gen.Condition) *favoriteDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f favoriteDo) Limit(limit int) *favoriteDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f favoriteDo) Offset(offset int) *favoriteDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f favoriteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *favoriteDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f favoriteDo) Unscoped() *favoriteDo {
	return f.withDO(f.DO.Unscoped())
}

func (f favoriteDo) Create(values ...*entity.Favorite) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f favoriteDo) CreateInBatches(values []*entity.Favorite, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f favoriteDo) Save(values ...*entity.Favorite) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f favoriteDo) First() (*entity.Favorite, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Favorite), nil
	}
}

func (f favoriteDo) Take() (*entity.Favorite, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Favorite), nil
	}
}

func (f favoriteDo) Last() (*entity.Favorite, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Favorite), nil
	}
}

func (f favoriteDo) Find() ([]*entity.Favorite, error) {
	result, err := f.DO.Find()
	return result.([]*entity.Favorite), err
}

func (f favoriteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Favorite, err error) {
	buf := make([]*entity.Favorite, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f favoriteDo) FindInBatches(result *[]*entity.Favorite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f favoriteDo) Attrs(attrs ...field.AssignExpr) *favoriteDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f favoriteDo) Assign(attrs ...field.AssignExpr) *favoriteDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f favoriteDo) Joins(fields ...field.RelationField) *favoriteDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f favoriteDo) Preload(fields ...field.RelationField) *favoriteDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f favoriteDo) FirstOrInit() (*entity.Favorite, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Favorite), nil
	}
}

func (f favoriteDo) FirstOrCreate() (*entity.Favorite, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Favorite), nil
	}
}

func (f favoriteDo) FindByPage(offset int, limit int) (result []*entity.Favorite, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f favoriteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f favoriteDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f favoriteDo) Delete(models ...*entity.Favorite) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *favoriteDo) withDO(do gen.Dao) *favoriteDo {
	f.DO = *do.(*gen.DO)
	return f
}
