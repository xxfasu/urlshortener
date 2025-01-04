package repository

import (
	"context"
	"github.com/xxfasu/urlshortener/internal/conf"
	"github.com/xxfasu/urlshortener/internal/repository/gen"
	"github.com/xxfasu/urlshortener/pkg/logs"
	"github.com/xxfasu/urlshortener/pkg/safe"
	"github.com/xxfasu/urlshortener/pkg/zapgorm2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

type transaction struct {
	DB *gorm.DB
	// rdb    *redis.Client
}

func NewTransaction(
	DB *gorm.DB,
) Transaction {
	return &transaction{
		DB: DB,
	}
}

type Transaction interface {
	Transaction(ctx context.Context, fn func(query *gen.Query) error) error
}

func (r *transaction) Transaction(ctx context.Context, fn func(query *gen.Query) error) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var flag error
		var wg sync.WaitGroup
		done := make(chan struct{})
		wg.Add(1)
		safe.Go(func() {
			defer wg.Done()
			query := gen.Use(tx)
			flag = fn(query)
		})
		safe.Go(func() {
			defer close(done)
			wg.Wait()
		})
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-done:
			return flag
		}
	})
}

func InitDB() (*gorm.DB, func(), error) {
	var (
		db  *gorm.DB
		err error
	)

	logger := zapgorm2.New(logs.Log.Logger)
	// GORM doc: https://gorm.io/docs/connecting_to_the_database.html

	db, err = gorm.Open(mysql.Open(conf.Config.Mysql.Source), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		return nil, nil, err
	}
	db = db.Debug()
	gen.SetDefault(db)
	// Connection Pool config
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, func() {
		sqlDB.Close()
	}, nil
}
