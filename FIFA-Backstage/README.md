```bash
DDD（Domain-Driven Design，领域驱动设计）

FIFA-Backstage
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
├── ui UI层 
│   └── api-server (API 设计)
├── vendor 
└── main.go 应用层
```

主要使用技术：

- [x] gorm 操作数据库
- [x] fresh 实现 web server 监听
- [x] viper 实现读取用户配置
- [x] 数据库 使用 postgre
- [x] goquery 实现网页解析
- [X] gin 快速搭建 web server (只列举了部分实例，举一反三)
- [X] gin-swagger 自动化构建 API 文档 (只列举了部分实例，举一反三)
- [X] jwt 认证，后台生成对应的 auth_token