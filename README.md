## 目录结构
* build: 编译打包脚本文件
* config: 配置文件
* database: 数据库文件
* mvc: mvc 常规模块，参数封装、验证、返回值处理，异常处理, 后端控制器，i18n: 国际化
* logger: 日志管理
* routes: 路由定义
* release: 安装包发布
* client: http client，websocket工具类，ssm接口
* test: 测试脚本
* api: 业务逻辑处理
* initialization: 服务器初始化
* global: 全局变量
* middware: 中间件
* frontend: 存放前端项目 

## 编译条件
1. go >= 1.19
安装gcc-aarch64-linux-gnu：`sudo apt-get install gcc-aarch64-linux-gnu`
2. 安装docker(前端项目的编译在node17的docker容器中进行)

## 构建  

1. 进入build目录，执行build脚本，编译过程大概持续5分钟 
``` bash
cd build
./build_2_release.sh 
```

3. 安装包文件  
``` bash
release/
├── sophliteos-linux_amd64.tgz
├── sophliteos-linux_arm64.tgz
├── sophliteos_pcie_1.1.2.deb
└── sophliteos_soc_1.1.2.deb
``` 

## 安装运行
- 安装x86版本
tar -xvf sophliteos-amd64.tgz && sudo ./install.sh
或者 sudo dpkg -i sophliteos_pcie_1.1.2.deb
- 安装arm版本
tar -xvf sophliteos-arm64.tgz && sudo ./install.sh
或者 sudo dpkg -i sophliteos_soc_1.1.2.deb