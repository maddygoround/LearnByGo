package Videoreader

import (
	"../../dbconnection"
	"../../models/videos"
	"log"
)

func GetAllVideos() *Videos.Videolist {
	rows, err := SQLCtx.SQLdb.Query("select id,videosize,name,filename from go_videos")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	vidObj := Videos.Video{}

	videolistObj := new(Videos.Videolist)

	for rows.Next() {

		err := rows.Scan(
			&vidObj.VideoId,
			&vidObj.VideoSize,
			&vidObj.VideoName,
			&vidObj.FileName,
		)

		videolistObj.VideoList = append(videolistObj.VideoList, vidObj)

		if err != nil {
			log.Fatal(err)
		}
	}

	return videolistObj

}
