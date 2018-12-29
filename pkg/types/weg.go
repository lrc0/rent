package types

import (
	"fmt"
	"time"
)

//WegDosage .
type WegDosage struct {
	UUID           string    `xorm:"char(36) not null 'uuid'" json:"uuid"`
	RoomName       string    `json:"room_name"`
	Water          int       `json:"water"`
	WaterDosage    float32   `json:"water_dosage"`
	Wprice         float32   `xorm:"'w_price'" json:"w_price"`
	Electric       int       `json:"electric"`
	ElectricDosage float32   `json:"electric_dosage"`
	Eprice         float32   `xorm:"'e_price'" json:"e_price"`
	Gas            int       `json:"gas"`
	GasDosage      float32   `json:"gas_dosage"`
	Gprice         float32   `xorm:"'g_price'" json:"g_price"`
	Total          float32   `json:"total_price"`
	CreateTime     time.Time `xorm:"created 'create_time'" json:"create_time,omitempty"`
	UpdateTime     time.Time `xorm:"updated 'update_time'" json:"update_time,omitempty"`
}

//IsWegDosageValid .
func IsWegDosageValid(weg WegDosage) error {
	if weg.Water < 0 {
		return fmt.Errorf("水的底度不能为负")
	}
	if weg.WaterDosage < 0 {
		return fmt.Errorf("水的使用量不能为负")
	}
	if weg.Wprice < 0 {
		return fmt.Errorf("水的价格不能为负")
	}
	if weg.Electric < 0 {
		return fmt.Errorf("电的底度不能为负")
	}
	if weg.ElectricDosage < 0 {
		return fmt.Errorf("电的使用量不能为负")
	}
	if weg.Eprice < 0 {
		return fmt.Errorf("电的价格不能为负")
	}
	if weg.Gas < 0 {
		return fmt.Errorf("气的底度不能为负")
	}
	if weg.GasDosage < 0 {
		return fmt.Errorf("气的使用量不能为负")
	}
	if weg.Gprice < 0 {
		return fmt.Errorf("气的价格不能为负")
	}
	return nil
}
