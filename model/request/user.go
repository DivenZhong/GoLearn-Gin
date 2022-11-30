/**
 * @Description TODO
 **/
package request

// 登录参数
type LoginParam struct {
	Phone    string `json:"phone"`    //手机
	Password string `json:"password"` //密码
}

// 注册参数
type RegisterParam struct {
	NickName string `json:"nickName"` //昵称
	Birthday string `json:"birthday"` //生日
	Phone    string `json:"phone"`    //手机
	Password string `json:"password"` //密码
	Address  string `json:"address"`  //地址
}
