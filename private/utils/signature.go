package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func GetSignatureNonce() string {
	u := uuid.NewV4()
	return u.String()
}

func GetSignString(Param map[string]string, mode string) string {
	SortedKeys := extractKeys(Param)
	SignString := ""
	for i, k := range SortedKeys {
		if mode == "full" {
			SignString += EncodeURIComponent(k) + "=" + EncodeURIComponent(Param[k])
			if i < len(SortedKeys)-1 {
				SignString += "&"
			}
		} else {
			SignString += k + Param[k]
		}

	}
	return SignString
}

func extractKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func EncodeURIComponent(str string) string {
	// r := url.QueryEscape(str)
	// r = strings.Replace(r, "+", "%20", -1)
	// return r
	// uri := tea.StringValue(str)
	str = url.QueryEscape(str)
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}
func HMACSHA1(keyStr, value string) string {
	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(value))
	//进行base64编码
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return res
}

func Base64Encode(value string) string {
	res := base64.StdEncoding.EncodeToString([]byte(value))
	return res
}

func HMACSHA256(keyStr, value string) string {
	mac := hmac.New(sha256.New, []byte(keyStr))
	mac.Write([]byte(value))
	//进行base64编码
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return res
}

func HMACSHA256Hex(keyStr, value string) string {
	hasher := hmac.New(sha256.New, []byte(keyStr))
	hasher.Write([]byte(value))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SHA256Hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
