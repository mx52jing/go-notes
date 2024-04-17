package main

/*
*
has One属于一对一的关系
*/

type HDog struct {
	ID         uint
	Name       string
	HGirlGodID uint // 外键，就是HGirlGod.ID
}

type HGirlGod struct {
	ID   uint
	Name string
	HDog HDog
}

func createData() {
	hGirlGod1 := HGirlGod{Name: "女神1号"}
	DB.Model(&HGirlGod{}).Create(&hGirlGod1)
	hDog1 := HDog{Name: "狗哥1号"}
	hGirlGod2 := HGirlGod{Name: "女神2号", HDog: hDog1}
	DB.Model(&HGirlGod{}).Create(&hGirlGod2)
}

func operateAssociation() {
	var hDog1 HDog
	DB.Model(&HDog{}).Find(&hDog1, 1)
	var hGirlGod1 HGirlGod
	DB.Model(&HGirlGod{}).Find(&hGirlGod1, 1)
	DB.Model(&hGirlGod1).Association("HDog").Append(&hDog1)
	//
	//var hGirlGod2 HGirlGod
	//DB.Model(&HGirlGod{}).Find(&hGirlGod2, 2)
	//DB.Model(&hGirlGod2).Association("HDog").Replace(&hDog1)
}

func createTable() {
	DB.AutoMigrate(&HGirlGod{}, &HDog{})
}

func main() {
	//createTable()
	//createData()
	operateAssociation()
}
