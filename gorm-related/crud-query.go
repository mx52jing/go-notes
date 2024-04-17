package main

import (
	"fmt"
)

func queryByFirst() {
	var firstTeacher Teacher
	DB.First(&firstTeacher)
	fmt.Println("firstTeacher =>", firstTeacher)
}

func queryByLast() {
	var lastTeacher Teacher
	DB.Last(&lastTeacher)
	fmt.Println("lastTeacher =>", lastTeacher)
}

func queryByTake() {
	var takeTeacher Teacher
	DB.Take(&takeTeacher)
	fmt.Println("takeTeacher =>", takeTeacher)
}

func queryByPrimaryKey() {
	var primaryKeyTeacher Teacher
	DB.Take(&primaryKeyTeacher, 8)
	fmt.Println("primaryKeyTeacher =>", primaryKeyTeacher)
}

func queryAll() {
	var allT []Teacher
	DB.Find(&allT)
	fmt.Println("allT =>", allT)
}

func queryById() {
	id := 6
	var allT []Teacher
	DB.Find(&allT, id)
	fmt.Println("find by id's result allT =>", allT)
}

func queryByCondition() {
	//var conditionT []Teacher
	//var singleTeacher Teacher
	//DB.Where("age = ?", "22").Find(&conditionT)
	//DB.Where("name <> ?", "如燕").Take(&conditionT)
	//DB.Where("name in ?", []string{"张三", "如燕"}).Find(&conditionT)
	//DB.Where("name like ?", "mul%").Find(&conditionT)
	//DB.Where("name LIKE ? AND age >= ?", "mul%", "36").Find(&conditionT)
	//DB.Where("age BETWEEN ? AND ?", "38", "60").Find(&conditionT)
	//DB.Where(&Teacher{Name: "如燕", Age: 22}).Find(&conditionT)
	//DB.Where(map[string]interface{}{"name": "如燕"}).Find(&conditionT)
	//DB.Where(&Teacher{Name: "如燕", Age: 0}).Find(&conditionT)
	//DB.Where([]int{22, 32}).Find(&conditionT)
	//DB.Where(map[string]interface{}{"name": "如燕", "age": 0}).Find(&conditionT)
	//DB.Find(&conditionT, "name = ? AND age = ?", "如燕", 22)
	//DB.First(&conditionT, "age >= ?", 22)
	//DB.Last(&conditionT, "name = ? OR age >= ?", "mul%", 20)
	//DB.Find(&conditionT, []int{10, 22})
	//DB.Find(&conditionT, map[string]interface{}{
	//	"Name": "如燕", "Age": 22,
	//})
	//DB.Find(&conditionT, Teacher{Name: "如燕"})
	//DB.Not("name LIKE ?", "mul%").Find(&conditionT)
	//DB.Not([]int{22, 36}).First(&conditionT)
	//DB.Not(&conditionT, map[string]interface{}{
	//	"Name": "如燕",
	//}).Find(&conditionT)
	//DB.Not(Teacher{Age: 18}).Find(&conditionT)
	//DB.Where("name LIKE ?", "mul%").Or("age <> ?", 18).Find(&conditionT)
	//DB.Where("name = ?", "如燕").Or("age >= 36").Find(&conditionT)
	//DB.Where("name = ?", "如燕").Or(Teacher{Age: 18}).Find(&conditionT)
	//DB.Where([]int{22, 66}).Or(map[string]interface{}{
	//	"Name": "如燕",
	//	"Age":  22,
	//}).Find(&conditionT)
	//DB.Select("name", "age").Where("age >= ?", 18).Find(&conditionT)
	//DB.Select("name").Find(&conditionT)
	//DB.Order("age desc, created_at").Find(&conditionT)
	//DB.Find()
	//DB.Limit(3).Find(&conditionT)
	//DB.Limit(2).Offset(2).Find(&conditionT)
	//limit := 2
	//page := 3
	//DB.Limit(limit).Offset((page - 1) * limit).Find(&conditionT)
	type Result struct {
		Name string
		Age  int
	}

	var res []Result
	//DB.Model(&Teacher{}).Select("name").Where("name LIKE ?", "mul%").Scan(&res)
	//DB.Table("teachers").Select("name", "age").Where("name LIKE ?", "mul%").Scan(&res)
	//type GroupResult struct {
	//	Name   string
	//	Num    int
	//	Gender int
	//}
	//var groupRes []GroupResult
	//DB.Model(&Teacher{}).
	//	Select(
	//		"GROUP_CONCAT(name, '-', age) AS name",
	//		"count(id) AS num",
	//		"gender",
	//	).
	//	Group("gender").
	//	Scan(&groupRes)
	DB.Raw("SELECT name, age FROM teachers WHERE name LIKE ?", "mul%").Scan(&res)
	fmt.Println("conditionT =>", res)
}
func main() {
	//queryByFirst()
	//queryByLast()
	//queryByTake()
	//queryByPrimaryKey()
	//queryAll()
	queryById()
	//queryByCondition()
}
