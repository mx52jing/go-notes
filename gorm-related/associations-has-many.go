package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Wallet struct {
	ID      uint
	Amount  float64
	HMDogID uint // 外键，就是HMDog.ID
}

type HMDog struct {
	ID          uint
	Name        string
	HMGirlGodID uint // 外键，就是HMGirlGod.ID
	Wallet      Wallet
}

type HMGirlGod struct {
	ID     uint
	Name   string
	HMDogs []HMDog
}

func migrateTable() {
	err := DB.AutoMigrate(&HMGirlGod{}, &HMDog{}, &Wallet{})
	if err != nil {
		return
	}
}

func createDGData() {
	hmDog1 := HMDog{Name: "狗哥1号"}
	hmDog2 := HMDog{Name: "狗哥2号"}
	hmDog3 := HMDog{Name: "狗哥3号"}
	hmDog4 := HMDog{Name: "狗哥4号"}
	hmDog5 := HMDog{Name: "狗哥5号"}
	hmDog6 := HMDog{Name: "狗哥6号"}
	hmGirlGod1 := HMGirlGod{
		Name:   "女神1号",
		HMDogs: []HMDog{hmDog1, hmDog2},
	}
	hmGirlGod2 := HMGirlGod{
		Name:   "女神2号",
		HMDogs: []HMDog{hmDog3, hmDog4},
	}
	hmGirlGod3 := HMGirlGod{
		Name:   "女神3号",
		HMDogs: []HMDog{hmDog5, hmDog6},
	}
	hmGirlGods := []HMGirlGod{hmGirlGod1, hmGirlGod2, hmGirlGod3}
	DB.Model(&HMGirlGod{}).Create(&hmGirlGods)
}
func clearAssociation() {
	// 找出女神1号
	var g1 HMGirlGod
	DB.Model(&HMGirlGod{}).Find(&g1, 1)
	// 清除其关联的所有狗哥
	err := DB.Model(&g1).Association("HMDogs").Clear()
	if err != nil {
		return
	}
}

func appendAssociation() {
	// 找出女神1号
	var g1 HMGirlGod
	DB.Model(&HMGirlGod{}).Find(&g1, 1)
	// 找出狗哥1号和狗哥2号
	var d1 HMDog
	var d2 HMDog
	DB.Model(&HMDog{}).Find(&d1, 1)
	DB.Model(&HMDog{}).Find(&d2, 2)
	// 将狗哥1号和女神1号关联
	//DB.Model(&g1).Association("HMDogs").Append(&d1)
	// 将狗哥1号、狗哥2号和女神1号关联
	ds := []HMDog{d1, d2}
	DB.Model(&g1).Association("HMDogs").Append(&ds)
}

// 添加新成员
func appendNewRowAssociation() {
	// 找出女神1号
	var g1 HMGirlGod
	DB.Model(&HMGirlGod{}).Find(&g1, 1)
	d7 := HMDog{Name: "狗哥7号"}
	d8 := HMDog{Name: "狗哥8号"}
	dSlice := []HMDog{d7, d8}
	DB.Model(&g1).Association("HMDogs").Append(&dSlice)
}

func replaceAssociation() {
	// 新创建狗哥9号
	d9 := HMDog{Name: "狗哥9号"}
	// 找出女神1号
	var g1 HMGirlGod
	DB.Model(&HMGirlGod{}).Find(&g1, 1)
	DB.Model(&g1).Association("HMDogs").Replace(&d9)
}

func replaceMulti() {
	// 找出狗哥3号
	var d3 HMDog
	DB.Model(&HMDog{}).Find(&d3, 3)
	// 找出狗哥4号
	var d4 HMDog
	DB.Model(&HMDog{}).Find(&d4, 4)
	// 找出狗哥9号
	var d9 HMDog
	DB.Model(&HMDog{}).Find(&d9, 9)
	// 将女神2号关联的狗哥替换为3号和9号
	// 找出女神2号
	var g2 HMGirlGod
	DB.Model(&HMGirlGod{}).Find(&g2, 2)
	g2Dogs := []HMDog{d3, d4, d9}
	//err := DB.Model(&g2).Association("HMDogs").Replace(&d3, &d4, &d9)
	err := DB.Model(&g2).Association("HMDogs").Replace(&g2Dogs)
	if err != nil {
		return
	}
}

func queryData() {
	// 查询女神2号的所有狗哥
	//var g2 HMGirlGod
	//DB.Preload("HMDogs").Find(&g2, 2)
	//fmt.Printf("g2 is %v\n", g2)
	// 查询女神2号下ID为3的狗哥
	var g2 HMGirlGod
	DB.Preload("HMDogs", func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", 3)
	}).Find(&g2, 2)
	fmt.Printf("g2 is %v\n", g2)
}

func queryByPreloadCondition() {
	// 查询女神2号下ID为3的狗哥
	var g2 HMGirlGod
	DB.Preload("HMDogs", "id = ?", 3).Find(&g2, 2)
	// g2 is {2 女神2号 [{3 狗哥3号 2 {0 0}}]}
	fmt.Printf("g2 is %v\n", g2)
}

func linkQuery() {
	// 找出女神1号下面的所有关联的狗哥
	var g1 HMGirlGod
	//DB.Preload("HMDogs").Find(&g1, 1)
	// 使用链式查询
	DB.Preload("HMDogs.Wallet").Find(&g1, 1)
	fmt.Printf("g1 is %v\n", g1)
}

func multiLinkQuery() {
	// 找出女神1号下面的所有关联的狗哥
	var g1 HMGirlGod
	//DB.Preload("HMDogs").Find(&g1, 1)
	// 使用链式查询
	DB.Preload("HMDogs.Wallet").Preload("HMDogs", "id > ?", 3).Find(&g1, 1)
	fmt.Printf("g1 is %v\n", g1)
}

func joinsQuery() {
	// 找出女神1号下面的所有关联的狗哥
	var g1 HMGirlGod
	// 查询女神1号关联的狗哥中  金额大于10000的
	DB.Preload("HMDogs", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Wallet").Where("amount > ?", 10000)
	}).Find(&g1, 1)
	// 查询女神1号关联的狗哥中  金额大于10000并且id > 3的
	//DB.Preload("HMDogs", func(db *gorm.DB) *gorm.DB {
	//	return db.Joins("Wallet").Where("amount > ? AND hm_dogs.id > ?", 10000, 3)
	//}).Find(&g1, 1)
	fmt.Printf("g1 is %v\n", g1)
}

func main() {
	//migrateTable()
	//createDGData()
	//clearAssociation()
	//appendAssociation()
	//appendNewRowAssociation()
	//replaceAssociation()
	//replaceMulti()
	//queryData()
	//queryByPreloadCondition()
	//linkQuery()
	//multiLinkQuery()
	joinsQuery()
}
