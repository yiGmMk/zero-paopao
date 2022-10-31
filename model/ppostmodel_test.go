package model

import (
	"context"
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func TestFetchPost(t *testing.T) {
	m := NewPPostModel(sqlx.NewMysql("paopao:paopao@tcp(localhost:10020)/paopao?charset=utf8mb4&parseTime=True"))
	b := m.RowBuilder().Where(squirrel.Eq{"user_id": []int{100058, 100059}})
	data, err := m.Fetch(context.Background(), b, 0, 10)
	if err != nil {
		t.Error(err, "")
	}
	fmt.Println(len(data))
}
