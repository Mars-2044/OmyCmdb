package logic

import (
	"OmyBlog/dao/mysql"
	"OmyBlog/models"
	"OmyBlog/pkg/snowflake"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err = mysql.QueryUserByUsername(p.Username); err != nil {
		// 数据库查询错误
		return err
	}

	// 2.生成UID
	userID := snowflake.GetID()

	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 3.保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.LoginSignUp) (err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
