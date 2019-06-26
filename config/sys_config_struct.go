package config

/***
 * 与yaml 对应的结构体
 */

 type DBConfig struct {
	 Host   string `profile:"host"`
	 Port   int `profile:"port"`
	 User   string `profile:"user"`
	 Pwd    string `profile:"pwd"`
	 DbName string `profile:"dbName"`
 }


type WebServer struct{
	Port int `profile:"port" profileDefault:"8188" `
	ContextPath string `profile:"context-path" `
}


type ApplicationConfig struct {
	Database DBConfig `profile:"database"`
	Server WebServer `profile:"server"`
}