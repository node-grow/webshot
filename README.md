# Webshot 网页截图服务

场景模拟：

1. 分享图片时需传入url，但由后端直接生成图片难度较大，有对应的h5分享页时则可使用此服务
2. 小程序生成分享图时会时不时地进入各种坑，则可以写一个h5分享页来曲线救国

## 使用方法

### 部署

1. docker部署（推荐）
   ```shell
    # 编译
    docker build -t go-webshot .
    # 运行
    docker run -p 8080:8080 go-webshot
   ```
2. 直接运行（确认本地已安装chrome浏览器和golang），windows用户请在Git Bash中运行
   ```shell
   # 编译
   go build -v -a -o build/go-webshot ./main.go
   # 运行
   ./build/go-webshot
   ```
   
#### 运行参数【可选】

- --cache-dir：缓存目录，默认为系统临时目录
- --cache-day：缓存天数，默认为7天

### 访问

[主机名:8080]/[配置项]/[url]
### 示例
http://127.0.0.1:8080/1600x900,q_50/https:/baidu.com

### 配置项说明

每个配置项以英文逗号```,```作为分隔

1. 分辨率设置：【宽度】x【高度】 例，1600x900
2. 质量设置：q_【质量】 例，q_50

### url说明

1. url 请带上 http(s):// 协议
2. 因该程序带有截图缓存机制，若要更新网页截图，请将缓存文件删除或在url后加上任意query参数
