package utils

import (
	"fmt"
	"strings"
	"time"
)

func ParseDurationStr(seconds int64) string {
	tempiStr := fmt.Sprint(time.Unix(seconds, 0).Sub(time.Unix(0, 0)))
	hIndex := strings.Index(tempiStr, "h")
	mIndex := strings.Index(tempiStr, "m")

	orarioStr := make(map[string]string)
	if hIndex != -1 {
		orarioStr["ore"] = tempiStr[0 : hIndex+1]
	}
	if mIndex != -1 {
		orarioStr["minuti"] = tempiStr[hIndex+1 : mIndex+1]
	}
	orarioStr["secondi"] = tempiStr[mIndex+1:]
	var orarioCompletoStr string
	for _, v := range orarioStr {
		if v[0] != '0' {
			orarioCompletoStr += v + " "
		}
	}
	return orarioCompletoStr
}
