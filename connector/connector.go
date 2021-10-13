package connector

import (
	//"database/sql"
	//"github.com/ClickHouse/clickhouse-go"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

const (
	username = "default"
	password = ",X7h(_JT){bKcL$k"
	database = "liquidity_pool"
	)
func ConnectToClickHouse() (db *gorm.DB, err error){
	dsn := "tcp://127.0.0.1:9000?database=" + database + "&username=" + username + "&password=" + password + "&read_timeout=10&write_timeout=20"
	db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto Migrate
	//err = db.AutoMigrate(&model.ReservePair{})
	//if err != nil {
	//	return nil, err
	//}
	//err = db.AutoMigrate(&model.TokenPair{})
	//if err != nil {
	//	return nil, err
	//}
	// Set table options
	//err = db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&model.ReservePair{})
	//if err != nil {
	//	return nil, err
	//}

	return db, nil
}