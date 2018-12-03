package models

import (
	u "github.com/dlockamy/goContacts/utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Contact struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	UserId uint `json:"usesr_id"`
}

func (contact *Contact) Validate() (map[string] interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Contact number should be on the payload"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	return u.Message(true, "success"), true
}

func (contact *Contact) Create() (map[string] interface{}) {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func Getcontact(id uint) (*Contact) {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}

	return contact
}

func Getcontacts(user uint) ([]*Contact) {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}