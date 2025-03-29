# 使用 goenv 安装多版本 go

> [goenv](https://github.com/go-nv)

查看开源项目的时候，经常看到不同项目对 go 的版本要求不一致，查了一下相关的工具，类似 Nodejs 中的 nvm 和 volta 工具，有好几个，以下是 goenv 的安装和使用。


## 安装
参考： https://github.com/go-nv/goenv/blob/master/INSTALL.md

```sh
# mac os
brew update
brew install goenv
```

同时配置 `.zshrc` 文件，添加到最后；

```sh
# goenv
export GOENV_ROOT="$HOME/.goenv"
export PATH="$GOENV_ROOT/bin:$PATH"
export GOROOT_MIRROR=https://goproxy.cn/dl # 清华大学镜像

eval "$(goenv init -)"

export PATH=$PATH:$(go env GOPATH)/bin
```

测试：


```sh
goenv --version 
# 1.21.13 (set by /Users/[Your Name]/.goenv/version)
```


## 常用：

```sh
# 查看所有命令
goenv commands


# 安装 go 和使用
goenv install go@1.24

goenv global go@1.24


# 卸载 go
goenev uninstall go@1.24

# 查看所有安装的go 版本
goenv versions

```


更多命令：https://github.com/go-nv/goenv/blob/master/COMMANDS.md
