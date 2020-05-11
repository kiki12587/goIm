package config

type Env struct {
	Static     string // 静态文件相对路径
	ServerPort string // web服务端口号
	ImageAddr  string
}

func GetEnv() *Env {
	return &env
}
