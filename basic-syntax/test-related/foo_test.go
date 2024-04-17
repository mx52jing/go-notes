package test_related

//
//import "testing"
//
////func TestFoo(t *testing.T) {
////	//result := foo(1, 2)
////	//expect := 3
////	//if result != expect {
////	//	t.Errorf("foo(1, 2) expect is %d, result is %d", expect, result)
////	//}
////	t.Run("case 1", func(t *testing.T) {
////		result := foo(1, 2)
////		expect := 3
////		if result != expect {
////			t.Errorf("foo(1, 2) expect is %d, result is %d", expect, result)
////		}
////	})
////	t.Run("case 2", func(t *testing.T) {
////		result := foo(4, 2)
////		expect := 8
////		if result != expect {
////			t.Errorf("foo(1, 2) expect is %d, result is %d", expect, result)
////		}
////	})
////	t.Run("case 3", func(t *testing.T) {
////		result := foo(8, 2)
////		expect := 10
////		if result != expect {
////			t.Errorf("foo(1, 2) expect is %d, result is %d", expect, result)
////		}
////	})
////}
//
//func TestFoo(t *testing.T) {
//	type args struct {
//		a int
//		b int
//	}
//	type tableStruct = struct {
//		name   string
//		args   args
//		expect int
//	}
//	tableData := []tableStruct{
//		{
//			name:   "case 1",
//			args:   args{a: 1, b: 2},
//			expect: 3,
//		},
//		{
//			name:   "case 2",
//			args:   args{a: 521, b: 520},
//			expect: 1041,
//		},
//		{
//			name:   "case 3",
//			args:   args{a: 100, b: 200},
//			expect: 300,
//		},
//	}
//	for _, val := range tableData {
//		t.Run(val.name, func(t *testing.T) {
//			result := foo(val.args.a, val.args.b)
//			expect := val.expect
//			if result != expect {
//				t.Errorf("foo(1, 2) expect is %d, result is %d", expect, result)
//			}
//		})
//	}
//}
