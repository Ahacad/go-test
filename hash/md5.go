package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func StringMd5(s string) string {
    md5 := md5.New()
    md5.Write([]byte(s))
    return hex.EncodeToString(md5.Sum(nil))
}
