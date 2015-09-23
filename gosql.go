package main

import (
	"./DBReader/VideoReader"
	"os"
)

func main() {
	vidObj, _ := Videoreader.GetAllVideos()
	os.Stdout.Write(vidObj)
}
