package Videos

type Videolist struct {
	VideoList []Video
}

type Video struct {
	Id        int
	Name      string
	Videopath string
}
