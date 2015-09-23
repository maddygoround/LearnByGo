package Videoreader

import (
	"../../Config"
	"../../Models/Videos"
	"database/sql"
	"encoding/json"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func GetAllVideos() ([]byte, error) {

	cObj := config.NewConfig()
	db, err := sql.Open("mssql", cObj.ConnectionString)

	if err != nil {
		log.Fatal(err)
	}

	errPing := db.Ping()

	if errPing != nil {
		log.Fatal("Ping ", errPing)
	}

	defer db.Close()

	rows, err := db.Query("select * from Videos")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	vidObj := Videos.Video{}

	videolistObj := new(Videos.Videolist)

	for rows.Next() {

		err := rows.Scan(
			&vidObj.Id,
			&vidObj.Name,
			&vidObj.Videopath,
		)

		videolistObj.VideoList = append(videolistObj.VideoList, vidObj)

		if err != nil {
			log.Fatal(err)
		}
	}

	return json.Marshal(videolistObj)

}
