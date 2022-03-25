package nsqFunc

import "github.com/nsqio/go-nsq"

type NSQConf struct{
	NSQServerAddress string
	Topic string
	config *nsq.Config
}

func NewNSQConf() *NSQConf {
	conf := new(NSQConf)
	conf.Init()
	return conf
}

func (n *NSQConf)Init(){
	n.NSQServerAddress = "47.99.166.254:4150"
	n.Topic = "Register"
	n.config = nsq.NewConfig()
}

