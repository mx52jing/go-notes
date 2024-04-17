package main

func saveAll() {
	//var first Teacher
	//DB.First(&first)
	//first.Name = "First Save"
	//first.Age = 88
	//DB.Save(&first)
	//var s Teacher
	//s.Age = 98
	//DB.Save(&s)
	//DB.Save(&Teacher{ID: 11, Name: "魔灵", Age: 20, Gender: 1})
}

func updateSingleColumn() {
	//DB.Model(&Teacher{}).Where("age = ? AND gender = ?", "98", "0").Update("name", "动灵")
	//var u Teacher
	//DB.Take(&u)
	//DB.Model(&u).Update("name", "剑灵")
	s := Teacher{ID: 2, Age: 19, Name: "吴文登", Gender: 1}
	//DB.Model(&s).Select("age", "gender").Updates(Teacher{Age: 19, Name: "吴文登", Gender: 1})
	DB.Save(&s)
}

func main() {
	//saveAll()
	updateSingleColumn()
}
