package nsqFunc

import "github.com/nsqio/go-nsq"

type NSQHelper struct {
	Producer *nsq.Producer
}

func NewNSQHelper(conf *NSQConf) (*NSQHelper, error) {
	helper := new(NSQHelper)
	var err error
	helper.Producer, err = nsq.NewProducer(conf.NSQServerAddress, conf.config)
	if err != nil {
		return nil, err
	}
	return helper, nil
}
