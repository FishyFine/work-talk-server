package mysqlFunc

type MysqlConf struct{
	Username string
	Password string
	Host string
	Port string
	Database string
}

func NewMysqlConf() *MysqlConf{
	conf := new(MysqlConf)
	conf.Init()
	return conf
}

func (m *MysqlConf)Init(){
	m.Username = "root"
	m.Password = "123456"
	m.Host = "47.99.166.254"
	m.Port = "13306" //dev:This also need to modify later.
	m.Database = "work_talk"
}
