package mapper

import (
	"github.com/stretchr/testify/assert"
	"rentmanagement/pkg/types"
	"testing"
	"time"
)

func TestAddHouseInfo(t *testing.T) {
	house := &types.HouseInfo{
		Name:      "lll",
		Phone:     180123123123,
		Address:   "四川",
		FreeRoom:  3,
		TotalRoom: 10,
	}

	house.CreateTime = time.Now()
	house.UpdateTime = time.Now()

	err := AddHouseInfo(house)
	assert.Nil(t, err)
}
