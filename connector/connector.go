package connector

import (
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	username = "default"
	password = ",X7h(_JT){bKcL$k"
	database = "liquidity_pool"
	)
func ConnectToClickHouse() (connect *sqlx.DB, err error){
	connect, err = sqlx.Open("clickhouse", "tcp://127.0.0.1:9000?username="+username+"&password=" + password +"&database="+database+"&debug=true")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = connect.Ping()
	if err != nil {
		log.Println(err)
	}

	return connect, nil
}