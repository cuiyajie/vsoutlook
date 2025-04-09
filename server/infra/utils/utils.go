package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/mail"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func SmsCode() string {
	return RandStringRunes(6, Number)
}

func Hash(length int) string {
	if length == 0 {
		length = 10
	}
	return RandStringRunes(length, AlphaNumber)
}

var (
	AlphaNumber = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	Number      = []rune("1234567890")
)

func RandStringRunes(n int, letterRunes []rune) string {
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// ChinaMobilePhoneReg https://github.com/VincentSit/ChinaMobilePhoneNumberRegex/blob/master/README-CN.md
var ChinaMobilePhoneReg = regexp.MustCompile(`(?m)^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[0-35-9]\d{2}|4(?:0\d|1[0-2]|9\d))|9[0-35-9]\d{2}|6[2567]\d{2}|4[579]\d{2})\d{6}$`)

func IsValidChinaMobile(mobile string) bool {
	return ChinaMobilePhoneReg.MatchString(mobile)
}

// todo: tagName 需要对逗号做处理
func ToMap(in any, tagName string) (Map, error) {
	out := make(Map)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func IsValidEmail(email string) bool {
	if len(email) == 0 {
		return false
	}
	if _, err := mail.ParseAddress(email); err == nil {
		if strings.Contains(strings.Split(email, "@")[1], ".") {
			return true
		}
	}
	return false
}

func IsASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}

	return true
}

// abcd.prt.4 has ext: .prt.4
func Ext(s string) string {
	if strings.Count(s, ".") <= 1 || !strings.Contains(s, ".prt.") {
		return filepath.Ext(s)
	}
	parts := strings.Split(s, ".")
	l := len(parts)
	if l >= 2 && parts[l-2] == "prt" {
		return ".prt." + parts[l-1]
	}
	return filepath.Ext(s)
}

func ParseCoreListString(input string) []int {
	var numberArray []int
	encountered := make(map[int]bool)

	segments := strings.Split(input, ",")
	for _, segment := range segments {
		if strings.Contains(segment, "-") {
			rangeValues := strings.Split(segment, "-")
			start, _ := strconv.Atoi(rangeValues[0])
			end, _ := strconv.Atoi(rangeValues[1])
			for i := start; i <= end; i++ {
				if !encountered[i] {
					numberArray = append(numberArray, i)
					encountered[i] = true
				}
			}
		} else {
			num, _ := strconv.Atoi(segment)
			if !encountered[num] {
				numberArray = append(numberArray, num)
				encountered[num] = true
			}
		}
	}
	return numberArray
}

func IntArrayToString(arr []uint32) string {
	var strArr []string
	for _, num := range arr {
		strArr = append(strArr, fmt.Sprintf("%d", num))
	}
	return strings.Join(strArr, ",")
}

func MapToString(m map[string]interface{}) string {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func KeysOf(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func IsValidIPv4WithPort(input string) bool {
	host, port, err := net.SplitHostPort(input)
	if err != nil {
		return false
	}

	ip := net.ParseIP(host)
	if ip == nil || ip.To4() == nil {
		return false
	}

	_, err = strconv.Atoi(port)

	return err == nil
}

func MergeMapArrays[T any](m map[string][]T) []T {
	var result []T
	for _, v := range m {
		result = append(result, v...)
	}
	return result
}

func FormatUint32Array(arr []uint32) string {
	// 1. 排序数组
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// 2. 遍历数组，生成 chunk
	var chunks []string
	start := arr[0]
	prev := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] == prev+1 {
			prev = arr[i]
		} else {
			if start == prev {
				chunks = append(chunks, strconv.FormatUint(uint64(start), 10))
			} else {
				chunks = append(chunks, fmt.Sprintf("%d-%d", start, prev))
			}
			start = arr[i]
			prev = arr[i]
		}
	}

	// 处理最后一个 chunk
	if start == prev {
		chunks = append(chunks, strconv.FormatUint(uint64(start), 10))
	} else {
		chunks = append(chunks, fmt.Sprintf("%d-%d", start, prev))
	}

	// 3. 用 "," 连接所有 chunk
	return strings.Join(chunks, ",")
}

func ConvertToIntToFloat64(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case int:
		return float64(v), true
	case int8:
		return float64(v), true
	case int16:
		return float64(v), true
	case int32:
		return float64(v), true
	case int64:
		return float64(v), true
	default:
		return 0.0, false
	}
}
