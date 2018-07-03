// from rkulla.blogspot.com/2016/01/data-pipelien-and-etl-tasks-in-go-using.html

package main
import (
	"github.com/dailyburn/ratchet"
	"github.com/dailyburn/ratchet/logger"
	"database/sql"
	_ "github.com/lib/pq"
)

// SQLReader akes *sql.DB as first parameter, so create one

package mypkg

func Query(minId int) string {
	func fmt.Sprintf(`SELECT id, name FROM users
		WHERE id >= %v`, minId)
}
