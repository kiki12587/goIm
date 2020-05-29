package config

type Env struct {
	Static     string // 静态文件相对路径
	ServerPort string // web服务端口号
	ImageAddr  string
	Sdkappid   int    //腾讯IM Appid
	Key        string //腾讯IM 密钥
	Expire     int    //过期时间
}

func GetEnv() *Env {
	return &env
}
