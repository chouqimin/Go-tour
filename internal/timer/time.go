
package timer

import "time"

var Location, _ = time.LoadLocation("Asia/Taipei")

func GetNowTime() time.Time {
	return time.Now().In(Location) // 確保取得的時間與期望的時區一致
	// return time.Now()
}

// 根據字串加減給定時間
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	// d:duration(持續時間) 
	// 支援的有效時間單位為"ns"、"us" (or "µs"), "ms", "s", "m", "h"，如:"300ms"、"-1.5h" 或 "2h45m"
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}
