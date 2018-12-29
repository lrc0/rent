package mapper

import (
	"github.com/stretchr/testify/assert"
	"github.com/xorm-page/page"
	"rentmanagement/pkg/types"
	"testing"
	"time"
)

func TestAddWEGDosage(t *testing.T) {
	weg := &types.WegDosage{
		RoomName:       "二楼前",
		Water:          100,
		WaterDosage:    10,
		Electric:       100,
		ElectricDosage: 10,
		Gas:            100,
		GasDosage:      10,
	}
	weg.CreateTime = time.Now()
	weg.UpdateTime = time.Now()

	err := AddWegDosage(weg)
	assert.Nil(t, err)
}

func TestFindWEGDosage(t *testing.T) {
	roomName := "二楼前"
	pa := &page.Pageable{
		PageIndex: 1,
		PageSize:  10,
	}

	_, err := FindWegDosage(roomName, pa)
	assert.Nil(t, err)
}
