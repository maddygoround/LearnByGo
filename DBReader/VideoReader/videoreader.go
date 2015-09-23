package Videoreader

import (
	"../../Config"
	"../../Models/Videos"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func GetAllVideos() *Videos.Videolist {

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

	rows, err := db.Query("select * from Video")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	vidObj := Videos.Video{}

	videolistObj := new(Videos.Videolist)

	for rows.Next() {

		err := rows.Scan(
			&vidObj.VideoId,
			&vidObj.VideoName,
			&vidObj.FileName,
			&vidObj.IsProccessed,
		)

		videolistObj.VideoList = append(videolistObj.VideoList, vidObj)

		if err != nil {
			log.Fatal(err)
		}
	}

	return videolistObj

}
