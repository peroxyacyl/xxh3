package xxh3

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkFixed(b *testing.B) {
	r := func(i int) {
		bench := func(b *testing.B) {
			b.SetBytes(int64(i))
			var acc uint64
			d := string(make([]byte, i))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				acc = HashString(d)
			}
			runtime.KeepAlive(acc)
		}
		if i > 240 {
			avx2Orig, sse2Orig, cleanup := override()
			defer cleanup()
			if avx2Orig {
				avx2, sse2 = true, false
				b.Run(fmt.Sprintf("%d-AVX2", i), bench)
			}
			if sse2Orig {
				avx2, sse2 = false, true
				b.Run(fmt.Sprintf("%d-SSE2", i), bench)
			}

			avx2, sse2 = false, false
		}

		b.Run(fmt.Sprintf("%d", i), bench)
	}

	r(0)
	r(1)
	r(2)
	r(3)
	r(4)
	r(8)
	r(9)
	r(16)
	r(17)
	r(32)
	r(33)
	r(64)
	r(65)
	r(96)
	r(97)
	r(128)
	r(129)
	r(240)
	r(241)
	r(512)
	r(1024)
	r(8192)
	r(100 * 1024)
}
