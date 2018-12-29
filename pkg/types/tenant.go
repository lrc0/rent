package types

import (
	"fmt"
	"time"
)

//Tenant .
type Tenant struct {
	UUID        string    `xorm:"char(36) not null 'uuid'" json:"uuid"`
	RoomName    string    `xorm:"'room_name'" json:"room_name"`       //房间名
	Name        string    `xorm:"'name'" json:"name"`                 //租客姓名
	IDNumber    string    `xorm:"'id_number'" json:"id_number"`       //身份证号码
	PhoneNumber string    `xorm:"'phone_number'" json:"phone_number"` //电话号码
	MonthlyRent int       `xorm:"'monthly_rent'" json:"monthly_rent"` //月租金
	LeasePeriod int       `xorm:"'lease_period'" json:"lease_period"` //租期
	DateFrom    string    `xorm:"'date_from'" json:"date_from"`       //起始时间
	DateTo      string    `xorm:"'date_to'" json:"date_to"`           //结束时间
	TotalRent   int       `xorm:"'total_rent'" json:"total_rent"`     //总租金
	CreateTime  time.Time `xorm:"created 'create_time'" json:"create_time,omitempty"`
	UpdateTime  time.Time `xorm:"updated 'update_time'" json:"update_time,omitempty"`
}

//IsTenantValid .
func IsTenantValid(t Tenant) error {
	if t.RoomName == "" {
		return fmt.Errorf("房间名不能为空")
	}
	if t.Name == "" {
		return fmt.Errorf("租客姓名不能为空")
	}
	if t.IDNumber == "" {
		return fmt.Errorf("租客身份证号码不能为空")
	}
	if t.PhoneNumber == "" {
		return fmt.Errorf("租客电话号码不能为空")
	}
	if t.MonthlyRent <= 0 {
		return fmt.Errorf("月租金有误")
	}
	if t.LeasePeriod <= 0 {
		return fmt.Errorf("租期有误")
	}

	return nil
}
