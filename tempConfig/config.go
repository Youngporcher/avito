package tempConfig

type Config struct {
	Login    string
	Password string
	Ip       string
	Port     string
	Db_name  string
}

var DB_Config = Config{
	Login:    "postgres",
	Password: "12345",
	Ip:       "localhost",
	Port:     "5432",
	Db_name:  "segmentService",
}
