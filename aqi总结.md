aqi启动项配置函数的作用是对AppConfig结构体的各种参数进行设置，然后根据设置来进行初始化，其中用到大量的viper库函数

设置环境变量的前缀，会匹配前缀为设置值的环境变量
viper.SetEnvPrefix("")

绑定环境变量，让viper优先读取环境变量
viper.AutomaticEnv()

设置配置文件名和类型，例如config.yaml。如开启了环境变量，环境变量中没有时会搜索配置文件。
viper.SetConfigName(acf.ConfigName)
viper.SetConfigType(acf.ConfigType)

设置配置文件路径，告诉viper在哪里寻找配置文件
viper.AddConfigPath(acf.ConfigPath)

读取配置文件后，保存至缓存，增加读取速度
viper.ReadInConfig()

返回viper.ReadInConfig()读取到的缓存值
viper.AllSettings()

1.LogConfig(configKeyPath string) Option
  应用日志文件配置路径

2.DataPath(path string) Option
  运行时数据存储基础路径
DataPath string//运行时数据存储基础路径

3.ConfigFile(file string) Option
  传入文件名设置配置文件
ConfigType string //配置文件类型
ConfigPath string //配置文件路径
ConfigName string //配置文件名称

4.Server(name ...string) Option
Servername             []string

5.HttpServer(name, portFindPath string) Option
//服务名称，support.Version
//当指定 HttpServerPortFindPath 时，在配置读取之后从配置路径获取http端口
Servername             []string
ServerPort             string
HttpServerPortFindPath string

6.WatchHandler(handler func()) Option