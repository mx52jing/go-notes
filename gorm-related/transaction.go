package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type TDog struct {
	Name string
}

func tAutoMigrate() {
	err := DB.AutoMigrate(&TDog{})
	if err != nil {
		return
	}
}

func normalTransaction() {
	err := DB.Transaction(func(tx *gorm.DB) error {
		d1 := TDog{Name: "狗哥1号"}
		d2 := TDog{Name: "狗哥2号"}
		dogs := []TDog{d1, d2}
		tx.Model(&TDog{}).Create(&dogs)
		return nil
	})
	if err != nil {
		return
	}
}

func normalErrorTransaction() {
	err := DB.Transaction(func(tx *gorm.DB) error {
		d3 := TDog{Name: "狗哥3号"}
		d4 := TDog{Name: "狗哥4号"}
		dogs := []TDog{d3, d4}
		tx.Model(&TDog{}).Create(&dogs)
		return errors.New("错误")
	})
	if err != nil {
		return
	}
}

func nestTransaction() {
	err := DB.Transaction(func(tx *gorm.DB) error {
		d5 := TDog{Name: "狗哥5号"}
		tx.Model(&TDog{}).Create(&d5)
		err := tx.Transaction(func(tx *gorm.DB) error {
			d6 := TDog{Name: "狗哥6号"}
			tx.Model(&TDog{}).Create(&d6)
			return errors.New("内部事务出错")
		})
		if err != nil {
			// 暂时不处理这个错误 让外部的事务可以顺利执行
			//return err
			fmt.Printf("内部事务报错 %s\n", err.Error())
		}
		return nil
	})
	if err != nil {
		return
	}
}

func manualTransaction() {
	tx := DB.Begin()
	d7 := TDog{Name: "狗哥7号"}
	d8 := TDog{Name: "狗哥8号"}
	tx.Model(&TDog{}).Create(&d7)
	tx.Model(&TDog{}).Create(&d8)
	tx.Commit()
}

func manualTransaction2() {
	tx := DB.Begin()
	d9 := TDog{Name: "狗哥9号"}
	d10 := TDog{Name: "狗哥10号"}
	tx.Model(&TDog{}).Create(&d9)
	tx.Commit()
	tx.Model(&TDog{}).Create(&d10)
}

func manualTransaction3() {
	tx := DB.Begin()
	d11 := TDog{Name: "狗哥11号"}
	d12 := TDog{Name: "狗哥12号"}
	dogs := []TDog{d11, d12}
	tx.Create(&dogs)
	tx.Commit()
	tx.Rollback()
}

func rollbackPoint() {
	tx := DB.Begin()
	d13 := TDog{Name: "狗哥13号"}
	d14 := TDog{Name: "狗哥14号"}
	tx.Model(&TDog{}).Create(&d13)
	pointName := "f"
	tx.SavePoint(pointName)
	tx.Model(&TDog{}).Create(&d14)
	tx.RollbackTo(pointName)
	tx.Commit()
}

func rollbackPoint2() {
	tx := DB.Begin()
	d15 := TDog{Name: "狗哥15号"}
	d16 := TDog{Name: "狗哥16号"}
	d17 := TDog{Name: "狗哥17号"}
	tx.Model(&TDog{}).Create(&d15)
	pointName := "f"
	tx.SavePoint(pointName)
	tx.Model(&TDog{}).Create(&d16)
	tx.RollbackTo(pointName)
	tx.Model(&TDog{}).Create(&d17)
	tx.Commit()
}

func main() {
	//tAutoMigrate()
	//normalTransaction()
	//normalErrorTransaction()
	//nestTransaction()
	//manualTransaction()
	//manualTransaction2()
	//manualTransaction3()
	//rollbackPoint()
	rollbackPoint2()
}
