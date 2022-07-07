package system

import (
	"errors"
	"fmt"
	"gms/global"
	system2 "gms/model/system"
	"gms/pkg"
	"gorm.io/gorm"
)

type UserService struct{}

// Login
//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser
func (userService *UserService) Login(u *system2.SysUser) (userInter *system2.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system2.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := pkg.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}

		var SysAuthorityMenus []system2.SysAuthorityMenu
		err = global.GVA_DB.Where("sys_authority_authority_id = ?", user.AuthorityId).Find(&SysAuthorityMenus).Error
		if err != nil {
			return
		}

		var MenuIds []string

		for i := range SysAuthorityMenus {
			MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
		}

		var am system2.SysBaseMenu
		ferr := global.GVA_DB.First(&am, "name = ? and id in (?)", user.Authority.DefaultRouter, MenuIds).Error
		if errors.Is(ferr, gorm.ErrRecordNotFound) {
			user.Authority.DefaultRouter = "404"
		}
	}

	return &user, err
}
