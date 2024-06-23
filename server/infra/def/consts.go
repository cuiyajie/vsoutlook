package def

const (
	NotDeleted    uint8 = 0
	SelfDeleted   uint8 = 1
	ParentDeleted uint8 = 2
	OwnerDeleted  uint8 = 3 // file deleted
)

const (
	Role_Admin  = 1
	Role_Normal = 2
)

const (
	LSize_4K = 1
	LSize_HD = 2
)

const SettingKey_Mtv = "mv_template_list"
const SettingKey_EndSwtApi = "endswt_api"

var LSizeStr = map[uint8]string{
	LSize_4K: "4K",
	LSize_HD: "HD",
}

const DefLayoutPath = "/opt/vsomediasoftware/layout/"
