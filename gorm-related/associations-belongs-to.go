package main

import "fmt"

/**
belongsTo属于一对一的关系
类似于自作多情的状况，一方会在自己这边记住对方，而对方可以毫不知情
*/

/**
下面结构体声明时，舔🐶在自己这边记住了女神和女神的id，但是声明的女神结构体可以不声明关于舔🐶的任何信息
使用DB.AutoMigrate(&Dog{})创建表时，会自动创建dogs表和girl_gods表
*/

type Dog struct {
	ID        uint
	Name      string
	GirlGod   GirlGod
	GirlGodID uint // 外键，就是GirlGod.ID
}

type GirlGod struct {
	ID   uint
	Name string
}

func dogCreate() {
	g6 := GirlGod{
		Name: "女神6号",
	}
	d5 := Dog{
		Name:    "狗哥5号",
		GirlGod: g6,
	}
	DB.Create(&d5)
}

func dogPreCreate() {
	// 先创建一些GirlGod数据
	girlGods := []GirlGod{
		{Name: "女神1号"},
		{Name: "女神2号"},
		{Name: "女神3号"},
		{Name: "女神4号"},
		{Name: "女神5号"},
		{Name: "女神6号"},
	}
	DB.Model(&GirlGod{}).Create(&girlGods)
	var girlGod1 GirlGod
	var girlGod2 GirlGod
	DB.Model(&GirlGod{}).Find(&girlGod1, 1)
	DB.Model(&GirlGod{}).Find(&girlGod2, 2)
	dog1 := Dog{Name: "狗哥1号", GirlGod: girlGod1}
	dog2 := Dog{Name: "狗哥2号", GirlGod: girlGod2}
	dogs := []Dog{dog1, dog2}
	DB.Model(&Dog{}).Create(&dogs)
}

// 操作关联
func operateAssociations() {
	//var dog1 Dog
	//DB.Model(&Dog{}).Find(&dog1, 1)
	//var girlGod6 GirlGod
	//DB.Model(&GirlGod{}).Find(&girlGod6, 6)
	//// 将狗哥1号关联的女神1号替换为女神6号
	//err := DB.Model(&dog1).Association("GirlGod").Append(&girlGod6)
	//if err != nil {
	//	return
	//}
	//var dog2 Dog
	//DB.Model(&Dog{}).Find(&dog2, 2)
	//var girlGod5 GirlGod
	//DB.Model(&GirlGod{}).Find(&girlGod5, 5)
	//// 将狗哥2号关联的女神2号替换为女神5号
	//err := DB.Model(&dog2).Association("GirlGod").Replace(&girlGod5)
	//if err != nil {
	//	return
	//}
	//var dog2 Dog
	//DB.Model(&Dog{}).Find(&dog2, 2)
	//var girlGod5 GirlGod
	//DB.Model(&GirlGod{}).Find(&girlGod5, 5)
	//var girlGod2 GirlGod
	//DB.Model(&GirlGod{}).Find(&girlGod2, 2)
	//// 将狗哥2号关联的女神2号替换为女神5号
	//err := DB.Model(&dog2).Association("GirlGod").Replace(&girlGod5, &girlGod2)
	//if err != nil {
	//	return
	//}
	//var dog1 Dog
	//DB.Model(&Dog{}).Find(&dog1, 1)
	//var girlGod6 GirlGod
	//DB.Model(&GirlGod{}).Find(&girlGod6, 6)
	//// 删除狗哥1号关联的女神6号
	//err := DB.Model(&dog1).Association("GirlGod").Delete(&girlGod6)
	//if err != nil {
	//	return
	//}
	var dog2 Dog
	DB.Model(&Dog{}).Find(&dog2, 2)
	// 将狗哥2号关联的女神清除
	err := DB.Model(&dog2).Association("GirlGod").Clear()
	if err != nil {
		return
	}
}

func setAssociation() {
	var dog1 Dog
	DB.Model(&Dog{}).Find(&dog1, 1)
	var dog2 Dog
	DB.Model(&Dog{}).Find(&dog2, 2)
	var girlGod6 GirlGod
	DB.Model(&GirlGod{}).Find(&girlGod6, 6)
	var girlGod5 GirlGod
	DB.Model(&GirlGod{}).Find(&girlGod5, 5)
	DB.Model(&dog1).Association("GirlGod").Append(&girlGod6)
	DB.Model(&dog2).Association("GirlGod").Append(&girlGod5)
}

// 查询关联
func queryAssociation() {
	var dog1 Dog
	DB.Preload("GirlGod").Find(&dog1, 1)
	fmt.Printf("dog1 is %v\n", dog1)
}

func main() {
	//DB.AutoMigrate(&Dog{})
	//dogCreate()
	//dogPreCreate()
	//operateAssociations()
	//setAssociation()
	queryAssociation()
}

/**
创建的dogs表如下
CREATE TABLE dogs
(
    id          bigint UNSIGNED AUTO_INCREMENT
        PRIMARY KEY,
    name        longtext        NULL,
    girl_god_id bigint UNSIGNED NULL,
    CONSTRAINT fk_dogs_girl_god
        FOREIGN KEY (girl_god_id) REFERENCES girl_gods (id)
);

创建的girl_gods表如下

CREATE TABLE girl_gods
(
    id   bigint UNSIGNED AUTO_INCREMENT
        PRIMARY KEY,
    name longtext NULL
);
*/
