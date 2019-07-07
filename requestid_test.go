package requestid

import (
	"sync"
	"testing"
)

func TestSetGet(t *testing.T) {
	tests := []interface{}{
		"123456",
		"abcdefghij",
		"ABCDEFG",
		0.1234,
		0,
		-1,
		1111111111,
	}

	var wg sync.WaitGroup
	for _, v := range tests {
		wg.Add(1)
		go func(v interface{}) {
			Set(v)
			defer Delete()

			func(v interface{}) {
				got := Get()
				if v != got {
					t.Errorf("get requestid failed. got:%v, want:%v", got, v)
				}
				t.Logf("set requestid:%v", v)
			}(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func TestConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	count := 100000
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(v interface{}) {
			Set(v)
			defer Delete()

			func(v interface{}) {
				got := Get()
				if v != got {
					t.Errorf("get requestid failed. got:%v, want:%v", got, v)
				}
			}(v)
			wg.Done()
		}(i)
	}
	t.Logf("test concurrency count:%v", count)

	wg.Wait()
}

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Set(i)
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get()
	}
}

func BenchmarkDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Delete()
	}
}

func BenchmarkGetGoID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getGoID()
	}
}
