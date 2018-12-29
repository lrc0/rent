package mapper

import (
	"github.com/stretchr/testify/assert"
	"rentmanagement/pkg/types"
	"testing"
	"time"
)

func TestAddTenant(t *testing.T) {
	tenant := &types.Tenant{
		RoomName:    "二楼前",
		Name:        "李蕊材",
		IDNumber:    "510703198907173115",
		PhoneNumber: "18081213028",
		MonthlyRent: 260,
		LeasePeriod: 3,
		DateFrom:    "2018-01-02",
		DateTo:      "2018-04-02",
	}
	tenant.CreateTime = time.Now()
	tenant.UpdateTime = time.Now()

	err := AddTenant(tenant)
	assert.Nil(t, err)
}
