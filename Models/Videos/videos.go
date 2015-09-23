package Videos

type Videolist struct {
	VideoList []Video `json:"VideoList"`
}

type Video struct {
	VideoId      int    `json:"videoId"`
	VideoName    string `json:"videoName"`
	FileName     string `json:"fileName"`
	IsProccessed bool   `json:"isProccessed"`
}
