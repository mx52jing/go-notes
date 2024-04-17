package main

import "fmt"

func createMap() {
	var nameMap = map[int]string {
		1: "张环",
		2: "李朗",
	}
	fmt.Println(nameMap)

	ageMap := map[string]int{
		"张环": 22,
		"李朗": 66,
		"杨方": 88,
		"仁阔": 99,
	}
	fmt.Println(ageMap)


	// varable_name := make(map[key_type]value_type)
	likeMap := make(map[int]string)
	fmt.Println(likeMap)
}

func operateMap() {
	personMap := make(map[string]string)
	fmt.Println("初始化的personMap值为:", personMap)

	// 添加元素 使用 map[key] = value
	personMap["name"] = "张环"
	personMap["brother"] = "李朗"
	personMap["title"] = "八大军头"
	fmt.Println("添加元素后personMap的值为:", personMap)

	// 删除元素 使用 delete(map, key)
	delete(personMap, "name")
	fmt.Println("删除name元素后personMap的值为:", personMap)

	// 修改元素 使用 map[key] = value
	personMap["title"] = "元芳手下的八大军头"
	fmt.Println("更新title元素后personMap的值为:", personMap)

	// 获取元素  使用 map[key]可获取对应key的value值,如果key不存在,会返回其value类型的零值
	fmt.Println(personMap["title"], personMap["age"])


	// 判断 key 是否存在
	brother_value, hasbrother := personMap["brother"]
	fmt.Println(brother_value, hasbrother)

	name_value, hasName := personMap["name"]
	fmt.Println(name_value, hasName)

	for key, value := range personMap {
		fmt.Printf("key: %s, value: %s \n", key, value)
	}


	fmt.Println("personMap的长度为", len(personMap))
}

func mapReference() {
	personMap := map[string]string{
		"one": "张环",
		"two": "李朗",
		"three": "沈韬",
		"four": "肖豹",
		"five": "杨方",
		"six": "仁阔",
	}
	fmt.Println(personMap)

	person1Map := personMap;

	person1Map["five"] = "齐虎"

	fmt.Println("person1Map值为:", person1Map)
	fmt.Println("personMap值为:", personMap)

	delete(personMap, "six")
	fmt.Println("person1Map值为:", person1Map)
	fmt.Println("personMap值为:", personMap)
}

func main() {
	// createMap()
	// operateMap()
	mapReference()
}