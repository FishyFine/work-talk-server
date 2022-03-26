package mysqlFunc

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlHelper struct {
	DB *gorm.DB
}

func NewMysqlHelper(conf *MysqlConf) (*MysqlHelper, error) {
	helper := new(MysqlHelper)
	err := helper.Conn(conf)
	return helper, err
}

func (m *MysqlHelper) Conn(conf *MysqlConf) error {
	args := conf.Username + ":" + conf.Password + "@" + "(" + conf.Host + ":" + conf.Port + ")/" + conf.Database + "?charset=utf8&parseTime=True"
	var err error
	m.DB, err = gorm.Open("mysql", args)
	if err != nil {
		return err
	}
	return nil
}
