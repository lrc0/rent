package mapper

import (
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	// "github.com/spf13/viper"
	"gopkg.in/logger.v1"

	"rentmanagement/pkg/config"
)

var engine *xorm.Engine

func connect() {
	mysql, err := config.ReadDBConfig()
	if err != nil {
		log.Error(err)
		return
	}

	engine, err = xorm.NewEngine(mysql.Database.Type, mysql.Database.URL)
	if err != nil {
		log.Fatal("DB connect error ", err)
	}
	if err := engine.Ping(); err != nil {
		log.Fatal("DB ping error ", err)
	}

	engine.SetMaxIdleConns(mysql.Database.MaxIdle)
	engine.SetMaxOpenConns(mysql.Database.MaxActive)
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "rent_")
	engine.SetTableMapper(tbMapper)
	engine.ShowSQL(mysql.Database.ShowSQL)
	engine.SetLogLevel(core.LOG_ERR)
	engine.SetConnMaxLifetime(time.Second * 30)

}

//GetEngine get db engine
func GetEngine() *xorm.Engine {
	if engine != nil {
		return engine
	}
	connect()
	return engine
}
