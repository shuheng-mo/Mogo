package common

import "encoding/json"

// SwapTo data assignment by checking json tag
func SwapTo(request, category interface{}) (err error) {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return
	}
	return json.Unmarshal(dataByte, category)
}
