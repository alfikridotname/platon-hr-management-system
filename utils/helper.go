package utils

import (
	"fmt"

	"gorm.io/gorm"
)

func RawSql(result *gorm.DB) {
	fmt.Println(result.Statement.SQL.String())
}
