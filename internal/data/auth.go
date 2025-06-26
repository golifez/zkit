package data

import (
	"context"

	"github.com/golifez/zkit/internal/auth"
	"github.com/golifez/zkit/internal/biz"
	"github.com/golifez/zkit/internal/conf"
	"github.com/golifez/zkit/internal/data/ent/user"
	"github.com/golifez/zkit/internal/domain"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type AutherRepo struct {
	data *Data
	log  *log.Helper
	c    *conf.Config
}

// NewGreeterRepo .
func NewAutherRepo(data *Data, c *conf.Config, logger log.Logger) biz.AutherRepo {
	return &AutherRepo{
		data: data,
		log:  log.NewHelper(logger),
		c:    c,
	}
}

func (a *AutherRepo) Login(ctx context.Context, g *domain.Auther) (rg *domain.Auther, err error) {
	// 从数据库查出来
	user, err := a.data.db.User.Query().Where(user.UserNameEQ(g.Username)).First(ctx)
	if err != nil {
		a.log.Errorf("查询用户失败: %v", err)
		return nil, err
	}

	a.log.Infof("查询到的用户: username=%s, password=%s", user.UserName, user.Password)

	// 解密数据库里的密码
	decryptedPassword, err := auth.Decrypt(user.Password, a.c.Secretkey.Passwdkey)
	if err != nil {
		a.log.Errorf("解密失败: %v", err)
		return nil, err
	}

	a.log.Infof("传入的密码: '%s' (长度: %d)", g.Password, len(g.Password))
	a.log.Infof("解密后的密码: '%s' (长度: %d)", decryptedPassword, len(decryptedPassword))

	// 详细比较
	if g.Password == decryptedPassword {
		a.log.Info("✅ 密码匹配成功")
		g.Uid = user.UID
		return g, nil
	} else {
		a.log.Errorf("❌ 密码不匹配")
		a.log.Errorf("传入密码字节: %v", []byte(g.Password))
		a.log.Errorf("解密密码字节: %v", []byte(decryptedPassword))
		return nil, errors.Unauthorized("password_error", "密码错误")
	}
}

// Register 注册
func (a *AutherRepo) Register(ctx context.Context, g *domain.Auther) error {
	// 从数据库查出来
	num, err := a.data.db.User.Query().Where(user.UserNameEQ(g.Username)).Count(ctx)
	if err != nil {
		return err
	}
	if num > 0 {
		return errors.Unauthorized("username_exist", "用户名已存在")
	}
	// 加密密码
	encpasswd := auth.Encryption(g.Password, a.c.Secretkey.Passwdkey)
	//获取UId
	uid := auth.GenUid()
	// 写入数据库
	_, err = a.data.db.User.Create().SetUID(uid).SetUserName(g.Username).SetPassword(encpasswd).SetNickName(g.NickName).Save(ctx)
	if err != nil {
		return err
	}
	return nil

}
