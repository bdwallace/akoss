package common

import (
	"archive/zip"
	"bytes"
	"crypto/des"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//判断一个数据是否为空，支持int, float, string, slice, array, map的判断
func Empty(value interface{}) bool {
	if value == nil {
		return true
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		if reflect.ValueOf(value).Len() == 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

//判断某一个值是否在列表(支持 slice, array, map)中
func InList(needle interface{}, haystack interface{}) bool {
	//interface{}和interface{}可以进行比较，但是interface{}不可进行遍历
	hayValue := reflect.ValueOf(haystack)
	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice, reflect.Array:
		//slice, array类型
		for i := 0; i < hayValue.Len(); i++ {
			if hayValue.Index(i).Interface() == needle {
				return true
			}
		}
	case reflect.Map:
		//map类型
		var keys []reflect.Value = hayValue.MapKeys()
		for i := 0; i < len(keys); i++ {
			if hayValue.MapIndex(keys[i]).Interface() == needle {
				return true
			}
		}
	default:
		return false
	}
	return false
}

//返回某一个值是否在列表位置(支持 slice, array, map) -1为不再列表中
func InListIndex(needle interface{}, haystack interface{}) int {
	//interface{}和interface{}可以进行比较，但是interface{}不可进行遍历
	hayValue := reflect.ValueOf(haystack)
	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice, reflect.Array:
		//slice, array类型
		for i := 0; i < hayValue.Len(); i++ {
			if hayValue.Index(i).Interface() == needle {
				return i
			}
		}
	case reflect.Map:
		//map类型
		var keys []reflect.Value = hayValue.MapKeys()
		for i := 0; i < len(keys); i++ {
			if hayValue.MapIndex(keys[i]).Interface() == needle {
				return i
			}
		}
	default:
		return -1
	}
	return -1
}

//string转int
func StrToInt(str string) int {
	intval, _ := strconv.Atoi(str)
	return intval
}

//string转f
func StrToFloat(str string) float64 {
	floatval, _ := strconv.ParseFloat(str, 64)
	return floatval
}

//string转int64
func StrToInt64(str string) int64 {
	int64val, _ := strconv.ParseInt(str, 10, 64)
	return int64val
}

//int转string
func IntToStr(i int) string {
	intval := strconv.Itoa(i)
	return intval
}

//浮点数四舍五入，并取前几位
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

//通过interface{}获取字符串
func GetString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

func AddStringListSingle(s string) (res string) {
	for _, v := range strings.Split(s, ",") {
		if res == "" {
			res = fmt.Sprintf("'%s'", v)
		} else {
			res = fmt.Sprintf("%s, '%s'", res, v)
		}
	}

	return
}

//通过interface{}获取数值型数据
//此获取比较灵活，转换规则如下
//1、如果接收数据为浮点string，则返回浮点数的整数部分，如果是整型string，则返回整数，如果是纯字符串，则返回0
//2、如果接收数据是float型，则返回float的整数部分
//3、如果接收数据是int, int32, int64型，则返回int
func GetInt(val interface{}) int {
	switch v := val.(type) {
	case int:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case string:
		n, err := strconv.Atoi(v)
		if err != nil {
			fval, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return 0
			}
			return int(fval)
		}
		return int(n)
	case float32:
		return int(v)
	case float64:
		return int(v)
	default:
		return 0
	}
}

//通过interface{}获取小数型数据
//此获取比较灵活，转换规则如下
//1、如果接收数据为浮点string，则将字符串转换为浮点数
//2、如果接收数据是float型，则返回float数据
//3、如果接收数据是int, int32, int64型，则转义成float类型
//4、返回的数据结果统一为float64
func GetFloat(val interface{}) float64 {
	switch v := val.(type) {
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case int32:
		return float64(v)
	case float64:
		return v
	case float32:
		return float64(v)
	case string:
		result, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0
		}
		return result
	}
	return 0
}

/**
 * 根据path读取文件中的内容，返回字符串
 * 建议使用绝对路径，例如："./schema/search/appoint.json"
 */
func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func ReadJson(path string) Info {
	jsonStr := ReadFile(path)
	ret := Info{}
	err := json.Unmarshal([]byte(jsonStr), &ret)
	if err != nil {
		panic("文件[" + path + "]的内容不是json格式")
	}
	return ret
}

/**
  判断内网IP
  A  10.0.0.0/8：10.0.0.0～10.255.255.255
  B  172.16.0.0/12：172.16.0.0～172.31.255.255
  C  192.168.0.0/16：192.168.0.0～192.168.255.255
**/
func CheckInternalIp(ip string) bool {
	if ip == "127.0.0.1" {
		return true
	}
	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		return false
	}
	a_from_ip := net.ParseIP("10.0.0.0")
	a_to_ip := net.ParseIP("10.255.255.255")
	b_from_ip := net.ParseIP("172.16.0.0")
	b_to_ip := net.ParseIP("172.31.255.255")
	c_from_ip := net.ParseIP("192.168.0.0")
	c_to_ip := net.ParseIP("192.168.255.255")
	if bytes.Compare(trial, a_from_ip) >= 0 && bytes.Compare(trial, a_to_ip) <= 0 {
		return true
	}
	if bytes.Compare(trial, b_from_ip) >= 0 && bytes.Compare(trial, b_to_ip) <= 0 {
		return true
	}
	if bytes.Compare(trial, c_from_ip) >= 0 && bytes.Compare(trial, c_to_ip) <= 0 {
		return true
	}
	return false
}
func Md5String(str string) string {

	h := md5.New()
	h.Write([]byte(str))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

//获取当前日期为当月第几周
func CountWeek(TimeFormat string) int {
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", TimeFormat, loc)
	month := t.Month()
	year := t.Year()
	days := 0
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	week := 1
	for i := 1; i <= days; i++ {
		dayString := GetString(i)
		if i < 10 {
			dayString = "0" + dayString
		}
		dateString := strings.Split(TimeFormat, "-")[0] + "-" + strings.Split(TimeFormat, "-")[1] + "-" + dayString + " 18:30:50"
		t1, _ := time.ParseInLocation("2006-01-02 15:04:05", dateString, loc)
		if t.YearDay() > t1.YearDay() {
			if t1.Weekday().String() == "Sunday" {
				week++
			}
		}

	}

	return week
}
func GetWeekday(TimeFormat string) string {
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", TimeFormat, loc)
	if t.Weekday().String() == "Monday" {
		return "星期一"
	}
	if t.Weekday().String() == "Tuesday" {
		return "星期二"
	}
	if t.Weekday().String() == "Wednesday" {
		return "星期三"
	}
	if t.Weekday().String() == "Thursday" {
		return "星期四"
	}
	if t.Weekday().String() == "Friday" {
		return "星期五"
	}
	if t.Weekday().String() == "Saturday" {
		return "星期六"
	}
	if t.Weekday().String() == "Sunday" {
		return "星期日"
	}
	return ""

}
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

/**
 * 新建目录
 */
func Mkdir(dir string) (err error) {
	err = os.MkdirAll(dir, os.ModePerm)
	return
}

/**
 * 删除目录,再新建目录
 */
func RmMkdir(dir string) (err error) {
	if err = os.RemoveAll(dir); err == nil {
		err = os.MkdirAll(dir, os.ModePerm)
	}
	return
}

// func DesECBEncrypt(data, key []byte) ([]byte, error) {
func DesECBEncrypt(text string, keys string) (string, error) {
	key := []byte(keys[:8])
	data := []byte(text)
	//NewCipher创建一个新的加密块
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println("-----初始化加密块，失败！-----")
		return "", err
	}
	fmt.Println("-----初始化加密块，成功！-----")
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		err := fmt.Errorf("need a multiple of the blocksize")
		return "", err
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Encrypt加密第一个块，将其结果保存到dst
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	result := base64.StdEncoding.EncodeToString(out)
	return result, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// srcFile could be a single file or a directory
func Zip(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		// header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/**
 * 覆盖文件
 */
func FileWriter(file string, wireteString string) (err error) {
	// 只能写不能追加
	if !checkFileIsExist(file) {
		_, err = os.Create(file)
		if err != nil {
			return
		}
	}
	var d = []byte(wireteString)
	err = ioutil.WriteFile(file, d, 0644)

	return
}

/**
 * 返回home目录
 */
func HomeDir() (home string) {
	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return
	} else if runtime.GOOS == "linux" {
		home = os.Getenv("HOME")
		return
	}

	home = os.Getenv("HOME")
	return
}

/*
	使用原生net/http方式请求URL

*/
func HttpGet(url string, headers map[string]string) (*http.Response, error) {
	tr := &http.Transport{ //解决x509: certificate signed by unknown authority
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr, //解决x509: certificate signed by unknown authority
	}
	req, err := http.NewRequest("GET", url, nil)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	client.CloseIdleConnections()
	return resp, nil
}

/*
	使用原生net/http方式请求URL

*/
func HttpPost(url string, headers map[string]string, data []byte) (*http.Response, error) {
	tr := &http.Transport{ //解决x509: certificate signed by unknown authority
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr, //解决x509: certificate signed by unknown authority
	}
	body := bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, body)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	client.CloseIdleConnections()
	return resp, nil
}

/*
	休息，sleep
*/
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}
