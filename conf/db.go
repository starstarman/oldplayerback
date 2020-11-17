package conf

const DriverName = "mysql"

type DbConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	IsRunning bool
}

var DbMasterList = []DbConfig{
	{
		Host:      "10.253.0.106",
		Port:      3306,
		User:      "root",
		Pwd:       "moqikaka3306",
		Database:  "liukuo",
		IsRunning: true,
	},
}

var DbMaster DbConfig = DbMasterList[0]
