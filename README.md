## 目录结构
* build: 编译打包脚本文件
* config: 配置文件
* database: 数据库文件
* mvc: mvc 常规模块，参数封装、验证、返回值处理，异常处理，handlers: 后端控制器，i18n: 国际化
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
1. 将前端项目下载到frontend文件夹  
``` bash
cd frontend
git clone http://172.28.141.70/peacenet/sophliteos-frontend.git
或者
git clone https://gerrit-ai.sophgo.vip:8443/AI_SE/sophliteos-frontend.git 
```  
2. 进入build目录，执行build脚本，编译过程大概持续5分钟 
``` bash
cd build
./build_2_release.sh 
```
- 如果系统执行docker命令需要root权限，需要执行下面命令
确保go环境变量加入到root用户中，参考如下：
``` bash
vim /root/.bashrc
# 文件的末尾添加Go语言的环境变量
export PATH=$PATH:/usr/local/go/bin
export GOPATH=/root/go
# 更改生效
source /root/.bashrc
```
- 执行build脚本 
``` bash
cd build
sudo ./build_2_release.sh 
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