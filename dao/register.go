package dao

type RegisterModel struct {
	Groupid       int    `json:"groupid" column:"groupid" form:"groupid"`
	Username      string `json:"username" column:"username" form:"username" `
	Password      string `json:"password" column:"password" form:"password"`
	Joindate      int64  `json:"joindate" column:"joindate" form:"joindate"`
	Joinip        string `json:"joinip" column:"joinip" form:"joinip"`
	Register_type int    `json:"register_type" column:"register_type" form:"register_type"`
	Remark        string `json:"remark" column:"remark" form:"remark"`
	Openid        string `json:"openid" column:"openid" form:"openid"`
	Status        int    `json:"status" column:"status" form:"status"`
	Usersig       string `json:"usersig" column:"usersig" form:"usersig"`
	Expire        int64  `json:"expire" column:"expire" form:"expire"`
}
