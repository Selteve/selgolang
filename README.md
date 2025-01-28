# Selgolang

Selgolang 是一个用于创建新项目的工具。

## 安装

首先，您需要将 Selgolang 安装到您的本地环境中。请确保您已经安装了 Go 语言环境，并且 Go 版本在 1.22.5 及以上。

使用以下命令安装 Selgolang：

```sh
go install github.com/Selteve/selgolang@latest
```

## 使用

安装完成后，您可以使用`selgolang`命令来创建新项目.

### 创建新项目

使用以下命令创建一个新项目：

```sh
selgolang new <project_name>
```
其中, `<projectname>` 是您要创建的项目名称。例如，如果您要创建一个名为 mapp 的项目，可以运行以下命令：

```sh
selgolang new mapp
```

该命令将从模板仓库中创建一个新项目，并将其放置在当前目录下的 mapp 文件夹中。

### 示例
以下是一个完整的示例：
```sh
# 安装 Selgolang
go install github.com/Selteve/selgolang@latest

# 创建一个名为 "myproject" 的新项目
selgolang new myproject
```