package config

import (
	"fmt"
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

var (
	AppVersion             = "v4.5.1"
	AppPort                = "3000"
	AppDebug               = false
	AppOs                  = fmt.Sprintf("AldinoKemal")
	AppPlatform            = waProto.DeviceProps_PlatformType(1)
	AppBasicAuthCredential string

	PathQrCode    = "statics/qrcode"
	PathSendItems = "statics/senditems"
	PathMedia     = "statics/media"
	PathStorages  = "storages"

	DBName = "whatsapp.db"

	WhatsappLogLevel            = "ERROR"
	WhatsappAutoReplyMessage    string
	WhatsappAutoReplyWebhook    string
	WhatsappSettingMaxFileSize  int64 = 50000000  // 50MB
	WhatsappSettingMaxVideoSize int64 = 100000000 // 100MB
)
