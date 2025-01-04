package repository

import (
	"context"
	"fmt"
	"github.com/xxfasu/urlshortener/internal/repository/gen"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sync"
	"testing"
)

var DB *gorm.DB

func TestMain(m *testing.M) {
	const MySQLDSN = "root:password@tcp(127.0.0.1:3306)/template?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	DB = db
	gen.SetDefault(db)
	code := m.Run()
	fmt.Println("test end")
	os.Exit(code)
}

func TransactionTest(ctx context.Context, fn func(query *gen.Query) error) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var flag error
		var wg sync.WaitGroup
		done := make(chan struct{})
		wg.Add(1)
		go func() {
			defer wg.Done()
			query := gen.Use(tx)
			flag = fn(query)
		}()
		go func() {
			defer close(done)
			wg.Wait()
		}()
		select {
		case <-ctx.Done():
			fmt.Println("Transaction Rollback due to timeout:", ctx.Err())
			return ctx.Err()
		case <-done:
			return flag
		}
	})
}
