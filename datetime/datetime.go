package datetime

import (
	"time"
)

var timeFormat map[string]string

func init() {
	timeFormat = map[string]string{
		"yyyy-mm-dd hh:mm:ss": "2006-01-02 15:04:05",
		"yyyy-mm-dd hh:mm":    "2006-01-02 15:04",
		"yyyy-mm-dd hh":       "2006-01-02 15",
		"yyyy-mm-dd":          "2006-01-02",
		"yyyy-mm":             "2006-01",
		"mm-dd":               "01-02",
		"dd-mm-yy hh:mm:ss":   "02-01-06 15:04:05",
		"yyyy/mm/dd hh:mm:ss": "2006/01/02 15:04:05",
		"yyyy/mm/dd hh:mm":    "2006/01/02 15:04",
		"yyyy/mm/dd hh":       "2006/01/02 15",
		"yyyy/mm/dd":          "2006/01/02",
		"yyyy/mm":             "2006/01",
		"mm/dd":               "01/02",
		"dd/mm/yy hh:mm:ss":   "02/01/06 15:04:05",
		"yyyymmdd":            "20060102",
		"mmddyy":              "010206",
		"yyyy":                "2006",
		"yy":                  "06",
		"mm":                  "01",
		"hh:mm:ss":            "15:04:05",
		"hh:mm":               "15:04",
		"mm:ss":               "04:05",
	}
}

func AddMinute(t time.Time, minutes int64) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func AddHour(t time.Time, hours int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func AddDay(t time.Time, days int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func AddWeek(t time.Time, weeks int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func AddMonth(t time.Time, months int64) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func AddYear(t time.Time, year int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func AddDaySafe(t time.Time, days int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func AddMonthSafe(t time.Time, months int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func AddYearSafe(t time.Time, years int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func GetNowDate() string { _ = "STUB: not implemented"; return "" }

func GetNowTime() string { _ = "STUB: not implemented"; return "" }

func GetNowDateTime() string { _ = "STUB: not implemented"; return "" }

func GetTodayStartTime() string { _ = "STUB: not implemented"; return "" }

func GetTodayEndTime() string { _ = "STUB: not implemented"; return "" }

func GetZeroHourTimestamp() int64 { _ = "STUB: not implemented"; return 0 }

func GetNightTimestamp() int64 { _ = "STUB: not implemented"; return 0 }

func FormatTimeToStr(t time.Time, format string, timezone ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func FormatStrToTime(str, format string, timezone ...string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func BeginOfMinute(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func EndOfMinute(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func BeginOfHour(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func EndOfHour(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func BeginOfDay(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func EndOfDay(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func BeginOfWeek(t time.Time, beginFrom time.Weekday) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func EndOfWeek(t time.Time, endWith time.Weekday) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func BeginOfMonth(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func EndOfMonth(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func BeginOfYear(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func EndOfYear(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func IsLeapYear(year int) bool { _ = "STUB: not implemented"; return false }

func BetweenSeconds(t1 time.Time, t2 time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func DayOfYear(t time.Time) int { _ = "STUB: not implemented"; return 0 }

func IsWeekend(t time.Time) bool { _ = "STUB: not implemented"; return false }

func NowDateOrTime(format string, timezone ...string) string { _ = "STUB: not implemented"; return "" }

func Timestamp(timezone ...string) int64 { _ = "STUB: not implemented"; return 0 }

func TimestampMilli(timezone ...string) int64 { _ = "STUB: not implemented"; return 0 }

func TimestampMicro(timezone ...string) int64 { _ = "STUB: not implemented"; return 0 }

func TimestampNano(timezone ...string) int64 { _ = "STUB: not implemented"; return 0 }

func TrackFuncTime(pre time.Time) func() { _ = "STUB: not implemented"; return nil }

func getCallerName() string { _ = "STUB: not implemented"; return "" }

func DaysBetween(start, end time.Time) int { _ = "STUB: not implemented"; return 0 }

func GenerateDatetimesBetween(start, end time.Time, layout string, interval string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Min(t1 time.Time, times ...time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func Max(t1 time.Time, times ...time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func MaxMin(t1 time.Time, times ...time.Time) (maxTime time.Time, minTime time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time)
}
