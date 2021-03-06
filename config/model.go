package config

type Conf struct {
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
	Mongo Mongo `json:"mongo"`
}

type (
	Mysql struct {
		Dbname   string `json:"dbname"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
	}
	Redis struct {
		Addr     string `json:"addr"`
		DB       int    `json:"db"`
		Password string `json:"password"`
	}

	Mongo struct {
		Uri string `json:"uri"`
		Db  string `json:"db"`
	}
)
