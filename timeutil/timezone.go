package timeutil

import (
	"os"
	"strconv"
	"time"

	"github.com/ChiKangMa/go-support/debug"
)

func init() {
	debug.PrintInfo("Set timezone to Asia/Shanghai")
	os.Setenv("TZ", "Asia/Shanghai")
}
func ConvertUnixTime(unixTime string) string {
	intUnixTime, err := strconv.ParseInt(unixTime, 10, 64)
	debug.PrintError(err)
	return time.Unix(intUnixTime, 0).Format(TimeLayout)
}
