package billing

import (
	"strings"

	"github.com/gavintony1990/Lexvia/common"
	"github.com/gavintony1990/Lexvia/setting/system_setting"
)

func PaymentReturnPath(suffix string) string {
	base := strings.TrimRight(system_setting.ServerAddress, "/")
	return base + common.ThemeAwarePath(suffix)
}
