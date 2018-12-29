package mapper

import (
	"gopkg.in/logger.v1"

	// "github.com/go-xorm/builder"
	"github.com/satori/go.uuid"
	"github.com/xorm-page/page"

	"rentmanagement/pkg/config"
	"rentmanagement/pkg/types"
)

//WEGUsage .
type WEGUsage struct {
}

//DefaultWegUsage .
var DefaultWegUsage = new(WEGUsage)

func (w *WEGUsage) add(weg *types.WegDosage) error {
	if err := types.IsWegDosageValid(*weg); err != nil {
		log.Error(err)
		return err
	}
	price, err := config.ReadPriceConfig()
	if err != nil {
		log.Error(err)
		return err
	}

	weg.Wprice = price.Price.Water * weg.WaterDosage
	weg.Eprice = price.Price.Electric * weg.ElectricDosage
	weg.Gprice = price.Price.Gas * weg.GasDosage

	weg.Total = weg.Wprice + weg.Eprice + weg.Gprice

	weg.UUID = uuid.NewV4().String()

	b, err := GetEngine().Insert(weg)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debug("结果: ", b)
	return nil
}

func (w *WEGUsage) find(roomName string, pa *page.Pageable) (*page.Page, error) {
	var wegs []types.WegDosage

	if pa.PageIndex == 0 {
		pa.PageIndex = 1
	}
	if pa.PageSize == 0 {
		pa.PageSize = 10
	}
	session := GetEngine().Alias("w").Where("room_name=?", roomName)
	return page.NewBuilder().Page(pa).Session(session).Data(&wegs).Build()
}

func (w *WEGUsage) findAll(pa *page.Pageable) (*page.Page, error) {
	var wegs []types.WegDosage

	if pa.PageIndex == 0 {
		pa.PageIndex = 1
	}
	if pa.PageSize == 0 {
		pa.PageSize = 10
	}
	session := GetEngine().Alias("w").Where("1=1")
	return page.NewBuilder().Page(pa).Session(session).Data(&wegs).Build()
}

//AddWegDosage .
func AddWegDosage(weg *types.WegDosage) error {
	return DefaultWegUsage.add(weg)
}

//FindWegDosage .
func FindWegDosage(roomName string, pa *page.Pageable) (*page.Page, error) {
	return DefaultWegUsage.find(roomName, pa)
}

//FindAllWegDosage .
func FindAllWegDosage(pa *page.Pageable) (*page.Page, error) {
	return DefaultWegUsage.findAll(pa)
}
