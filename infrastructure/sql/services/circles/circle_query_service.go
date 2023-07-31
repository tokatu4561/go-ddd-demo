package circles

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


type CircleSummaryData struct {
	Id string `db:"circleId"`
	Name string `db:"ownerName"`
}

type CircleGetSummaryCommand struct {
	Size int
	Page int
}

type CircleQueryService struct {
	db *sqlx.DB
}

func NewCircleQueryService() (*CircleQueryService, error) {
	dsn := os.Getenv("dsn")
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &CircleQueryService{
		db: db,
	}, nil
}

func (sqs CircleQueryService) GetSummaries(command *CircleGetSummaryCommand) ([]*CircleSummaryData, error) {
	// sql で サークルのサマリーを取得する
	// FIXME: implement 正しく実装する 動くかわからない
	rows, err := sqs.db.Query("SELECT circles.id as circleId, users.name as ownerName FROM circles LEFT OUTER JOIN users ON circles.ownerId = users.id ORDER BY circles.id OFFSET $1 ROWS FETCH NEXT $2 ROWS ONLY", command.Size * command.Page, command.Size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var circleSummaries []*CircleSummaryData
	for rows.Next() {
		var circleId string
		var ownerName string
		err := rows.Scan(&circleId, &ownerName)
		if err != nil {
			return nil, err
		}
		circleSummary := &CircleSummaryData{Id: circleId, Name: ownerName}
		circleSummaries = append(circleSummaries, circleSummary)
	}
	
	return circleSummaries, nil
}