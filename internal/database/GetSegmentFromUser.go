package database

import (
	"context"
	"github.com/Youngporcher/avito/tempConfig"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"strings"
)

func GetSegmentFromUser(user_id string) []string {
	if user_id == "" {
		println("user id can't be empty")
		os.Exit(1)
	}

	db, err := pgxpool.Connect(context.Background(), "postgres://"+tempConfig.DB_Config.Login+":"+tempConfig.DB_Config.Password+"@"+tempConfig.DB_Config.Ip+":"+tempConfig.DB_Config.Port+"/"+tempConfig.DB_Config.Db_name)
	if err != nil {
		println("error connecting to database")
		os.Exit(1)
	}
	defer db.Close()
	var segments []string
	pgxscan.Select(context.Background(), db, &segments, "SELECT segment FROM users_segments WHERE user_id=$1", strings.ToUpper(user_id))
	return segments
}
