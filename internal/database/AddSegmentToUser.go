package database

import (
	"context"
	"fmt"
	"github.com/Youngporcher/avito/tempConfig"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"strings"
)

func AddSegmentToUser(addSegments []string, deleteSegments []string, user_id string) {
	if user_id == "" {
		println("user id can't be empty")
		os.Exit(1)
	}
	if addSegments == nil && deleteSegments == nil {
		println("adding segments and deleting segments can't be empty")
		os.Exit(1)
	}
	db, err := pgxpool.Connect(context.Background(), "postgres://"+tempConfig.DB_Config.Login+":"+tempConfig.DB_Config.Password+"@"+tempConfig.DB_Config.Ip+":"+tempConfig.DB_Config.Port+"/"+tempConfig.DB_Config.Db_name)

	if err != nil {
		println("error connecting to database")
		os.Exit(1)
	}
	defer db.Close()

	// Добавление сегмента пользователю
	if addSegments != nil {
		for _, segment := range addSegments {
			var segmentIsExist int
			//Проверка сегмента, есть ли он в списке сегментов
			err = db.QueryRow(context.Background(), "SELECT Count(id) From segments WHERE segment = $1", strings.ToUpper(segment)).Scan(&segmentIsExist)
			if err != nil {
				println("Request error when adding segments")
				os.Exit(1)
			}
			if segmentIsExist > 0 {
				err = db.QueryRow(context.Background(), "SELECT Count(id) From users_segments WHERE segment = $1 AND user_id = $2", strings.ToUpper(segment), strings.ToUpper(user_id)).Scan(&segmentIsExist)
				if err != nil {
					println("Request error when adding segments")
					os.Exit(1)
				}
				//Проверка есть ли сегмент у пользователя
				if segmentIsExist > 0 {
					fmt.Printf("User \"%s\" already has a \"%s\" segment \n", user_id, segment)
				} else {
					_, err = db.Exec(context.Background(), "INSERT INTO users_segments(user_id,segment) values ($1,$2)", strings.ToUpper(user_id), strings.ToUpper(segment))
					if err != nil {
						println("Request error when adding segments")
						os.Exit(1)
					}
				}

			} else {

				fmt.Printf("error adding segment: \"%s\" segment does not exist\n", segment)
			}
		}
	}
	//Удаление сегментов у пользователя
	if deleteSegments != nil {
		for _, segment := range deleteSegments {
			_, err = db.Exec(context.Background(), "DELETE FROM users_segments WHERE segment = $1 AND user_id = $2", strings.ToUpper(segment), strings.ToUpper(user_id))
			if err != nil {
				println("error deleting from user_segments table")
				os.Exit(1)
			}
		}
	}
}
