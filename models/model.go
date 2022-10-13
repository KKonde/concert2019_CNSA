package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

const (
	// SQLConnectionString : MySQL Connection String
	SQLConnectionString = "dbuser:cnsaconcertdb@tcp(45.77.134.35:3306)/cnsaconcert2019?charset=utf8&parseTime=True"
	// SALT : SALT
	SALT = "5SoWWSqK0Kq$}_jqN9&AaMGFkJ_[(}O}O.k<Mq!N9&Aa^WN<d@VbKd~.ZFJKNJ%^fZWz)IC^2.+/U`YiJadcD]CMG@M&KNJ%^fZ}_jqF2a15*@e4}S8UdiHjw%65SkJ_[(}O}Aaz!gMN1QA@nwRIG/]1+`u?.O14N1QA@6kK:EnwRIGHjw,I2YMnwRIN1QA@G/]1+`S`>q/921*%65SoWWSqK0Kq$O.Lt*#E4<>g%?{#ai2M.^YqVYuy<^WN<d@&"
)

// Database Connection
var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", SQLConnectionString)
	if err != nil {
		fmt.Println("DBDBDBDBDBDBDBDB")
		panic(err)
	}
}

// User models
type User struct {
	StudentNumber string `gorm:"type:VARCHAR(10); primary_key" json:"studentNumber"`
	Password      string `gorm:"type:VARCHAR(100)" json:"password"`
	Name          string `gorm:"type:VARCHAR(30)" json:"name"`
}

// TableName of User
func (c *User) TableName() string {
	return "users"
}

// Apply models
type Apply struct {
	StudentNumber string `gorm:"type:VARCHAR(10); primary_key" json:"studentNumber"`
	Name          string `gorm:"type:VARCHAR(30)" json:"name"`
	Class         string `gorm:"type:VARCHAR(10); unique_key" json:"class"`
	Row           string `gorm:"type:VARCHAR(10); unique_key" json:"row"`
	Col           string `gorm:"type:VARCHAR(10); unique_key" json:"col"`
}

// TableName of Apply
func (c *Apply) TableName() string {
	return "applys"
}

// Nope models
type Nope struct {
	StudentNumber string `gorm:"type:VARCHAR(10); primary_key" json:"studentNumber"`
	Name          string `gorm:"type:VARCHAR(30)" json:"name"`
}

// TableName of Nope
func (c *Nope) TableName() string {
	return "nopes"
}

// Login 학생 아이디 인증(SALT)
func Login(studentNumber string, password string) (bool, string, string) {
	user := User{}
	err := db.Table("users").Where("student_number = ?", studentNumber).First(&user).Error
	if err != nil {
		return false, "", "studentNumber"
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+SALT))
	if err != nil {
		return false, "", "password"
	}

	return true, user.Name, "success"
}

// ChangePasswordAdmin 관리자용 비번 변경
func ChangePasswordAdmin(studentNumber string, newPassword string) error {
	user := User{}
	db.Table("users").Where("student_number = ?", studentNumber).First(&user)

	bytes, _ := bcrypt.GenerateFromPassword([]byte(newPassword+SALT), bcrypt.DefaultCost)
	user.Password = string(bytes)

	return db.Save(&user).Error
}

// ChangePassword 비밀번호 변경
func ChangePassword(studentNumber string, newPassword string) error {
	user := User{}
	db.Table("users").Where("student_number = ?", studentNumber).First(&user)

	bytes, _ := bcrypt.GenerateFromPassword([]byte(newPassword+SALT), bcrypt.DefaultCost)
	user.Password = string(bytes)

	return db.Save(&user).Error
}

// GetApplysByClass 신청좌석 가져오기 by Class
func GetApplysByClass(class string) []Apply {
	applys := []Apply{}
	db.Table("applys").Where("class = ?", class).Find(&applys)
	return applys
}

// GetApplyMine 내 좌석 가져오기
func GetApplyMine(studentNumber string) (string, string, string) {
	apply := Apply{}
	err := db.Table("applys").Where("student_number = ?", studentNumber).First(&apply).Error
	if err == nil {
		return apply.Class, apply.Row, apply.Col
	}
	return "0", "0", "0"
}

// AddApply 좌석 신청하기
func AddApply(studentNumber string, name string, class string, row string, col string) error {
	apply := Apply{}
	err := db.Table("applys").Where("student_number = ?", studentNumber).First(&apply).Error
	if err == nil {
		apply.Row = row
		apply.Col = col
		return db.Save(&apply).Error
	}

	return db.Create(&Apply{
		StudentNumber: studentNumber,
		Name:          name,
		Class:         class,
		Row:           row,
		Col:           col,
	}).Error
}

// Insert 학생 db 추가
func Insert(studentNumber string, name string) error {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(studentNumber+SALT), bcrypt.DefaultCost)

	return db.Create(&User{
		StudentNumber: studentNumber,
		Password:      string(bytes),
		Name:          name,
	}).Error
}

// MakeNope 미신청자
func MakeNope(studentNumber string, name string) error {
	apply := Apply{}
	err := db.Table("applys").Where("student_number = ?", studentNumber).First(&apply).Error

	if err == nil {
		return db.Save(&apply).Error
	}

	return db.Create(&Nope{
		StudentNumber: studentNumber,
		Name:          name,
	}).Error
}
