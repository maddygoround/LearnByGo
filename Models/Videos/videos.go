package Videos

type Videolist struct {
	VideoList []Video `json:"VideoList"`
}

type Video struct {
	VideoId   int    `json:"videoId"`
	VideoSize int64  `json:"videoSize"`
	VideoName string `json:"videoName"`
	FileName  string `json:"fileName"`
}
