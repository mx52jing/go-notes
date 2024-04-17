package go_redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

var client *redis.Client

func TestStringValue(t *testing.T) {
	stringValue(context.TODO(), client)
}

func TestListValue(t *testing.T) {
	listValue(context.TODO(), client)
}

func TestSetValue(t *testing.T) {
	setValue(context.TODO(), client)
}

func TestZSetValue(t *testing.T) {
	zSetValue(context.TODO(), client)
}

func TestHashValue(t *testing.T) {
	hashValue(context.TODO(), client)
}

func init() {
	// 连接redis数据库
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:9379",
		Password: "", //没有密码
		DB:       0,  //redis默认会创建0-15号DB，这里使用默认的DB
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic("redis connect fail")
	}
	fmt.Printf("redis connect success\n")
}
