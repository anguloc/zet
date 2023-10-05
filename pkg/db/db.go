package db

import (
	"context"
	"sync"
	"time"

	"github.com/anguloc/zet/internal/app/repo/zetdao/query"
	"github.com/anguloc/zet/internal/errx"
	"github.com/anguloc/zet/pkg/conf"
	"github.com/anguloc/zet/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbMap = sync.Map{}
var initOnce = sync.Once{}

func Init(ctx context.Context) error {
	if conf.Conf().DBZet.Skip() {
		log.Info(ctx, "db_init_skip")
		return nil
	}
	var (
		err   error
		zetDB *gorm.DB
	)
	initOnce.Do(func() {
		if conf.Conf().DBZet.Host == "" {
			return
		}
		zetDB, err = gormDB(ctx, conf.Conf().DBZet)
		if err != nil {
			return
		}
		query.SetDefault(zetDB)
		dbMap.Store("db_zet", zetDB)
	})

	if err != nil {
		log.Error(ctx, "zet_db_init_failed", log.Err(err))
		if conf.Conf().DBZet.Must() {
			return err
		}
	}

	return nil
}

func gormDB(ctx context.Context, c conf.Mysql) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(ZetDSN(c)), gormConfig(ctx))
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(c.MaxLifeTime))
	sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(c.MaxIdleTime))
	return db, nil
}

func gormConfig(ctx context.Context) *gorm.Config {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   log.NewGormLogger(ctx),
	}
}

// ZetDSN 数据库zet的dsn
func ZetDSN(c conf.Mysql) string {
	return c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.DBName + "?" + c.Config
}

func Zet() (*gorm.DB, error) {
	if v, ok := dbMap.Load("db_zet"); ok {
		db := v.(*gorm.DB)
		return db, nil
	}
	return nil, errx.ErrSystem.SetMsg("db未初始化")
}
