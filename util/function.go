package util

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/godcong/wego/log"
	"github.com/satori/go.uuid"
)

/*CustomHeader xml header*/
const CustomHeader = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>`

/*CDATA xml cdata defines */
type CDATA struct {
	Value string `xml:",cdata"`
}

/* error types */
var (
	ErrorSignType  = errors.New("sign type error")
	ErrorParameter = errors.New("JsonApiParameters() check error")
	ErrorToken     = errors.New("EditAddressParameters() token is nil")
)

/*RandomKind RandomKind */
type RandomKind int

/*random kinds */
const (
	RandomNum      RandomKind = iota // 纯数字
	RandomLower                      // 小写字母
	RandomUpper                      // 大写字母
	RandomLowerNum                   // 数字、小写字母
	RandomUpperNum                   // 数字、大写字母
	RandomAll                        // 数字、大小写字母
)

/*RandomString defines */
var (
	RandomString = map[RandomKind]string{
		RandomNum:      "0123456789",
		RandomLower:    "abcdefghijklmnopqrstuvwxyz",
		RandomUpper:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		RandomLowerNum: "0123456789abcdefghijklmnopqrstuvwxyz",
		RandomUpperNum: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		RandomAll:      "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
)

/*ParseNumber parse interface to number */
func ParseNumber(v interface{}) (float64, bool) {
	switch v0 := v.(type) {
	case float64:
		return v0, true
	case float32:
		return float64(v0), true
	}
	return 0, false
}

/*ParseInt parse interface to int64 */
func ParseInt(v interface{}) (int64, bool) {
	switch v0 := v.(type) {
	case int:
		return int64(v0), true
	case int32:
		return int64(v0), true
	case int64:
		return int64(v0), true
	case uint:
		return int64(v0), true
	case uint32:
		return int64(v0), true
	case uint64:
		return int64(v0), true
	case float64:
		return int64(v0), true
	case float32:
		return int64(v0), true
	default:
	}
	return 0, false
}

/*ParseString parse interface to string */
func ParseString(v interface{}) (string, bool) {
	switch v0 := v.(type) {
	case string:
		return v0, true
	case []byte:
		return string(v0), true
	case bytes.Buffer:
		return v0.String(), true
	default:
	}
	return "", false
}

/*MapToXML Convert MAP to XML */
func MapToXML(m Map) (string, error) {
	return mapToXML(m, false)
}

func mapToXML(m Map, needHeader bool) (string, error) {

	buff := bytes.NewBuffer([]byte(CustomHeader))
	if needHeader {
		buff.Write([]byte(xml.Header))
	}

	enc := xml.NewEncoder(buff)

	_ = enc.EncodeToken(xml.StartElement{Name: xml.Name{Local: "xml"}})
	for k, v := range m {
		if v0, b := v.(string); b {
			if _, err := strconv.ParseInt(v0, 10, 0); err != nil {
				_ = enc.EncodeElement(
					CDATA{Value: v0}, xml.StartElement{Name: xml.Name{Local: k}})
				continue
			}
		}
		_ = enc.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: k}})
	}
	_ = enc.EncodeToken(xml.EndElement{Name: xml.Name{Local: "xml"}})
	_ = enc.Flush()
	return buff.String(), nil
}

/*XMLToMap Convert XML to MAP */
func XMLToMap(contentXML []byte) Map {
	return xmlToMap(contentXML, false)
}

/*JSONToMap Convert JSON to MAP */
func JSONToMap(content []byte) Map {
	m := Map{}
	_ = json.Unmarshal(content, &m)
	return m
}

func xmlToMap(contentXML []byte, hasHeader bool) Map {
	m := make(Map)
	dec := xml.NewDecoder(bytes.NewReader(contentXML))
	ele, val := "", ""

	for t, err := dec.Token(); err == nil; t, err = dec.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			ele = token.Name.Local
			// fmt.Printf("This is the sta: %s\n", ele)
			if strings.ToLower(ele) == "xml" {
				// xmlFlag = true
				continue
			}

			// 处理元素结束（标签）
		case xml.EndElement:
			name := token.Name.Local
			// fmt.Printf("This is the end: %s\n", name)
			if strings.ToLower(name) == "xml" {
				break
			}
			if ele == name && ele != "" {
				m.Set(ele, val)
				ele = ""
				val = ""
			}
			// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			// content := string(token)
			// fmt.Printf("This is the content: %v\n", content)
			val = string(token)
			// 异常处理(Log输出）
		default:
			log.Println(token)
		}

	}

	return m
}

/*Time get time string */
func Time(t ...time.Time) string {
	if t == nil {
		return strconv.Itoa(time.Now().Nanosecond())
	}
	return strconv.Itoa(t[0].Nanosecond())
}

/*GenerateNonceStr GenerateNonceStr */
func GenerateNonceStr() string {
	return GenerateUUID()
}

/*GenerateUUID GenerateUUID */
func GenerateUUID() string {
	s := uuid.NewV1().String()
	s = strings.Replace(s, "-", "", -1)
	run := ([]rune)(s)[:32]
	return string(run)
}

/*In check v is in source */
func In(source []string, v string) bool {
	for _, v0 := range source {
		if v0 == v {
			return true
		}
	}
	return false
}

/*MapToString MapToString */
func MapToString(data Map, skip []string) string {
	var keys sort.StringSlice
	for k := range data {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	var sign []string

	for _, k := range keys {
		if In(skip, k) {
			continue
		}

		v := strings.TrimSpace(data.GetString(k))
		if len(v) > 0 {
			sign = append(sign, strings.Join([]string{k, v}, "="))
		}
	}
	log.Debug(strings.Join(sign, "&"))
	return strings.Join(sign, "&")
}

/*ToURLParams map to url params */
func ToURLParams(data Map, skip []string) string {
	keys := data.SortKeys()
	var sign []string

	for _, k := range keys {
		if In(skip, k) {
			continue
		}
		v := strings.TrimSpace(data.GetString(k))
		if len(v) > 0 {
			sign = append(sign, strings.Join([]string{k, v}, "="))
		}
	}
	return strings.Join(sign, "&")
}

// CurrentTimeStampMS get current time with millisecond
func CurrentTimeStampMS() int64 {
	return time.Now().UnixNano() / time.Millisecond.Nanoseconds()
}

// CurrentTimeStampNS get current time with nanoseconds
func CurrentTimeStampNS() int64 {
	return time.Now().UnixNano()
}

// CurrentTimeStamp get current time with unix
func CurrentTimeStamp() int64 {
	return time.Now().Unix()
}

// CurrentTimeStampString get current time to string
func CurrentTimeStampString() string {
	return strconv.FormatInt(CurrentTimeStamp(), 10)
}

// SHA1 transfer string to sha1
func SHA1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

func signatureSHA1(m Map) string {
	keys := m.SortKeys()
	var sign []string
	for _, k := range keys {
		if v := strings.TrimSpace(m.GetString(k)); v != "" {
			log.Debug(k, v)
			sign = append(sign, strings.Join([]string{k, v}, "="))
		} else if v, b := m.GetInt64(k); b {
			log.Debug(k, v)
			sign = append(sign, strings.Join([]string{k, strconv.FormatInt(v, 10)}, "="))
		}
	}
	sb := strings.Join(sign, "&")
	return SHA1(sb)
}

//GenerateRandomString2 随机字符串
func GenerateRandomString2(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

//GenerateRandomString 随机字符串
func GenerateRandomString(size int, kind ...RandomKind) string {
	bytes := RandomString[RandomAll]
	if kind != nil {
		if k, b := RandomString[kind[0]]; b == true {
			bytes = k
		}
	}
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
