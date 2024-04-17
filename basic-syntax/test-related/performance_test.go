package test_related

import "testing"

func BenchmarkHandleJsonByJsonLibrary(b *testing.B) {
	b.Run("Benchmark HandleJsonByJsonLibrary", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			handleJsonByJsonLibrary()
			//if err != nil {
			//	b.Errorf("HandleJsonByJsonLibrary err %s", err)
			//	return
			//}
			//if data.Id != C1.Id || len(data.Students) != len(C1.Students) {
			//	b.Error("HandleJsonByJsonLibrary fail")
			//	return
			//}
		}
		//b.Logf("HandleJsonByJsonLibrary success")
	})
}

func BenchmarkHandleJsonBySonicLibrary(b *testing.B) {
	b.Run("Benchmark HandleJsonBySonicLibrary", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			handleJsonBySonicLibrary()
			//if err != nil {
			//	b.Errorf("HandleJsonBySonicLibrary err %s", err)
			//	return
			//}
			//if data.Id != C1.Id || len(data.Students) != len(C1.Students) {
			//	b.Error("HandleJsonBySonicLibrary fail")
			//	return
			//}
		}
		//b.Logf("HandleJsonBySonicLibrary success")
	})
}
