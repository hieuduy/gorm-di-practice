package entity

type Address struct {
	// integer PrioritizedPrimaryField enables AutoIncrement by default
	ID 			int64  `gorm:"primary_key;autoIncrement:false"`
	Street      string `gorm:"not null" json:"street"`
	City        string `gorm:"not null" json:"city"`
	ZipCode     string `gorm:"index" json:"zip_code"`
	State       string `json:"state"`
}
