package models

import (
	"github.com/provider-go/user/global"
	"time"
)

type UserInfo struct {
	DID        string    `json:"did" gorm:"column:did;type:varchar(100);not null;default:'';primary_key;comment:'主键'"`
	Username   string    `json:"username" gorm:"column:username;type:varchar(20);not null;default:'';comment:用户名"` // 用户名
	Password   string    `json:"password" gorm:"column:password;type:varchar(50);not null;default:'';comment:密码"`  // 密码
	Nickname   string    `json:"nickname" gorm:"column:nickname;type:varchar(20);not null;default:'';comment:昵称"`  // 昵称
	Sex        int       `json:"sex" gorm:"column:sex;type:tinyint(1);not null;default:0;comment:性别：0(男)1(女)"`     // 性别：0：男；1：女
	Avatar     string    `json:"avatar" gorm:"column:avatar;type:varchar(200);not null;default:'';comment:头像地址"`   // 头像地址
	Phone      string    `json:"phone" gorm:"column:phone;type:varchar(20);not null;default:'';comment:电话"`        // 电话
	Email      string    `json:"email" gorm:"column:email;type:varchar(50);not null;default:'';comment:邮箱"`        // 邮箱
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`                                   // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`                                   // 更新时间
}

func CreateUserInfo(did, username, nickname string, sex int, avatar, phone, email string) error {
	return global.DB.Table("user_infos").
		Create(&UserInfo{DID: did, Username: username, Nickname: nickname, Sex: sex, Avatar: avatar, Phone: phone, Email: email}).Error
}

func UpdateUserInfo(did, username, nickname string, sex int, avatar, phone, email string) error {
	return global.DB.Table("user_infos").Where("did = ?", did).Updates(map[string]interface{}{
		"username": username,
		"nickname": nickname,
		"sex":      sex,
		"avatar":   avatar,
		"phone":    phone,
		"email":    email,
	}).Error
}

func DeleteUserInfo(did string) error {
	return global.DB.Table("user_infos").Where("did = ?", did).Delete(&UserInfo{}).Error
}

func ListUserInfo(pageSize, pageNum int) ([]*UserInfo, int64, error) {
	var rows []*UserInfo
	//计算列表数量
	var count int64
	global.DB.Table("user_infos").Count(&count)

	if err := global.DB.Table("user_infos").Order("create_time desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func ViewUserInfo(did string) (*UserInfo, error) {
	row := new(UserInfo)
	if err := global.DB.Table("user_infos").Where("did = ?", did).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}
