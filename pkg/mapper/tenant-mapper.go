package mapper

import (
	"fmt"
	"time"

	"github.com/satori/go.uuid"
	"github.com/xorm-page/page"
	"gopkg.in/logger.v1"

	"rentmanagement/pkg/types"
)

//TenantMapper .
type TenantMapper struct {
}

//DefaultTenantMapper .
var DefaultTenantMapper = new(TenantMapper)

func (r *TenantMapper) add(t *types.Tenant) error {
	if err := types.IsTenantValid(*t); err != nil {
		log.Error(err)
		return err
	}
	t1, err := time.Parse("2006-01-02", t.DateFrom)
	if err != nil {
		return err
	}
	t2, err := time.Parse("2006-01-02", t.DateTo)
	if err != nil {
		return err
	}
	if t1.Before(t2) == false {
		err := fmt.Errorf("起租时间不能晚于到期时间")
		log.Error(err)
		return err
	}

	t.TotalRent = t.MonthlyRent * t.LeasePeriod

	t.UUID = uuid.NewV4().String()

	rs, err := GetEngine().Insert(t)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debug("结果: ", rs)
	return nil
}

func (r *TenantMapper) findAll(pa *page.Pageable) (*page.Page, error) {
	var ts []types.Tenant

	if pa.PageIndex == 0 {
		pa.PageIndex = 1
	}
	if pa.PageSize == 0 {
		pa.PageSize = 10
	}

	session := GetEngine().Alias("w").Where("1=1")
	return page.NewBuilder().Page(pa).Session(session).Data(&ts).Build()
}

func (r *TenantMapper) find(roomName string, pa *page.Pageable) (*page.Page, error) {
	var tenants []types.Tenant

	if pa.PageIndex == 0 {
		pa.PageIndex = 1
	}
	if pa.PageSize == 0 {
		pa.PageSize = 10
	}
	session := GetEngine().Alias("r").Where("room_name=?", roomName)
	return page.NewBuilder().Page(pa).Session(session).Data(&tenants).Build()
}

//AddTenant .
func AddTenant(t *types.Tenant) error {
	return DefaultTenantMapper.add(t)
}

//FindAllTenant .
func FindAllTenant(pa *page.Pageable) (*page.Page, error) {
	return DefaultTenantMapper.findAll(pa)
}

//FindTenant .
func FindTenant(roomName string, pa *page.Pageable) (*page.Page, error) {
	return DefaultTenantMapper.find(roomName, pa)
}
