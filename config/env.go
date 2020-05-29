package config

var env = Env{
	ServerPort: "9999",     //服务端口
	Static:     "./static", //静态资源路径
	ImageAddr:  "http://134.175.138.178:9000/uploads/",
	Sdkappid:   1400362640,
	Key:        "34e1bb60f73163964f7a857101e718348faa10de88f3d6ce91e3119a4196b171",
	Expire:     2592000, //签名过期时间设置成一个月
}
