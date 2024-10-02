package models

// ItemChecklist represents the checklist structure for items.
type ItemChecklist struct {
	ItemID        uint   `json:"itemID"`
	StorageKindID uint   `json:"storageKindID"`
	Category      string `json:"category"`
	DisplayName   string `json:"displayName"`
	Length        uint   `json:"length"`
	Width         uint   `json:"width"`
	Height        uint   `json:"height"`
	Area          uint   `json:"area"`
	Volume        uint   `json:"volume"`
}
