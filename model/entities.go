package model

type Entity interface {
}

type Member struct {
	ID        string `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Active    bool   `gorm:"not null"`
}

type Season struct {
	ID      string `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Created string `gorm:"not null"`
}

type SeasonGame struct {
	ID             string `gorm:"primaryKey"`
	Opponent       string `gorm:"not null"`
	Home           bool   `gorm:"not null"`
	Mode           string `gorm:"not null"`
	Date           string `gorm:"not null"`
	SeasonID       string `gorm:"not null"`
	Season         Season
	PresentMembers []Member `gorm:"many2many:present_members;"`
}

type SeasonPresence struct {
	Name       string
	HomeGames  int
	TotalGames int
}

type MemberPresence struct {
	Season     string
	LastName   string
	FirstName  string
	HomeGames  int
	TotalGames int
}
