package config

import (
	"io/ioutil"

	"gopkg.in/logger.v1"
	"gopkg.in/yaml.v2"

	"rentmanagement/pkg/types"
	"rentmanagement/pkg/util"
)

//Filepath .
var Filepath = "/src/rentmanagement/conf/config.yml"

//ReadDBConfig .
func ReadDBConfig() (*types.Mysql, error) {
	var mysql types.Mysql
	file, err := ioutil.ReadFile(util.Home() + Filepath)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = yaml.Unmarshal(file, &mysql)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &mysql, nil
}

//ReadPriceConfig .
func ReadPriceConfig() (types.Prices, error) {
	var price types.Prices
	file, err := ioutil.ReadFile(util.Home() + Filepath)
	if err != nil {
		log.Error(err)
		return types.Prices{}, err
	}

	err = yaml.Unmarshal(file, &price)
	if err != nil {
		log.Error(err)
		return types.Prices{}, err
	}
	return price, nil
}
