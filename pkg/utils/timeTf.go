package utils

var MonthMap map[string]string = map[string]string{
	"Jan": "01",
	"Feb": "02",
	"Mar": "03",
	"Apr": "04",
	"May": "05",
	"Jun": "06",
	"Jul": "07",
	"Aug": "08",
	"Sep": "09",
	"Oct": "10",
	"Nov": "11",
	"Dec": "12",
}

func TimeTransform(yearStr string, monthStr string, dayStr string, timeStr string) (dateStr string) {
	for k, v := range MonthMap {
		if k == monthStr {
			dateStr := yearStr + "-" + v + "-" + dayStr + " " + timeStr
			return dateStr
		}
	}
	return "null"
}
