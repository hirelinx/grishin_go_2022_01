package models

import (
	//"01._Structures/internal/commons/phones"
	"errors"
	"fmt"
)

const MinPasswordLength = 7

type User struct {
	Email    string
	Password string
	Age      int
	//Phone    phones.PhoneNumber
	Phone string
}

func (user *User) ChangePassword(newPassword string) {
	if len(newPassword) >= MinPasswordLength {
		user.Password = newPassword // (*user).Password
	}
}

func CheckPhone(phone string) (err error) {
	if len(phone) != 12 {
		err = errors.New(fmt.Sprintf("wrong length of phone (%d)", len(phone)))
		return
	}
	if phone[0:2] == "+7" {
		return
	}
	err = errors.New("malformed phone")
	return
}

func (user *User) ChangePhone(newPhone string) (err error) {
	err = CheckPhone(newPhone)
	if err == nil {
		user.Phone = newPhone
		return
	}
	return fmt.Errorf("invalid new value \"%v\": %w", newPhone, err)
}

func (user *User) PrintInformation() {
	fmt.Printf("Email - %v, password - %v, age - %v, phone - %v \n", user.Email, user.Password, user.Age, user.Phone)

}

//func (user *User) ChangePhone(newPhone string) (err error) {
//	phone, err := phones.New(newPhone)
//	if phone == nil {
//		return
//	}
//	user.Phone = phone
//	return
//}

//func (user *User) PrintInformation() {
//	fmt.Printf("Email - %v, password - %v, age - %v, phone - +%d \n", user.Email, user.Password, user.Age, user.Phone.Value())
//
//}
