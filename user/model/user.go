package model

import (
	"database/sql/driver"
	"html"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RoleList []string

type User struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string   `gorm:"type:varchar(255)"`
	Password  string   `gorm:"type:varchar(255)"`
	Name      string   `gorm:"type:varchar(255)"`
	Roles     RoleList `gorm:"type:text"`
	Phone     string   `gorm:"type:varchar(255)"`
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "user"
}

func (r *RoleList) Scan(src any) error {
	// fmt.Println(src)
	// bytes, ok := src.([]byte)
	// if !ok {
	// 	return errors.New("src value cannot cast to []byte")
	// }
	*r = strings.Split(src.(string), ",")
	return nil
}
func (r RoleList) Value() (driver.Value, error) {
	if len(r) == 0 {
		return nil, nil
	}
	return strings.Join(r, ","), nil
}

// Generate encrypted password
func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Name = html.EscapeString(strings.TrimSpace(user.Name))
	return nil
}
