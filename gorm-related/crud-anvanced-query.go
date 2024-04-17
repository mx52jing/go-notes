package main

import (
	"fmt"
	"gorm.io/gorm"
)

type SmartFields struct {
	Name   string
	Age    int
	Gender int
}

func smartSelectField() {
	var smartT []SmartFields
	DB.Model(&Teacher{}).Where("age >= ?", "22").Find(&smartT)
	fmt.Println("smartT is=>", smartT)
}

func subQuery() {
	var subQueryT []SmartFields
	//DB.Where("age >= (?)", DB.Model(&Teacher{}).Select("AVG(age)")).Find(&subQueryT)
	DB.Table("(?) as t", DB.Model(&Teacher{}).Select("name", "age")).Where("age > ?", 18).Find(&subQueryT)
	fmt.Println("subQueryT is=>", subQueryT)
}

func namedQuery() {
	var namedQueryT []Teacher
	//DB.Where("name1 = @name OR name2 = @name", sql.Named("name", "如燕")).Find(&namedQueryT)
	DB.Where("name1 = @name OR name2 = @name", map[string]interface{}{"name": "如燕"}).Find(&namedQueryT)
	fmt.Println("namedQueryT is=>", namedQueryT)
}

func scanMap() {
	//var res map[string]interface{}
	//DB.Model(&Teacher{}).Where("age >= ?", 18).Take(&res)
	var res []map[string]interface{}
	DB.Table("teachers").Where("age >= ?", 18).Find(&res)
	fmt.Println("res is=>", res)
}

func ageGreaterThan18(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 18)
}

func nameLikeMul(db *gorm.DB) *gorm.DB {
	return db.Where("name LIKE ?", "张%")
}

func scopeExample() {
	var scopeT []Teacher
	//DB.Scopes(ageGreaterThan18).Find(&scopeT)
	DB.Scopes(ageGreaterThan18, nameLikeMul).Find(&scopeT)
	fmt.Println("scopeT =>", scopeT)
}

func nums() {
	var count int64
	//var res []Teacher
	DB.Model(&Teacher{})
	DB.Model(&Teacher{}).Where("age >= ?", 18).Count(&count)
	fmt.Println("count =>", count)
}

func main() {
	//smartSelectField()
	//subQuery()
	//namedQuery()
	//scanMap()
	//scopeExample()
	nums()
}
