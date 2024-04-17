package main

func createOneTeacher() {
	teacher := Teacher{
		Name:   "李朗",
		Age:    28,
		Gender: 0,
	}
	DB.Create(&teacher)
}

func createMultiTeacher() {
	teachers := []Teacher{
		{Name: "杨方", Age: 16, Gender: 1},
		{Name: "崔亮", Age: 36, Gender: 1},
		{Name: "仁阔", Age: 26, Gender: 1},
		{Name: "如燕", Age: 19, Gender: 0},
		{Name: "齐虎", Age: 22, Gender: 1},
		{Name: "黑衣天王", Age: 16, Gender: 0},
		{Name: "潘越", Age: 28, Gender: 1},
		{Name: "宁氏", Age: 18, Gender: 0},
	}
	DB.Create(&teachers)
}

func createBySelectedField() {
	selectedFieldTeacher := Teacher{
		Name:   "如燕",
		Age:    18,
		Gender: 1,
	}
	DB.Select("Name", "Age").Create(&selectedFieldTeacher)
}

func createByOmitField() {
	selectedFieldTeacher := Teacher{
		Name:   "如燕1号",
		Age:    19,
		Gender: 1,
	}
	DB.Omit("Age", "Created").Create(&selectedFieldTeacher)
}

func createByMap() {
	teacherMap := map[string]interface{}{
		"Name":   "mapName",
		"Age":    36,
		"Gender": 1,
	}
	DB.Model(&Teacher{}).Create(teacherMap)
}

func createByMultiMap() {
	multiMap := []map[string]interface{}{
		{"Name": "multiMap1", "Age": 22},
		{"Name": "multiMap2", "Age": 32},
		{"Name": "multiMap3", "Age": 42},
		{"Name": "multiMap4", "Age": 52},
	}
	DB.Model(&Teacher{}).Create(multiMap)
}

func main() {
	//DB.AutoMigrate(&Teacher{}) //
	//createOneTeacher()
	createMultiTeacher()
	//createBySelectedField()
	//createByOmitField()
	//createByMap()
	//createByMultiMap()
}
