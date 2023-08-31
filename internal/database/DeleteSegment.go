package database

import (
	"context"
	"github.com/Youngporcher/avito/tempConfig"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"strings"
)

func DeleteSegment(name string) {
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

	//Удаление сегмента у пользователей
	_, err = db.Exec(context.Background(), "DELETE FROM users_segments WHERE segment = $1", strings.ToUpper(name))
	if err != nil {
		println("error deleting from user_segments table")
		os.Exit(1)
	}
	//Удаление сегмента у списка сегментов
	_, err = db.Exec(context.Background(), "DELETE FROM segments WHERE segment = $1", strings.ToUpper(name))
	if err != nil {
		println("error deleting from segments table")
		os.Exit(1)
	}
}
