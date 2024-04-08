# go http配置sqlite ini配置文件
## 编译 win exe
    //sudo apt-get install gcc-mingw-w64
    CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build
## 编译 aarch64
    //根据情况选择链版本
    CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build
## 编译 arm
    //根据情况选择链版本
    CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=arm-linux-gnueabihf-gcc go build

## 注意

    编译aarch64和arm的时候输出产物的名称都是工程名称，需要用file指令看下产物的cpu架构信息

### 20230901
    1.0.5 增加多目远程控制

### 20231212
    1.0.6 增加调试控制,增加违法抓拍控制