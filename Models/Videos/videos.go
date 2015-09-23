package Videos

type Videolist struct {
	VideoList []Video `json:"VideoList"`
}

type Video struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Videopath string `json:"category"`
}
