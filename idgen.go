package utils

import (
	"crypto/md5"
	"encoding/base64"
	"time"
	"encoding/binary"
	"sort"
	"strings"
	"crypto/sha1"
	"fmt"
)

//  generate url safe, pseudo unique id based on given input
func GenerateId(input string, numbytes int) string{
	b := md5.Sum([]byte(input))
	return base64.URLEncoding.EncodeToString(b[:numbytes])
}

//  generate url safe id based on current unix nano second
func GenerateTimeBasedId() (string,int64) {
	t := time.Now().UnixNano()

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b,uint64(t))

	return base64.URLEncoding.EncodeToString(b),t
}


//  recover the unit time from string id
func GetUnixTimeFromTimeBasedId(id string) (int64, error) {
	b, err := base64.URLEncoding.DecodeString(id)

	if err != nil {
		return 0, err
	}
	ut := binary.LittleEndian.Uint64(b)

	return int64(ut)/10e8, nil
}

func GenerateSignature(strs ...string) string {
	sort.Strings([]string(strs))
	strToDigest := strings.Join([]string(strs), "")

	digest := sha1.Sum([]byte(strToDigest))
	return fmt.Sprintf("%x", digest)
}
