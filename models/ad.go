package models

import "strconv"

// Ad table ads
type Ad struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

// ListAllAd 广告列表
func ListAllAd() ([]*Ad, error) {
	var ads []*Ad
	err := DB.Find(&ads).Error
	return ads, err
}

// GetAdById 广告详情
func GetAdById(id string) (*Ad, error) {
	pid, err := strconv.ParseUint(id, 10, 64)
	ad := &Ad{}
	err = DB.First(&ad, pid).Error
	return ad, err
}

// CountAd 广告总数 TODO delete
func CountAd() int {
	var count int
	DB.Model(&Ad{}).Count(&count)
	return count
}
