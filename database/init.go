/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-19 18:16:52
 */

package database

import (
	"time"

	"github.com/glebarez/sqlite"
	_ "github.com/go-sql-driver/mysql" // just for init
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	ConnectRetryTime     = 30          // 重试次数
	ConnectRetryInternal = time.Second // 重试间隔
)

// ConnectDB 连接数据库
func connectionDB(path string) (db *gorm.DB, err error) {
	for i := 0; i <= ConnectRetryTime; i++ {
		db, err = gorm.Open(mysql.New(mysql.Config{
			DriverName: "mysql",
			DSN:        path,
		}), &gorm.Config{
			SkipDefaultTransaction: true,
			CreateBatchSize:        64,
		})

		if err == nil {
			break
		} else {
			log.Errorf("connection to database fail: %v", err)
			time.Sleep(ConnectRetryInternal)
		}
	}
	return db, err
}

func runMigrations(db *gorm.DB, models []interface{}) (err error) {
	if err = db.AutoMigrate(models...); err != nil {
		return err
	}

	for _, m := range models {
		if !db.Migrator().HasTable(m) {
			return errors.Errorf("create table %#v fail", m)
		}
	}

	return nil
}

// InitDB 初始化DB
func InitDB(dbDSN string, models []interface{}) (db *gorm.DB, err error) {
	db, err = connectionDB(dbDSN)
	if err != nil {
		return
	}

	if err = runMigrations(db, models); err != nil {
		return
	}

	return
}

// connectTestDB 连接测试DB，用于单元测试
func connectTestDB(path string) (db *gorm.DB, err error) {
	return gorm.Open(sqlite.Open(path),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
}

// InitTestDB 初始化测试DB，用于单元测试，注意单测按需创建表即可
func InitTestDB(dsn string, models []interface{}) (db *gorm.DB, err error) {
	db, err = connectTestDB(dsn)
	if err != nil {
		return
	}

	if len(models) > 0 {
		if err = runMigrations(db, models); err != nil {
			return
		}
	}
	return
}
