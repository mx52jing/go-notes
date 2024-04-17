package go_redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

// redis存储string类型的数据
func stringValue(ctx context.Context, client *redis.Client) {
	key := "serial_name"
	value := "神探狄仁杰"
	// 删除string类型的值
	defer client.Del(ctx, key)
	// 设置存储string类型的值 第四个参数表示失效时间 这里设置的是10秒失效 0 表示永不失效
	err := client.Set(ctx, key, value, time.Second*10).Err()
	checkError(err)
	// 设置失效时间还可以调用client.Expire函数，该方法对任意类型的redis value都适用
	client.Expire(ctx, key, time.Second*10)
	// 获取存储的string类型的值
	val, err := client.Get(ctx, key).Result()
	checkError(err)
	fmt.Printf("当前获取到string类型【%s】的值为【%s】\n", key, val)
}

// redis存储List类型的数据
func listValue(ctx context.Context, client *redis.Client) {
	key := "killers"
	rValue := []interface{}{"闪灵", 88, "xue-ling", "魔灵"}
	lValue := []interface{}{"肖清芳", "袁天罡"}
	// 删除list类型的值
	defer client.Del(ctx, key)
	// 存储list的值使用RPush 或者 LPush
	// RPush表示向List右侧插入，LPush表示向List左侧插入。如果List不存在会先创建
	err := client.RPush(ctx, key, rValue...).Err()
	checkError(err)
	// 左侧插入
	err = client.LPush(ctx, key, lValue...).Err()
	checkError(err)
	// 获取存储的list类型的值
	// 截取，双闭区间。LRange表示List Range，即遍历List。0表示第一个，-1表示倒数第一个
	// 获取到的list_data是个[]string，即数字88存到redis里实际上是string
	listData, err := client.LRange(ctx, key, 0, -1).Result()
	checkError(err)
	fmt.Printf("当前获取到list类型【%s】的值为%v\n", key, listData)
}

// redis存储Set类型的数据
func setValue(ctx context.Context, client *redis.Client) {
	v1Key := "setKeyV1"
	defer client.Del(ctx, v1Key)
	values := []interface{}{18, "虎敬晖", "22", "hello world", 88}
	// 向Set中添加元素 使用client.sAdd方法。
	err := client.SAdd(ctx, v1Key, values...).Err()
	checkError(err)
	// 判断Set中是否包含某个值 使用SIsMember
	numberAge := 18 // 数字18会转成string再去redis里查找
	if isAgeExist := client.SIsMember(ctx, v1Key, numberAge).Val(); isAgeExist {
		fmt.Printf("Set中包含%#v\n", numberAge)
	} else {
		fmt.Printf("Set中不包含%#v\n", numberAge)
	}
	stringAge := "18"
	if isAgeExist := client.SIsMember(ctx, v1Key, stringAge).Val(); isAgeExist {
		fmt.Printf("Set中包含%#v\n", stringAge)
	} else {
		fmt.Printf("Set中不包含%#v\n", stringAge)
	}
	// 遍历Set 使用SMembers
	for index, val := range client.SMembers(ctx, v1Key).Val() {
		fmt.Printf("Set中第【%d】项的值为%#v\n", index, val)
	}
	// 新的key
	v2Key := "setKeyV2"
	defer client.Del(ctx, v2Key)
	v2Value := []interface{}{"林青霞", "虎敬晖", 22, "JavaScript", "106"}
	err = client.SAdd(ctx, v2Key, v2Value...).Err()
	checkError(err)
	// 找出 v1Key 和 v2Key的差集
	fmt.Println("v1Key 和 v2Key的差集(v1中有v2中没有的)")
	for _, v1ToV2DiffValue := range client.SDiff(ctx, v1Key, v2Key).Val() {
		fmt.Println(v1ToV2DiffValue)
	}
	fmt.Println("v2Key 和 v1Key的差集(v2中有v1中没有的)")
	for _, v2ToV1DiffValue := range client.SDiff(ctx, v2Key, v1Key).Val() {
		fmt.Println(v2ToV1DiffValue)
	}
	// 找出 v1Key 和 v2Key的交集
	fmt.Println("v2Key 和 v1Key的交集(v1和有v2中同时存在的)")
	for _, v1ToV2InnerValue := range client.SInter(ctx, v1Key, v2Key).Val() {
		fmt.Println(v1ToV2InnerValue)
	}
}

// redis存储有序的Set
func zSetValue(ctx context.Context, client *redis.Client) {
	v1key := "v1Key"
	defer client.Del(ctx, v1key)
	// 按照Score从小到大的顺序排列
	v1Value := []redis.Z{
		{Member: "杨方", Score: 100},
		{Member: "袁天罡", Score: 98},
		{Member: "张环", Score: 92.89},
		{Member: "李朗", Score: 92},
	}
	err := client.ZAdd(ctx, v1key, v1Value...).Err()
	checkError(err)
	fmt.Println("输出有顺序的ZSet成员")
	for _, val := range client.ZRange(ctx, v1key, 0, -1).Val() {
		fmt.Println(val)
	}
}

// value是哈希表(即map)
func hashValue(ctx context.Context, client *redis.Client) {
	stu1Key := "stu1"
	stu1Value := map[string]interface{}{"Name": "血灵", "Age": 29}
	err := client.HSet(ctx, stu1Key, stu1Value).Err()
	checkError(err)
	stu2Key := "stu2"
	stu2Value := map[string]interface{}{"Name": "魔灵", "Age": 22}
	err = client.HSet(ctx, stu2Key, stu2Value).Err()
	checkError(err)
	// 获取单个字段
	stu1Name, err := client.HGet(ctx, stu1Key, "Name").Result()
	fmt.Printf("%s的【%s】字段的值为%#v\n", stu1Key, "Name", stu1Name)
	stu2Age, err := client.HGet(ctx, stu2Key, "Age").Result()
	fmt.Printf("%s的【%s】字段的值为%#v\n", stu2Key, "Age", stu2Age)
	// 获取完整的map数据
	for fieldKey, fieldVal := range client.HGetAll(ctx, stu1Key).Val() {
		fmt.Printf("%s中的【%s】字段值为%#v\n", stu1Key, fieldKey, fieldVal)
	}
}

func setMultiValue(ctx context.Context, client *redis.Client) {

}

func checkError(err error) {
	if err != nil {
		if err == redis.Nil { //读redis发生error，大部分情况是因为key不存在
			fmt.Println("当前key不存在")
			return
		}
		fmt.Println(err)
		os.Exit(1)
	}
}
