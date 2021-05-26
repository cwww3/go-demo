package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hashCode(s string) int {
	var hash int32 = 0
	if len(s) == 0 {
		return int(hash)
	}
	for _, c := range s {
		hash = ((hash << 5) - hash) + int32(c)
		hash &= hash // Convert to 32bit integer
	}
	return int(hash)
}
func main() {
	str := "[{\"scene\":\"SystemInfo\",\"extra\":\"\\\"{\\\\\\\"benchmarkLevel\\\\\\\":10,\\\\\\\"language\\\\\\\":\\\\\\\"zh_CN\\\\\\\",\\\\\\\"wifiEnabled\\\\\\\":false,\\\\\\\"safeArea\\\\\\\":{\\\\\\\"bottom\\\\\\\":667,\\\\\\\"height\\\\\\\":647,\\\\\\\"top\\\\\\\":20,\\\\\\\"width\\\\\\\":375,\\\\\\\"left\\\\\\\":0,\\\\\\\"right\\\\\\\":375},\\\\\\\"bluetoothEnabled\\\\\\\":false,\\\\\\\"locationAuthorized\\\\\\\":true,\\\\\\\"deviceOrientation\\\\\\\":\\\\\\\"portrait\\\\\\\",\\\\\\\"notificationSoundAuthorized\\\\\\\":true,\\\\\\\"screenHeight\\\\\\\":667,\\\\\\\"windowHeight\\\\\\\":555,\\\\\\\"version\\\\\\\":\\\\\\\"8.0.2\\\\\\\",\\\\\\\"fontSizeSetting\\\\\\\":18,\\\\\\\"system\\\\\\\":\\\\\\\"iOS 14.4.1\\\\\\\",\\\\\\\"locationReducedAccuracy\\\\\\\":false,\\\\\\\"notificationAuthorized\\\\\\\":true,\\\\\\\"pixelRatio\\\\\\\":2,\\\\\\\"statusBarHeight\\\\\\\":20,\\\\\\\"notificationBadgeAuthorized\\\\\\\":true,\\\\\\\"windowWidth\\\\\\\":375,\\\\\\\"locationEnabled\\\\\\\":true,\\\\\\\"model\\\\\\\":\\\\\\\"iPhone 6s\\u003ciPhone8,1\\u003e\\\\\\\",\\\\\\\"batteryLevel\\\\\\\":92,\\\\\\\"screenWidth\\\\\\\":375,\\\\\\\"microphoneAuthorized\\\\\\\":true,\\\\\\\"cameraAuthorized\\\\\\\":true,\\\\\\\"albumAuthorized\\\\\\\":true,\\\\\\\"notificationAlertAuthorized\\\\\\\":true,\\\\\\\"brand\\\\\\\":\\\\\\\"iPhone\\\\\\\",\\\\\\\"platform\\\\\\\":\\\\\\\"ios\\\\\\\",\\\\\\\"SDKVersion\\\\\\\":\\\\\\\"2.16.0\\\\\\\",\\\\\\\"enableDebug\\\\\\\":false,\\\\\\\"devicePixelRatio\\\\\\\":2,\\\\\\\"host\\\\\\\":{\\\\\\\"env\\\\\\\":\\\\\\\"WeChat\\\\\\\",\\\\\\\"appId\\\\\\\":\\\\\\\"\\\\\\\",\\\\\\\"version\\\\\\\":402653753}}\\\"\",\"key\":\"\"},{\"scene\":\"open1014\",\"extra\":\"{\\\"path\\\":\\\"pages/index/share\\\",\\\"query\\\":{\\\"type\\\":\\\"track\\\",\\\"name\\\":\\\"subsMsgCheckIn\\\"},\\\"scene\\\":1014,\\\"referrerInfo\\\":{},\\\"mode\\\":\\\"default\\\",\\\"ifStrictNoPay\\\":true}\",\"key\":\"\"},{\"scene\":\"lazyBegin\",\"extra\":\"{\\\"todoCnt\\\":0,\\\"wxsdk\\\":\\\"2.16.0\\\"}\",\"key\":\"\"},{\"scene\":\"lazyEnd\",\"extra\":\"{\\\"appendCnt\\\":0,\\\"prependCnt\\\":0,\\\"elsped\\\":1}\",\"key\":\"\"},{\"scene\":\"btnPage\",\"extra\":\"\",\"key\":\"pages/index/share\"}]" + "Ar32ifhewiuhewiuw"
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(str)), `\\u`, `\u`, -1))
	if err != nil {
		fmt.Println(err)
	}
	s := hashCode(str)

	old := hashCodeLegacy(str)
	fmt.Println(s)
	fmt.Println(old)
}

func hashCodeLegacy(s string) int {
	var hash int32 = 0
	if len(s) == 0 {
		return int(hash)
	}
	var c byte
	for i := 0; i < len(s); i++ {
		c = s[i]
		hash = ((hash << 5) - hash) + int32(c)
		hash &= hash // Convert to 32bit integer
	}
	return int(hash)
}
