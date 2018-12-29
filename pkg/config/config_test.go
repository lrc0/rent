package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadDBConfig(t *testing.T) {
	mysql, err := ReadDBConfig()
	fmt.Printf("=====>>>> %+v\n", mysql)

	assert.Nil(t, err)
}

func TestReadPriceConfig(t *testing.T) {
	price, err := ReadPriceConfig()
	fmt.Printf("=======>>> %+v \n", price)
	assert.Nil(t, err)
}
