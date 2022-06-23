package repository

import "log"

type UserInfo struct {
	ID       int64
	UserName string
}

type userInfoRepository struct{}

func GetTodos() (userInfos []UserInfo, err error) {
	userInfos = []UserInfo{}
	rows, err := Db.Query("SELECT * FROM UserInfo")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		userInfo := UserInfo{}
		err = rows.Scan(&userInfo.ID, &userInfo.UserName)
		if err != nil {
			log.Print(err)
			return
		}
		userInfos = append(userInfos, userInfo)
	}

	return
}
