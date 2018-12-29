package mapper

import (
	"github.com/stretchr/testify/assert"
	"rentmanagement/pkg/types"
	"testing"
	"time"
)

func TestAddHouseInfo(t *testing.T) {
	house := &types.HouseInfo{
		Name:      "李蕊材",
		Phone:     18081213028,
		Address:   "四川省绵阳市高新区普明后街五组29号",
		FreeRoom:  3,
		TotalRoom: 10,
	}

	house.CreateTime = time.Now()
	house.UpdateTime = time.Now()

	err := AddHouseInfo(house)
	assert.Nil(t, err)
}
