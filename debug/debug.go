package debug

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var debugID string

func init() {
	debugID = uuid.New().String()
}

func currentTime() string {
	// return ""
	return time.Now().Format("2006-01-02 15:04:05")
}

func PrintError(err error) {
	if err != nil {
		fmt.Println(debugID+" (ERROR): "+err.Error(), currentTime())
	}
}

func PrintLog(msg ...interface{}) {
	fmt.Println(debugID+" (LOG): ", msg, "on", currentTime())
}

func PrintInfo(msg ...interface{}) {
	// fmt.Println(debugID+" (INFO): ", msg, "on", currentTime())
}

func PrintInfoValue(x interface{}) {
	fmt.Printf("INFO: %+v\n", x)
}

func PrintDebug(msg ...interface{}) {
	// fmt.Println(debugID + " (DEBUG): ", msg, " on ", currentTime())
}

func PrintValue(x interface{}) {
	// fmt.Printf("DEBUG: %+v\n", x)
}
