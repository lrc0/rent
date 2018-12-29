package types

import (
	"fmt"
	"time"
)

// HouseInfo .
type HouseInfo struct {
	UUID       string    `xorm:"varchar(36) not null 'uuid'" json:"uuid"`
	Name       string    `json:"name" xorm:"name"`
	Phone      int       `json:"phone" xorm:"phone"`
	Address    string    `json:"address" xorm:"address"`
	FreeRoom   int       `json:"free_room" xorm:"free_room"`
	TotalRoom  int       `json:"total_room" xorm:"total_room"`
	CreateTime time.Time `xorm:"created 'create_time'" json:"create_time,omitempty"`
	UpdateTime time.Time `xorm:"updated 'update_time'" json:"update_time,omitempty"`
}

//IsHouseInfoValid .
func IsHouseInfoValid(h *HouseInfo) error {
	switch {
	case h.Name == "":
		return fmt.Errorf("房东姓名不能为空")
	case h.Phone == 0:
		return fmt.Errorf("房东电话有误")
	case h.Address == "":
		return fmt.Errorf("房屋地址不能为空")
	case h.TotalRoom == 0:
		return fmt.Errorf("房屋总数不能为0")
	}
	return nil
}
