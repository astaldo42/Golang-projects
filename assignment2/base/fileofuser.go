package base

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users []User

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func Registration(username, password string) {
	user := User{Username: username}
	user.SetPassword(password)
	Users = append(Users, user)
}

//
//func Authorization(username, password string) bool {
//	for _, user := range Users {
//		if user.Username == username && user.Password == password {
//			return true
//		}
//	}
//	return false
//}
//
//func encryptPassword(password string) string {
//	hash := sha256.Sum256([]byte(password))
//	return fmt.Sprintf("%x", hash)
//}
