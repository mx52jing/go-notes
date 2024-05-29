package main

import "fmt"

type People struct {
	ID     uint
	Name   string
	Groups []Group `gorm:"many2many:people_groups"`
}

type Group struct {
	ID      uint
	Name    string
	Peoples []People `gorm:"many2many:people_groups"`
}

type PeopleGroup struct {
	PeopleId  uint   `gorm:"primaryKey"`
	GroupId   uint   `gorm:"primaryKey"`
	IsDefault *bool  `gorm:"TINYINT(1);default:0"` // 添加了一个新的column
	People    People `gorm:"foreignKey:PeopleId"`
	Group     Group  `gorm:"foreignKey:GroupId"`
}

func autoMigrateE() {
	DB.AutoMigrate(&People{}, &Group{}, &PeopleGroup{})
}

func createPeople() {
	p1 := People{Name: "人员1号"}
	p2 := People{Name: "人员2号"}
	p3 := People{Name: "人员3号"}
	p4 := People{Name: "人员4号"}
	p5 := People{Name: "人员5号"}
	p6 := People{Name: "人员6号"}
	peoples := []People{p1, p2, p3, p4, p5, p6}
	DB.Model(&People{}).Create(&peoples)
}

func groupPeopleAssociation() {
	// people 1号
	//var p1 People
	//DB.Model(&People{}).Find(&p1, 1)
	//g1 := Group{Name: "群组1"}
	//g2 := Group{Name: "群组2"}
	// 关联用户和群组
	//err := DB.Model(&p1).Association("Groups").Append(&g1, &g2)
	//if err != nil {
	//	return
	//}
	//===============================
	//var p2 People
	//DB.Model(&People{}).Find(&p2, 2)
	//var groups []Group
	//DB.Model(&Group{}).Where("id IN ?", []uint{1, 2}).Find(&groups)
	//DB.Model(&p2).Association("Groups").Append(&groups)
	//===============================
	var p2 People
	DB.Model(&People{}).Find(&p2, 2)
	g3 := Group{Name: "群组3"}
	g4 := Group{Name: "群组4"}
	g5 := Group{Name: "群组5"}
	g6 := Group{Name: "群组6"}
	groups := []Group{g3, g4, g5, g6}
	DB.Model(&p2).Association("Groups").Append(&groups)
}

func changeDefault(peopleId, groupId int) {
	// 先将peopleId关联的其他的群组的is_default设置为false
	DB.Model(&PeopleGroup{}).Where("people_id = ?", peopleId).Update("is_default", false)
	// 再设置传入的peopleId关联的群组的is_default为true
	DB.Model(&PeopleGroup{}).Where("people_id = ? AND group_id = ?", peopleId, groupId).Update("is_default", true)
}

func queryPeopleAssociation(peopleId uint) {
	//var p2 People
	//DB.Model(&People{}).Find(&p2, peopleId)
	var groupsData []Group
	// 查询id为peopleId的row关联的所有群组
	//DB.Model(&p2).Association("Groups").Find(&groupsData)
	// 查询people id为peopleId的row关联的所有群组，并且默认群组排到最前面
	DB.Model(&Group{}).Joins("JOIN people_groups pg ON pg.group_id = groups.id").Where("pg.people_id = ?", peopleId).Order("pg.is_default DESC").Find(&groupsData)
	fmt.Printf("groupsData is %v\n", groupsData)
}

func main() {
	autoMigrateE()
	//createPeople()
	//groupPeopleAssociation()
	//changeDefault(2, 5)
	//queryPeopleAssociation(2)
}
