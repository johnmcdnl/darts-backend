package targets

import (
	"github.com/johnmcdnl/darts-backendv2/data"
	"github.com/jinzhu/gorm"
)

type TargetMatch struct {
	Id         uint `gorm:"primary_key" sql:"AUTO_INCREMENT"`
	UserID     int     `sql:"index"`
	TargetName string
	Attempts   int
	Successful int
	gorm.Model
}

type AverageTargetMatch struct {
	TargetName        string
	AveragePercentage float64
	TotalAttempts     int
	TotalSuccess      int
}

func Migrate() {
	data.GetDB().AutoMigrate(&TargetMatch{})
}

func createNewMatch(tm TargetMatch) {
	data.GetDB().Save(tm)
}

func allMatches() []TargetMatch {
	var m []TargetMatch
	data.GetDB().Find(&m)
	return m
}

func allMatchesForUser(userId int) []TargetMatch {
	var m []TargetMatch
	data.GetDB().Where(&TargetMatch{UserID: userId}).Find(&m)
	return m
}

func allMatchesByTargetName(targetName string) []TargetMatch {
	var m []TargetMatch
	data.GetDB().Where(&TargetMatch{TargetName: targetName}).Find(&m)
	return m
}

func allMatchesByUserAndTarget(userId int, targetName string) []TargetMatch {
	var m []TargetMatch
	data.GetDB().Where(&TargetMatch{UserID: userId, TargetName: targetName}).Find(&m)
	return m
}

func allMatchesByUserAndTargetAverage(userId int, targetName string) AverageTargetMatch {

	var avg = AverageTargetMatch{TargetName:targetName}

	matches := allMatchesByUserAndTarget(userId, targetName)

	for _, match := range matches {
		avg.TotalAttempts += match.Attempts
		avg.TotalSuccess += match.Successful
	}
	avg.AveragePercentage = float64(avg.TotalSuccess) / float64(avg.TotalAttempts)
	return avg
}

func allMatchesByAreaType(at string) []TargetMatch {
	var m []TargetMatch
	//TODO
	return m
}

func allMatchesByUserAndAreaType(userId int, areaType string) []TargetMatch {
	var m []TargetMatch
	//TODO
	return m
}