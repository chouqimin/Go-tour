package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 因DataType欄位類型與Go結構中的類型不是完全一致的(如varchar、longtext、TimeStamp等)
var DBTypeToStructType = map[string]string{
	"int":       "int32",
	"tinyint":   "int8",
	"smallint":  "int",
	"mediumint": "int64",
	"bigint":    "int64",
	"bit":       "int",

	"bool": "bool",

	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",

	"date":      "time.Time",
	"datetime":  "time.Time",
	"timestamp": "time.Time",
	"time":      "time.Time",

	"float":  "float64",
	"double": "float64",
}

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	Username string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"

	// dsn:Data source name
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.Username,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)

	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn) // 第一個參數為驅動名稱(如:mysql)，第二個參數為驅動連接資料庫的連接資訊
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := `
		SELECT 
			COLUMN_NAME,
			DATA_TYPE,
			COLUMN_KEY,
			IS_NULLABLE,
			COLUMN_TYPE,
			COLUMN_COMMENT
		FROM
			information_schema.COLUMNS
		WHERE
			TABLE_SCHEMA = ?
				AND TABLE_NAME = ?;
	`
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("沒有資料")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(
			&column.ColumnName,
			&column.DataType,
			&column.IsNullable,
			&column.ColumnKey,
			&column.ColumnType,
			&column.ColumnComment,
		)

		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}
