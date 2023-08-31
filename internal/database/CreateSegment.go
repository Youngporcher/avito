package database

import (
	"context"
	"fmt"
	"github.com/Youngporcher/avito/tempConfig"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"strings"
)

func CreateSegment(name string) {
	if name == "" {
		println("segment name can't be empty")
		os.Exit(1)
	}
	db, err := pgxpool.Connect(context.Background(), "postgres://"+tempConfig.DB_Config.Login+":"+tempConfig.DB_Config.Password+"@"+tempConfig.DB_Config.Ip+":"+tempConfig.DB_Config.Port+"/"+tempConfig.DB_Config.Db_name)
	if err != nil {
		println("error connecting to database")
		os.Exit(1)
	}
	defer db.Close()
	_, err = db.Exec(context.Background(), "INSERT INTO segments(segment) VALUES ($1)", strings.ToUpper(name))
	if err != nil {
		fmt.Printf("error create segment: perhaps such a segment \"%s\" already exists\n", name)
	}
}
