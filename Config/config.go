package config

var ObjConfig *config
var ObjDBParams *DBParams

type config struct {
	ConnectionString string
}

type DBParams struct {
	domain   string
	username string
	password string
	port     string
	database string
}

func NewDBParams() *DBParams {
	if ObjDBParams == nil {
		ObjDBParams = &DBParams{"172.16.1.5", "imedia", "imedia2006", "1433", "GoLearn"}
	}
	return ObjDBParams
}

func NewConfig() *config {
	dbParams := NewDBParams()
	if ObjConfig == nil {
		ObjConfig = &config{"server=" + dbParams.domain + ";user id=" + dbParams.username + ";password=" + dbParams.password + ";port=" + dbParams.port + ";database=" + dbParams.database}
	}
	return ObjConfig
}
