DDD（Domain-Driven Design，领域驱动设计）
├── configs
├── db
├── docs
├── domain 领域层 (所抓取网页的每个 Tab )
├── images
├── infra 基础设施层
│   ├── adapter (字符串处理的 适配器)
│   ├── config (配置信息 数据库配置等)
│   ├── crypt (认证密钥加密 登录)
│   ├── download (下载说要爬取网页信息)
│   ├── init (初始化DB)
│   └── model (分析网页结构后 说要爬取的对应字段定义 以及 json 字段定义)
├── Makefile
├── ui UI层 
│   └── api-server (API 设计)
│       ├── admins (用户认证)
│       ├── api_server.go (API 定义集合)
│       ├── awards (对应模块 API 处理)
│       ├── classic
│       ├── coaches
│       ├── controller
│       ├── groups
│       ├── matches
│       ├── players
│       ├── statistics
│       └── teams
├── vendor
│     └── vendor.json  
└── main.go 应用层

viper
Viper 是 Go 应用程序的完整配置解决方案，包括 12-Factor 应用程序。它旨在在应用程序中工作，并可以处理所有类型的配置需求和格式。它支持：
- 设置默认值
- 从 JSON，TOML，YAML，HCL 和 Java 属性配置文件中读取
- 实时观看和重新读取配置文件（可选）
- 从环境变量中读取
- 从远程配置系统（etcd 或 Consul）读取，并观察变化
- 从命令行标志读取
- 从缓冲区读取
- 设置显式值
- Viper 可以被认为是所有应用程序配置需求的注册表

构建现代应用程序时，您不必担心配置文件格式
