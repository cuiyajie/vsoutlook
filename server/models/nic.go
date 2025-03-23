package models

type Uint32Slice []uint32
type MapUint32Slice map[string]Uint32Slice
type StringSlice []string
type MapStringSlice map[string]StringSlice

type Nic struct {
	ModelId
	NicNameMain   string         `gorm:"type:varchar"`
	NicNameBackup string         `gorm:"type:varchar"`
	ReceiveMain   bool           `gorm:"type:boolean"`
	ReceiveBackup bool           `gorm:"type:boolean"`
	SendMain      bool           `gorm:"type:boolean"`
	SendBackup    bool           `gorm:"type:boolean"`
	CoreList      string         `gorm:"type:varchar"`
	AllocatedCore MapUint32Slice `gorm:"type:jsonb"`
	DMAList       string         `gorm:"type:varchar"`
	AllocatedDMA  MapStringSlice `gorm:"type:jsonb"`
	VFCount       uint32         `gorm:"type:integer;default:4"`
	NodeID        string         `gorm:"type:varchar"`
	Deleted       uint8          `gorm:"default:0"`
}

type NicInfo struct {
	NicNameMain   string `json:"nicNameMain"`
	NicNameBackup string `json:"nicNameBackup"`
	ReceiveMain   bool   `json:"receiveMain"`
	ReceiveBackup bool   `json:"receiveBackup"`
	SendMain      bool   `json:"sendMain"`
	SendBackup    bool   `json:"sendBackup"`
	CoreList      string `json:"coreList"`
	DMAList       string `json:"dmaList"`
	VFCount       uint32 `json:"vfCount"`
}

type NicAsBasic struct {
	ID string `json:"id"`
	NicInfo
	AllocatedCore MapUint32Slice `json:"allocatedCore"`
	AllocatedDMA  MapStringSlice `json:"allocatedDMA"`
	NodeID        string         `json:"nodeId"`
}

func (nic *Nic) AsBasic() NicAsBasic {
	var result NicAsBasic
	result.ID = nic.ID
	result.NicNameMain = nic.NicNameMain
	result.NicNameBackup = nic.NicNameBackup
	result.ReceiveMain = nic.ReceiveMain
	result.ReceiveBackup = nic.ReceiveBackup
	result.SendMain = nic.SendMain
	result.SendBackup = nic.SendBackup
	result.CoreList = nic.CoreList
	result.AllocatedCore = nic.AllocatedCore
	result.DMAList = nic.DMAList
	result.AllocatedDMA = nic.AllocatedDMA
	result.VFCount = nic.VFCount
	result.NodeID = nic.NodeID
	return result
}

func (nic Nic) IsDeleted() bool {
	return nic.Deleted > 0
}
