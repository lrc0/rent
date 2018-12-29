package mapper

import (
	"fmt"

	"github.com/satori/go.uuid"
	"github.com/xorm-page/page"
	"gopkg.in/logger.v1"

	"rentmanagement/pkg/types"
)

//HouseInfoMapper .
type HouseInfoMapper struct {
}

//DefaultHouseInfoMapper .
var DefaultHouseInfoMapper = new(HouseInfoMapper)

func (h *HouseInfoMapper) add(house *types.HouseInfo) error {
	err := types.IsHouseInfoValid(house)
	if err != nil {
		log.Error(err)
		return err
	}

	house.UUID = uuid.NewV4().String()

	rs, err := GetEngine().Insert(house)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debug("结果: ", rs)
	return nil
}

func (h *HouseInfoMapper) findAll(pa *page.Pageable) (*page.Page, error) {
	var ts []types.HouseInfo

	if pa.PageIndex == 0 {
		pa.PageIndex = 1
	}
	if pa.PageSize == 0 {
		pa.PageSize = 10
	}

	session := GetEngine().Alias("w").Where("1=1")
	return page.NewBuilder().Page(pa).Session(session).Data(&ts).Build()
}

func (h *HouseInfoMapper) find(name string, pa *page.Pageable) (*page.Page, error) {
	var houseInfo []types.HouseInfo

	if pa.PageIndex == 0 {
		pa.PageIndex = 1
	}
	if pa.PageSize == 0 {
		pa.PageSize = 10
	}
	session := GetEngine().Alias("r").Where("name=?", name)
	return page.NewBuilder().Page(pa).Session(session).Data(&houseInfo).Build()
}

func (h *HouseInfoMapper) delete(uuid string) error {
	if uuid == "" {
		err := fmt.Errorf("uuid 不能为空")
		log.Error(err)
		return err
	}

	rs, err := GetEngine().Where("uuid = ?", uuid).Delete(types.HouseInfo{})
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debug("结果: ", rs)
	return nil
}

//AddHouseInfo .
func AddHouseInfo(house *types.HouseInfo) error {
	return DefaultHouseInfoMapper.add(house)
}

//FindAllHouseInfo .
func FindAllHouseInfo(pa *page.Pageable) (*page.Page, error) {
	return DefaultHouseInfoMapper.findAll(pa)
}

//FindHouseInfo .
func FindHouseInfo(name string, pa *page.Pageable) (*page.Page, error) {
	return DefaultHouseInfoMapper.find(name, pa)
}

//DeleteHouseInfo .
func DeleteHouseInfo(uuid string) error {
	return DefaultHouseInfoMapper.delete(uuid)
}
