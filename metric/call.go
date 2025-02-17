package metric

import (
	"database/sql"
	"fmt"
)

type Metric struct {
	db *sql.DB
}

func NewMetric(db *sql.DB) *Metric {
	return &Metric{db}
}

func (m *Metric) GetCountCall() int {
	rows, err := m.db.Prepare("SELECT COUNT(id) AS cnt FROM call_summaries USE INDEX (`PRIMARY`) WHERE line_number = '6030000'")

	if err != nil {
		panic(err.Error())
	}

	var countCalls int

	err = rows.QueryRow(1).Scan(&countCalls)

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	return countCalls
}
