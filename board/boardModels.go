package board

import "github.com/johnmcdnl/darts-backend/data"

type Area struct {
	AreaName     string
	areaType     AreaType
	AreaTypeName string
}

type AreaType struct {
	AreaTypeName string
}

func getAllAreas() []Area {
	var a []Area
	data.GetDB().Find(&a)
	return a
}

func getAreaByName(name string) Area {
	var a Area
	data.GetDB().Where(Area{AreaName: name}).Find(&a)
	return a
}

func getAllAreaTypes() []AreaType {
	var a []AreaType
	data.GetDB().Find(&a)
	return a
}

func getAreaTypeByName(name string) AreaType {
	var a AreaType
	data.GetDB().Where(AreaType{AreaTypeName: name}).Find(&a)
	return a
}
