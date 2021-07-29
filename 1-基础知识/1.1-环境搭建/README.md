### go语言环境搭建

自己目前只在使用mac，所以一下是mac的go语言环境搭建

### 使用brew安装go

```bash
brew install go
```

查看go是否安装成功
```bash
go version
```

### 配置环境变量

打开.zshrc
```bash
vim ~/.zshrc
```
将一下配置加入到.zshrc中
```bash
#GO
export GOROOT=/usr/local/Cellar/go/1.16.3/libexec
export GOPATH=$HOME/sdk/go
export PATH=${PATH}:$GOPATH/bin
```

查看结果
```bash
go env
```
我自己的结果如下
```bash
GO111MODULE="on"
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/zhengzhuang/Library/Caches/go-build"
GOENV="/Users/zhengzhuang/Library/Application Support/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/zhengzhuang/sdk/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/zhengzhuang/sdk/go"
GOPRIVATE=""
GOPROXY="https://goproxy.cn,direct"
GOROOT="/usr/local/Cellar/go/1.16.3/libexec"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/Cellar/go/1.16.3/libexec/pkg/tool/darwin_amd64"
GOVCS=""
GOVERSION="go1.16.3"
GCCGO="gccgo"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/dev/null"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -arch x86_64 -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/n9/0ssjxdp57594d8vvrgr61z700000gn/T/go-build1527436439=/tmp/go-build -gno-record-gcc-switches -fno-common"
```

### go module
go module是go官方自带的go依赖管理库,在1.13版本正式推荐使用<br>
go module可以将某个项目(文件夹)下的所有依赖整理成一个 go.mod 文件,里面写入了依赖的版本等<br>
使用go module之后我们可不用将代码放置在src下了

查看是否开启go module

```bash
go env
```
其中显示`GO111MODULE="on`则为开启，如果没有开始输入一下命令即可
```bash
export GO111MODULE=on
```
