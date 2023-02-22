package basic

import "fmt"

func CapTrack_() {
	s := make([]int, 0)

	oldCap := cap(s)
	for i := 0; i < 2048; i++ {
		s = append(s, i)

		newCap := cap(s)
		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}
	/*
		运行结果(1.18版本之前)：
		[0 ->   -1] cap = 0     |  after append 0     cap = 1
		[0 ->    0] cap = 1     |  after append 1     cap = 2
		[0 ->    1] cap = 2     |  after append 2     cap = 4
		[0 ->    3] cap = 4     |  after append 4     cap = 8
		[0 ->    7] cap = 8     |  after append 8     cap = 16
		[0 ->   15] cap = 16    |  after append 16    cap = 32
		[0 ->   31] cap = 32    |  after append 32    cap = 64
		[0 ->   63] cap = 64    |  after append 64    cap = 128
		[0 ->  127] cap = 128   |  after append 128   cap = 256
		[0 ->  255] cap = 256   |  after append 256   cap = 512		// 小于1024扩容为2倍
		[0 ->  511] cap = 512   |  after append 512   cap = 1024
		[0 -> 1023] cap = 1024  |  after append 1024  cap = 1280	// 超过1024扩容为1.25倍
		[0 -> 1279] cap = 1280  |  after append 1280  cap = 1696
		[0 -> 1695] cap = 1696  |  after append 1696  cap = 2304

		运行结果(1.18版本)：
		[0 ->   -1] cap = 0     |  after append 0     cap = 1
		[0 ->    0] cap = 1     |  after append 1     cap = 2
		[0 ->    1] cap = 2     |  after append 2     cap = 4
		[0 ->    3] cap = 4     |  after append 4     cap = 8
		[0 ->    7] cap = 8     |  after append 8     cap = 16
		[0 ->   15] cap = 16    |  after append 16    cap = 32
		[0 ->   31] cap = 32    |  after append 32    cap = 64
		[0 ->   63] cap = 64    |  after append 64    cap = 128		// 小于256扩容为2倍
		[0 ->  127] cap = 128   |  after append 128   cap = 256
		[0 ->  255] cap = 256   |  after append 256   cap = 512		// 超过256时newCap = oldCap + (oldCap + 256*3)/4
		[0 ->  511] cap = 512   |  after append 512   cap = 848
		[0 ->  847] cap = 848   |  after append 848   cap = 1280
		[0 -> 1279] cap = 1280  |  after append 1280  cap = 1792
		[0 -> 1791] cap = 1792  |  after append 1792  cap = 2560
	*/
}
