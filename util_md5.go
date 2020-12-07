package base64Captcha

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateMD5ID(content string, password string, timeout int, idLenth int) (id string) {
	salt := int(int(time.Now().Unix()) / timeout)
	data := fmt.Sprintf(`%s_%d_%s`, content, salt, password)
	h := md5.New()
	h.Write([]byte(data))
	id = hex.EncodeToString(h.Sum(nil))
	id = id[0:idLenth]
	return
}
