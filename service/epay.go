package service

import (
	"github.com/gavintony1990/Lexvia/setting/operation_setting"
	"github.com/gavintony1990/Lexvia/setting/system_setting"
)

func GetCallbackAddress() string {
	if operation_setting.CustomCallbackAddress == "" {
		return system_setting.ServerAddress
	}
	return operation_setting.CustomCallbackAddress
}
