# go http配置sqlite ini配置文件
## 编译 win exe
//sudo apt-get install gcc-mingw-w64
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build
## 编译
CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build


### 20230901
    1.0.5 增加多目远程控制

### 20231212
    1.0.6 增加调试控制