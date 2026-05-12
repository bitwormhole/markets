package normalizers

import "strings"

func PurifyCode(code string) string {

	i := strings.IndexByte(code, '@')
	if i < 0 {
		return code
	}
	str := code[0:i]
	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)
	return str
}
