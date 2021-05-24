# Go-tour
Go小型專案 練習



## 單字格式轉換

查看說明

```sh
go run main.go help word
```

使用範例 : 字串 getAttritube 透過mode=5 得到 get_attritube

```sh
go run main.go word -s=getAttritube -m=5
```

**自動化測試** word_test.go

---

## 時間工具

查看說明

```sh
go run main.go help time

go run main.go help time now

go run main.go help time calc
```

使用範例 :  返回當下時間以及Unix timestamp

```sh
go run main.go time now
```

使用範例 : 根據-c的時間 對-d的參數做增減 返回時間以及Unix timestamp

1. -c="2021-05-12 00:00:00" -d=5m 得到 2021-05-12 00:05:00, 1620749100

2. -c="2021-05-12 00:00:00" -d=-2h 得到 2021-05-11 22:00:00, 1620741600

```sh
go run main.go time calc -c="2021-05-12 00:00:00" -d=5m
go run main.go time calc -c="2021-05-12 00:00:00" -d=-2h
```

**自動化測試** time_test.go
