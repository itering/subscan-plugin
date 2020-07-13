package tools

import (
	"encoding/json"
	"strconv"
	"strings"
)

func AddHex(s string) string {
	if strings.TrimSpace(s) == "" {
		return ""
	}
	if strings.HasPrefix(s, "0x") {
		return s
	}
	return strings.ToLower("0x" + s)
}

func StringToInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}

func IntFromInterface(i interface{}) int {
	switch i := i.(type) {
	case int:
		return i
	case int64:
		return int(i)
	case uint64:
		return int(i)
	case float64:
		return int(i)
	case string:
		return StringToInt(i)
	}
	return 0
}

func UnmarshalToAnything(r interface{}, raw interface{}) {
	switch raw := raw.(type) {
	case string:
		_ = json.Unmarshal([]byte(raw), &r)
	case []uint8:
		_ = json.Unmarshal(raw, &r)
	default:
		b, _ := json.Marshal(raw)
		_ = json.Unmarshal(b, r)
	}
}
