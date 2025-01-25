package CopyDatabase

import (
	"context"
	"log"
	"strings"

	"github.com/KEHyeon/pgxCopy/internal/utils"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CopyToDatabase(pool *pgxpool.Pool, data []tableNameProvider) (int64, error) {
	// 첫 번째 요소에서 컬럼 추출
	columns := utils.GetColumnFromStruct(data[0], "json")

	// 데이터 값 추출
	valueRows := make([][]interface{}, len(data))
	for i, item := range data {
		values := utils.GetValuesFromStruct(item, "json")
		valueRows[i] = values
	}

	// CopyFrom 실행
	tableName := data[0].TableName()
	parts := strings.Split(tableName, ".")

	if len(parts) != 2 {
		log.Fatalf("Invalid table name format: %v. Expected format 'schema.table'", tableName)
	}

	// 분리된 스키마와 테이블 이름을 사용하여 pgx.Identifier 생성
	return pool.CopyFrom(context.Background(), pgx.Identifier{parts[0], parts[1]}, columns, pgx.CopyFromRows(valueRows))

}
