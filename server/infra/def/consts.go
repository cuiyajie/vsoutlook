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
