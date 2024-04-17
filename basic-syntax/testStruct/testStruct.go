package main

import (
	"fmt"
	"reflect"
)

// func createStruct() {
// 	person := struct {
// 		name   string
// 		age    int
// 		isBoss bool
// 	}{"张环", 22, false}
// 	fmt.Println(person)
// }

func createStruct() {
	type Person struct {
		name 	 string
		age  	 int
		isBoss bool
	}
	person := Person {
		name: "张环",
		age: 22,
		isBoss: false,
	}
	fmt.Println("person is:", person)

	person1 := Person {
		"张环",
		222,
		false,
	}
	fmt.Println("person1 is:", person1)

	person2 := struct {
		name 	 string
		age  	 int
		isBoss bool	
	}{
		name: "张环",
		age: 22,
		isBoss: false,	
	}
	fmt.Println("person2 is:", person2)

	person3 := struct {
		name, skill string
		age        int
	}{
		"张环",
		"棍",
		22,
	}
	fmt.Println("person3 is:", person3)
}

func zeroValue() {
	type Person struct {
		name 	 string
		age  	 int
		isBoss bool
	}
	person5 := Person {
		name: "李朗",
	}
	fmt.Println("person 5 is:", person5)

	person6 := Person{}
	fmt.Println("person6 is:", person6)
}

func operationStruct() {
	type Person struct {
		name 	 string
		age  	 int
		isBoss bool
	}
	person7 := Person {
		name: "蛇灵",
		age: 88,
		isBoss: true,
	}

	fmt.Println("person7 is:", person7)
	fmt.Println("person7 name is:", person7.name)
	fmt.Println("person7 age is:", person7.age)
	fmt.Println("person7 isBoss is:", person7.isBoss)

	person7.age = 22

	fmt.Println("person7 is:", person7)
}

func structPointer() {
	type Person struct {
		name 	 string
		age  	 int
		isBoss bool
	}
	person8 := &Person {
		name: "蛇灵",
		age: 88,
		isBoss: true,
	}
	fmt.Println("person8 name is:", person8.name)
	fmt.Println("person8 name is:", (*person8).name)
}

func anonymousField() {
	type Person struct {
		string
		int
		bool
	}
	person9 := Person {
		"蛇灵",
		88,
		true,
	}

	fmt.Println("person9 is:", person9)
	fmt.Println("person9 string is:", person9.string)
	fmt.Println("person9 bool is:", person9.bool)
}

func nestedStruct() {
	type Skill struct {
		name  string
		level int
	}
	type Person struct {
		name 	 string
		age  	 int
		isBoss bool
		skill  Skill
	}
	person10 := Person {
		name: "张环",
		age: 22,
		isBoss: false,
		skill: Skill {
			name: "棍",
			level: 2,
		},
	}
	fmt.Println("person10 is:", person10)
}

func varivaleElevation() {
	type Skill struct {
		skillName  string
		level int
	}
	type Person struct {
		name 	 string
		age  	 int
		isBoss bool
		Skill // 注意这里
	}

	person11 := Person {
		name: "张环",
		age: 22,
		isBoss: false,
		Skill: Skill {
			skillName: "枪",
			level: 2,
		},
	}

	fmt.Println("person11 is:", person11)
	fmt.Println("person11 skillName is:", person11.skillName)
}

func compareStruct() {
	type Person struct {
		name 	 string
		age  	 int
		isBoss bool
	}

	person12 := Person {
		name: "张环",
		age: 22,
		isBoss: false,
	}
	person13 := Person {
		name: "张环",
		age: 22,
		isBoss: false,
	}
	person15 := Person {
		name: "张环",
		age: 228,
		isBoss: false,
	}

	fmt.Println("person12和person13比较", person12 == person13)
	fmt.Println("person12和person13比较", reflect.DeepEqual(person12, person13))

	fmt.Println("person12和person15比较", person12 == person15)
	fmt.Println("person12和person15比较", reflect.DeepEqual(person12, person15))
}

func structMethod() {
}

type Person struct {
	name 	 string
	age  	 int
	isBoss bool
}
func (p Person) showPersonInfo() {
	fmt.Println("Person name", p.name)
	fmt.Println("Person age", p.age)
	fmt.Println("Person isBoss", p.isBoss)
}

func (p *Person) changePersonAge(age int) {
	fmt.Println("修改age前的Person为", *p)
	p.age = age
	fmt.Println("修改age后的Person为", *p)
}

func main() {
	// createStruct()
	// zeroValue()
	// operationStruct()
	// structPointer()
	// anonymousField()
	// nestedStruct()
	// varivaleElevation()
	// compareStruct()
	person17 := Person {
		name: "张环",
		age: 22,
		isBoss: false,	
	}
	person17.changePersonAge(88)
}