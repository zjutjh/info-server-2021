# Info-backend

ZJUT新生信息查询后端

[![Go](https://github.com/zjutjh/info-backend/actions/workflows/go.yml/badge.svg)](https://github.com/zjutjh/info-backend/actions/workflows/go.yml)![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/zjutjh/info-backend)![GitHub](https://img.shields.io/github/license/zjutjh/info-backend)

## 如何编译

1. Clone 仓库到本地

   ```
   git clone https://github.com/zjutjh/info-backend
   ```

2. 更改工作目录

   ```
   cd info-backend
   ```

3. 编译

   ```
   go build
   ```

## 如何运行

1. 创建配置文件`config.yaml`

   配置文件自动查找目录：`.` `/etc/info/`

   或者使用`-c` `--config` 参数指定配置文件

   ```yaml
   # example config file
   # server config >>>
   server-port: ":8080"
   # database config >>>
   db-username: "root"
   db-password: "passwd"
   db-database: "info_test"
   # "db-hostname: 127.0.0.1" & "port: 3306" can be omitted
   # db-hostname: 127.0.0.1
   # db-port: 3306
   ```

2. 从Excel载入数据到MySQL

   > 注：Excel数据表格式很可能不通用。（导入已写的尽量智能，但是出了问题还是可能需要修改源码）

   ```
   ./info-backend -l [excel文件路径]
   ```

   如果Excel文件有密码：

   ```
   ./info-backend -l [excel文件路径] -p [密码]
   ```

   > 如果已导入基本数据，然后再导入寝室数据(更新数据库)，可以后面再加个`-u`，这样就不会执行数据库插入操作，避免打印记录重复的错误。

3. 启动服务端

   ```
   ./info-backend
   ```

## API定义

1. 获取基本信息

   ```http
   POST /api/v1/info HTTP/1.1
   Content-Type: application/json; charset=utf-8
   Host: 127.0.0.1:8080
   Connection: close
   User-Agent: Paw/3.2.2 (Macintosh; OS X/11.4.0) GCDHTTPRequest
   Content-Length: 51
   
   {"stu_name":"\u7cbe\u5c0f\u5f18","stu_id":"210001"}
   ```

   ``` http
   HTTP/1.1 200 OK
   Server: Tengine
   Date: Sat, 07 Aug 2021 03:35:39 GMT
   Content-Type: application/json; charset=utf-8
   Content-Length: 132
   Connection: close
   Expires: Sat, 07 Aug 2021 15:35:39 GMT
   Cache-Control: max-age=43200
   Cache-Control: no-cache
   
   {"code":200,"data":{"uid":"20210101","name":"精小弘","faculty":"精弘学院","class":"番茄一班","campus":"莫干山校区"}}
   ```

2. 获取寝室信息

   ```http
   POST /api/v1/dorm HTTP/1.1
   Content-Type: application/json; charset=utf-8
   Host: 127.0.0.1:8080
   Connection: close
   User-Agent: Paw/3.2.2 (Macintosh; OS X/11.4.0) GCDHTTPRequest
   Content-Length: 49
   
   {"stu_name":"info","stu_id":"330304200001010001"}
   ```

   ``` http
   HTTP/1.1 200 OK
   Server: Tengine
   Date: Sat, 07 Aug 2021 03:35:39 GMT
   Content-Type: application/json; charset=utf-8
   Content-Length: 132
   Connection: close
   Expires: Sat, 07 Aug 2021 15:35:39 GMT
   Cache-Control: max-age=43200
   Cache-Control: no-cache
   
   {"code":200,"data":{"uid":"20210101","name":"精小弘","faculty":"精弘学院","class":"番茄一班","campus":"莫干山校区"}
   ```

3. 异常情况 （例）

   ```http
   POST /api/v1/info HTTP/1.1
   Content-Type: application/json; charset=utf-8
   Host: 127.0.0.1:8080
   Connection: close
   User-Agent: Paw/3.2.2 (Macintosh; OS X/11.4.0) GCDHTTPRequest
   Content-Length: 51
   
   {"stu_name":"nobody","stu_id":"none"}
   ```

   ```http
   HTTP/1.1 200 OK
   Content-Type: application/json; charset=utf-8
   Date: Sat, 07 Aug 2021 03:53:36 GMT
   Content-Length: 35
   Connection: close
   
   {"code":404,"msg":"RecordNotFound"}
   ```

## 运行参数一览

```
Usage:
  info-backend [OPTIONS]

Application Options:
  -c, --config=  [PATH] Config file path
  -l, --load=    [PATH] Read & load data from excel
  -p, --passwd=  [PASS] Password of excel file
  -s, --sheet=   [Sheet] Read sheet
  -u, --update   Update database by dormitory info
  -v, --version  Show Info server version & quit

Help Options:
  -h, --help     Show this help message
```

## 项目结构

```
.
├── README.md
├── config.yaml
├── controller
│   └── get_stu_info.go
├── data
│   └── data.go
├── go.mod
├── go.sum
├── handler
│   ├── init_db.go
│   ├── query_record.go
│   ├── read_excel.go
│   └── request_check.go
├── main.go
└── model
    ├── models.go
    ├── options.go
    ├── request.go
    └── response.go

5 directories, 20 files
```

