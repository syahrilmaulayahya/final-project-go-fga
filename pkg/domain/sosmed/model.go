package sosmed

type Sosmed struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"not null"`
	Url    string `gorm:"not null"`
	UserID string `gorm:"not null"`
}
