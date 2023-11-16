# SOPHLITEOS算法业务
## 目录结构
* api: 接口业务逻辑处理
* build: 编译打包脚本文件
* client: http client请求
* config: 配置文件和配置读取驱动
* database: 数据库文件
* global: 全局变量
* initialization: 服务启动初始化
* logger: 日志管理
* middware: 中间件
* mvc: 结构体配置
* release: 安装包发布
* routes: 路由定义
* scrip: 部署脚本


## 编译条件
1. go版本 >= 1.19
2. 安装gcc-aarch64-linux-gnu：`sudo apt-get install gcc-aarch64-linux-gnu`


## 构建  
1. 进入build目录，执行build脚本  
``` bash
cd build
./build.sh 
```


3. 安装包文件  
``` bash
release/
├── algoliteos-linux_amd64.tgz
└── algoliteos-linux_arm64.tgz
``` 

## 安装运行
- 安装x86版本  
tar -xvf algoliteos-amd64.tgz && sudo ./install.sh

- 安装arm版本  
tar -xvf algoliteos-arm64.tgz && sudo ./install.sh

