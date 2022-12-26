package models

import (
	//"01._Structures/internal/commons/phones"
	"errors"
	"fmt"
)

const MinPasswordLength = 7

type User struct {
	Identificator[UserId]
	Email    string
	Password string
	Age      uint8
	//Phone    phones.PhoneNumber
	Phone string
}

func NewUser(identificator Identificator[UserId], email string, password string, age uint8, phone string) *User {

	return &User{Identificator: identificator, Email: email, Password: password, Age: age, Phone: phone}
}

type UserId = uint64

func (user User) Identify() UserId {
	return user.Identificator.id()
}

func (user User) Reidentify(new UserId) {
	user.Identificator.setId(new)
}

func (user *User) SetPassword(newPassword string) {
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

func (user *User) SetPhone(newPhone string) (err error) {
	err = CheckPhone(newPhone)
	if err == nil {
		user.Phone = newPhone
		return
	}
	return fmt.Errorf("invalid new value \"%v\": %w", newPhone, err)
}

func (user *User) ToString() string {
	return fmt.Sprintf("Email - %v, password - %v, age - %v, phone - %v \n", user.Email, user.Password, user.Age, user.Phone)

}

//func (user *User) SetPhone(newPhone string) (err error) {
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
