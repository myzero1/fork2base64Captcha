package base64Captcha

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func generateMD5ID(content string, timeout int) (id string) {
	salt := int(int(time.Now().Unix()) / timeout)
	data := fmt.Sprintf(`%s_%d`, content, salt)
	h := md5.New()
	h.Write([]byte(data))
	id = hex.EncodeToString(h.Sum(nil))
	return
}
