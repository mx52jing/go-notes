package main

import (
	"context"
	"fmt"
	"time"
)

func withTimeoutFn() {
	parentCtx := context.TODO()
	child, cancel := context.WithTimeout(parentCtx, 500 * time.Millisecond)
	defer cancel()
	select {
	case <-child.Done():
		err := child.Err()
		fmt.Printf("结束了，原因是%s", err)
	case <-time.After(600 * time.Millisecond):
		fmt.Println("超过600ms")
	}
}

func timeoutInherit() {
	parentCtx, cancel1 := context.WithTimeout(context.TODO(), 500 * time.Millisecond) // 500ms 后会超时
	t0 := time.Now()
	defer cancel1()
	time.Sleep(300 * time.Millisecond) // 睡眠300ms
	childCtx, cancel2 := context.WithTimeout(parentCtx, 100 * time.Millisecond) // 100ms后会超时
	t1 := time.Now()
	defer cancel2()
	select {
	case <-parentCtx.Done():
		t2 := time.Now()
		err := parentCtx.Err()
		fmt.Printf("parentCtx超时，原因为：%v，parentCtx运行时间为%v, childCtx运行时间为%v", err, t2.Sub(t0), t2.Sub(t1))
	case <-childCtx.Done():
		t2 := time.Now()
		err := childCtx.Err()
		fmt.Printf("childCtx超时，原因为：%v，parentCtx运行时间为%v, childCtx运行时间为%v", err, t2.Sub(t0), t2.Sub(t1))
	}
}


func withCancelFn() {
	parentCtx := context.Background()
	childCtx, cancel := context.WithCancel(parentCtx)
	go func ()  {
		time.Sleep(300 * time.Millisecond)
		cancel()
	}()
	select {
	case <-childCtx.Done():
		err := childCtx.Err()
		fmt.Printf("结束了，原因是%s", err)	
	case <-time.After(500 * time.Millisecond):
		fmt.Println("超过500ms")	
	}
}

func withValueFn() {
	parentCtx := context.TODO()
	childCtx := context.WithValue(parentCtx, "name", "南山")
	name := childCtx.Value("name")
	fmt.Printf("获取到的name为%s", name)
}


func addAge(ctx context.Context) context.Context {
	childCtx := context.WithValue(ctx, "age", 22)
	return childCtx
}

func addLike(ctx context.Context) context.Context {
	childCtx := context.WithValue(ctx, "like", []string{"篮球","乒乓球"})
	return childCtx;
}

func withValueInherit() {
	parentCtx := context.TODO()
	childCtx := context.WithValue(parentCtx, "name", "南山")
	child1Ctx := addAge(childCtx)
	fmt.Printf("child1的name为%s,age为%d\n", child1Ctx.Value("name"), child1Ctx.Value("age"))
	child2Ctx := addLike(child1Ctx)
	fmt.Printf("child2的name为%s,age为%d, like为%v\n", child2Ctx.Value("name"), child2Ctx.Value("age"),  child2Ctx.Value("like"))
}

func main() {
	// withTimeoutFn()
	// withCancelFn()
	// withValueFn()
	// withValueInherit()
	timeoutInherit()
}