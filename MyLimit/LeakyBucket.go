package MyLimit

import "time"

type LeakyBucket struct {
	water       int `water count at refreashTime`
	rate        int
	burst       int `max burst`
	refreshTime int
}

func NewLeakBucket(rate, burst int) *LeakyBucket {
	lb := new(LeakyBucket)
	lb.water = 0
	lb.rate = rate
	lb.burst = burst
	lb.refreshTime = time.Now().Second()
	return lb
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (lb *LeakyBucket) refreshWater() {
	now := time.Now().Second()
	//到目前为止， now-refreshTime 时间内,当前water所剩的水至少以rate/s 速度流失
	lb.water = max(0, lb.water-(now-lb.refreshTime)*lb.rate)
}

func (lb) grantWater() bool {
	lb.refreshTime()
	if lb.water < lb.burst {
		//没潢+1
		lb.water++
		return true
	}
	return false
}
