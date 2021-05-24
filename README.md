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

---

## MySQL根據連線資訊、資料庫名稱和表名稱，轉換成Go語言的結構

先建立如下表，以供後面結果參考

```sql
CREATE TABLE `articles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(50) NOT NULL COMMENT '標題',
  `content` longtext,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `articles_user_id_b9d9571c_fk_auth_user_id` (`user_id`),
  CONSTRAINT `articles_user_id_b9d9571c_fk_auth_user_id` FOREIGN KEY (`user_id`) REFERENCES `auth_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```



查看說明

```sh
go run main.go help sql

go run main.go help sql struct
```

> $ go run main.go help sql struct
> sql轉換
>
> Usage:
>    sql struct [flags]
>
> Flags:
>       --charset string    請輸入資料庫的編碼 (default "utf8mb4")
>       --db string         請輸入資料庫名稱
>   -h, --help              help for struct
>       --host string       請輸入資料庫的HOST (default "127.0.0.1:3306")
>       --password string   請輸入資料庫的密碼
>       --table string      請輸入表名稱
>       --type string       請輸入資料庫實例類型 (default "mysql")
>       --username string   請輸入資料庫的帳號**

使用範例 :  

```sh
go run main.go sql struct --username=帳號 --password=密碼 --db=資料庫名稱 --table=表名稱
```

返回結果(Go的程式碼，該table的struct結構)

```go
type Articles struct {
         // content
         Content        string  `json:"content"`
         // created_at
         CreatedAt      time.Time       `json:"created_at"`
         // id
         Id     int32   `json:"id"`
         // 標題
         Title  string  `json:"title"`
         // updated_at
         UpdatedAt      time.Time       `json:"updated_at"`
         // user_id
         UserId int32   `json:"user_id"`
}

func (model Articles) TableName() string {
        return "articles"
}
```

