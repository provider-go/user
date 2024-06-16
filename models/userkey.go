package models

import (
	"github.com/provider-go/user/global"
	"time"
)

type UserKey struct {
	DID        string    `json:"did" gorm:"column:did;type:varchar(100);not null;default:'';comment:'主键'"`
	PubKey     string    `json:"pubkey" gorm:"column:pubkey;type:varchar(100);not null;default:'';comment:用户公钥"`         // 用户公钥
	Status     int       `json:"status" gorm:"column:status;type:tinyint(1);not null;default:0;comment:账号状态:0(正常)1(禁用)"` // 账号状态:0(正常)1(禁用)
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`                                         // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`                                         // 更新时间
}

func CreateUserKey(did, pubkey string) error {
	return global.DB.Table("user_keys").Create(&UserKey{DID: did, PubKey: pubkey}).Error
}

func UpdateUserKey(did, pubkey string) error {
	return global.DB.Table("user_keys").Where("did = ?", did).Updates(map[string]interface{}{
		"pubkey": pubkey,
	}).Error
}

func DeleteUserKey(did string) error {
	return global.DB.Table("user_keys").Where("did = ?", did).Delete(&UserKey{}).Error
}

func ListUserKey(pageSize, pageNum int) ([]*UserKey, int64, error) {
	var rows []*UserKey
	//计算列表数量
	var count int64
	global.DB.Table("user_keys").Count(&count)

	if err := global.DB.Table("user_keys").Order("create_time desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func ViewUserKey(did string) (*UserKey, error) {
	row := new(UserKey)
	if err := global.DB.Table("user_keys").Where("did = ?", did).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}
