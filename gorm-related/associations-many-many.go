package main

import (
	"fmt"
)

type M2MWallet struct {
	ID       uint
	Amount   float64
	M2MDogID uint // 外键，就是M2MDog.ID
}

type M2MDog struct {
	ID          uint
	Name        string
	M2MGirlGods []M2MGirlGod `gorm:"many2many:dog_girl_god"` // 一定要声明这个tag, 一个狗哥可以关联多个女神
	M2MWallet   M2MWallet
}

type M2MGirlGod struct {
	ID      uint
	Name    string
	M2MDogs []M2MDog `gorm:"many2many:dog_girl_god"` // 一定要声明这个tag, 一个女神可以关联多个狗哥
}

func autoMigrate() {
	DB.AutoMigrate(&M2MDog{}, &M2MGirlGod{}, &M2MWallet{})
}

func createM2MData() {
	d1 := M2MDog{Name: "狗哥1号"}
	d2 := M2MDog{Name: "狗哥2号"}
	d3 := M2MDog{Name: "狗哥3号"}
	d4 := M2MDog{Name: "狗哥4号"}
	d5 := M2MDog{Name: "狗哥5号"}
	d6 := M2MDog{Name: "狗哥6号"}
	dogs := []M2MDog{d1, d2, d3, d4, d5, d6}
	DB.Model(&M2MDog{}).Create(&dogs)
}

func girl2DogAssociation() {
	// 狗哥1号
	var d1 M2MDog
	DB.Model(&M2MDog{}).Find(&d1, 1)
	// 创建女神1号 并将其和狗哥1号关联
	//g1 := M2MGirlGod{Name: "女神1号"}
	//DB.Model(&d1).Association("M2MGirlGods").Append(&g1)
	// 创建女神2号 并将其和狗哥1号关联
	//g2 := M2MGirlGod{Name: "女神2号"}
	//DB.Model(&d1).Association("M2MGirlGods").Append(&g2)
	// 创建女神3、4号 并将其和狗哥1号关联
	g3 := M2MGirlGod{Name: "女神3号"}
	g4 := M2MGirlGod{Name: "女神4号"}
	DB.Model(&d1).Association("M2MGirlGods").Append(&g3, &g4)
}

func queryGirl() {
	// 查出女神1号
	var g1 M2MGirlGod
	DB.Model(&M2MGirlGod{}).Find(&g1, 1)
	var dogs []M2MDog
	// 查询女神1号关联的狗哥
	//DB.Model(&g1).Preload("M2MWallet").Association("M2MDogs").Find(&dogs)
	DB.Model(&g1).Preload("M2MWallet").Preload("M2MGirlGods").Association("M2MDogs").Find(&dogs)
	/**
	女神1号关联的狗哥 => [{1 狗哥1号 [{1 女神1号 []} {2 女神2号 []} {3 女神3号 []} {4 女神4号 []}] {1 10000 1}} {2 狗哥2号 [{1 女神1号 []}] {2 80000 2}} {3 狗哥3号 [{1 女神1号 []}]
	*/
	fmt.Printf("女神1号关联的狗哥 => %v\n", dogs)
}

func queryByJoin() {
	// 查出女神1号
	var g1 M2MGirlGod
	DB.Model(&M2MGirlGod{}).Find(&g1, 1)
	var dogs []M2MDog
	DB.Model(&M2MDog{}).Preload("M2MWallet").Joins("JOIN dog_girl_god dg ON dg.m2_m_dog_id = m2_m_dogs.id").Joins("JOIN m2_m_wallets w on w.m2_m_dog_id = m2_m_dogs.id").Where("dg.m2_m_girl_god_id = ?", g1.ID).Where("w.amount >= ?", 10000).Find(&dogs)
	fmt.Printf("女神1号关联的狗哥 => %v\n", dogs)
}

func main() {
	//autoMigrate()
	//createM2MData()
	//girl2DogAssociation()
	//queryGirl()
	queryByJoin()
}
