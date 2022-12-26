package phones

//
//import (
//	"errors"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//var countryCodes = []int32{
//	1, 7, 20, 27, 30, 31, 32, 33, 34, 36, 39, 40, 41, 43, 44, 45, 46, 47, 48, 49, 51, 52, 53, 54, 55, 56, 57, 58, 60, 61, 62, 63, 64, 65, 66, 81, 82, 84, 86, 90, 91, 92, 93, 94, 95, 98, 211, 212, 213, 216, 218, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 260, 261, 262, 263, 264, 265, 266, 267, 268, 269, 290, 291, 297, 298, 299, 350, 351, 352, 353, 354, 355, 356, 357, 358, 359, 370, 371, 372, 373, 374, 375, 376, 377, 378, 379, 380, 381, 382, 383, 385, 386, 387, 389, 420, 421, 423, 500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 590, 591, 592, 593, 594, 595, 596, 597, 598, 670, 672, 673, 674, 675, 676, 677, 678, 679, 680, 681, 682, 683, 685, 686, 687, 688, 689, 690, 691, 692, 800, 808, 850, 852, 853, 855, 856, 870, 878, 880, 881, 883, 886, 888, 960, 961, 962, 963, 964, 965, 966, 967, 968, 970, 971, 972, 973, 974, 975, 976, 977, 979, 992, 993, 994, 995, 996, 998, 1242, 1246, 1264, 1268, 1284, 1340, 1345, 1441, 1473, 1649, 1664, 1670, 1671, 1684, 1721, 1758, 1767, 1784, 1808, 1849, 1868, 1869, 1876, 1939, 2908, 3735, 4428, 4779, 5993, 5994, 5997, 5999, 6721, 6723, 8811, 8813, 8817, 8819, 25524, 35818, 37497, 88213, 88216, 90392, 99534, 262639, 441534, 447911, 447924, 6189162, 6189164,
//}
//
//const CodeMask = 10000000000
//
//func CheckCountryCode(code int32) bool {
//	i := sort.Search(len(countryCodes), func(i int) bool { return code <= countryCodes[i] })
//	return i < len(countryCodes) && code == countryCodes[i]
//
//}
//
//func GetCountryCodes() []int32 {
//	return countryCodes
//}
//
//type PhoneNumber interface {
//	ToSignedString() (phone string)
//	//ToFormattedString() (phone string)
//	Value() (number int64)
//}
//
//type phoneNumber struct {
//	value int64
//}
//
//func (phon *phoneNumber) Value() (number int64) {
//	return phon.value
//}
//
//func codeToNumber(code string) (number int, err error) {
//	if len(code) < 1 {
//		err = errors.New("malformed phone number format")
//	}
//	if code[0] == '+' {
//		code = code[1:]
//		number = 0
//	} else {
//		number = -1
//	}
//	numb, err := strconv.ParseInt(code, 10, 32)
//	number = number + int(numb)
//	return
//}
//
//func toNumber(phone string, checkStringPlusCode int8) (number int64, err error) {
//	var checkString, checkPlus, checkCountry bool
//	{
//		ch := checkStringPlusCode
//		checkString, checkPlus, checkCountry = ch&1 != 0, (ch&2)>>1 != 0, (ch&4)>>2 != 0
//	}
//	var code int32 = -1
//	if checkPlus && len(phone) != 0 && phone[0] == '+' {
//		code = 0
//		phone = phone[1:]
//	}
//	if checkString && len(phone) < 11 {
//		return 0, errors.New("invalid length of phone number")
//	}
//	number, err = strconv.ParseInt(phone[len(phone)-10:], 10, 64)
//	if err != nil {
//		return
//	}
//	{
//		tmp, terr := strconv.ParseInt(phone[0:len(phone)-10], 10, 32)
//		if terr != nil {
//			return 0, terr
//		}
//		code = code + int32(tmp)
//	}
//	if checkCountry && !CheckCountryCode(code) {
//		return 0, errors.New("such country code doesn't exists")
//	}
//
//	number = number%CodeMask + int64(code)*CodeMask
//	return
//}
//
//func RoughToNumber(phone string) (number int64, err error) {
//	var str = strings.Join(strings.FieldsFunc(phone, func(r rune) bool {
//		return r == ' ' || r == '-' || r == '(' || r == ')'
//	}), "")
//	return toNumber(str, 7)
//}
//
//func (phon *phoneNumber) ToSignedString() (phone string) {
//	return ToSignedString(phon.value)
//}
//
////func (phon *phoneNumber) ToFormattedString(format string, HasNext func() bool, Pass func() string) (phone string) {
////	return ToFormattedString(phon.value)
////}
////
////func ToFormattedString(value int64, format string, Pass func() (string, error)) string {
////
////}
//
//func ToSignedString(number int64) (phone string) {
//	return "+" + strconv.Itoa(int(number))
//}
//
//func New(phone string) (phoneN *phoneNumber, err error) {
//	phoneNum, er := RoughToNumber(phone)
//	if er != nil {
//		return nil, er
//	}
//	return &phoneNumber{value: phoneNum}, nil
//}
