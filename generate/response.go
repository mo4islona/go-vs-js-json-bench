package generate

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"math/rand"
	"strconv"
)

type Meta struct {
	Bulls MetaInfo `json:"bulls"`
	Bears MetaInfo `json:"bears"`
}

type MetaInfo struct {
	Bets int     `json:"bets"`
	Bank float64 `json:"bank"`
}

type AssetValue struct {
	Value     float64   `json:"asset"`
	Timestamp time.Time `json:"timestamp"`
}

type Bet struct {
	Id         bson.ObjectId `json:"id"`
	Team       int64         `json:"team"`
	Bet        float64       `json:"bet"`
	AuthorId   bson.ObjectId `json:"authorId"`
	Payout     float64       `json:"payout"`
	Commission float64       `json:"commission"`
	Created    time.Time     `json:"created"`
}

type PublicUser struct {
	Id       bson.ObjectId `json:"id"`
	Username string        `json:"username"`
	Avatar   bson.ObjectId `json:"avatar"`
}

type Model struct {
	Id          bson.ObjectId   `json:"id"`
	Title       string          `json:"title"`
	Meta        *Meta           `json:"meta"`
	Type        int64           `json:"type"`
	Asset       string          `json:"asset"`
	Bet         float64         `json:"bet"`
	Bets        []*Bet          `json:"bets"`
	FullMembers []*PublicUser   `json:"members"`
	FistAsset   AssetValue      `json:"firstAsset"`
	LastAsset   AssetValue      `json:"lastAsset"`
	AuthorId    bson.ObjectId   `json:"authorId"`
	Status      int64           `json:"status"`
	GameStatus  int64           `json:"gameStatus"`
	Created     time.Time       `json:"createdAt"`
	Deadline    time.Time       `json:"deadlineAt"`
	Expired     time.Time       `json:"expiredAt"`
	BotOnly     bool            `json:"-"`
}

var Assets []string = []string{
	"DEMO",
	"EURUSD",
	"USDJPY",
	"GBPUSD",
	"AUDUSD",
	"USDCHF",
	"NZDUSD",
	"USDCAD",
}

func Response(l int) []*Model {
	rooms := make([]*Model, l)
	for i := 0; i < l; i++ {
		rooms[i] = &Model{
			Id:    bson.NewObjectId(),
			Title: "title" + strconv.Itoa(i),
			Meta: &Meta{
				Bears: MetaInfo{
					Bets: 10,
					Bank: rand.Float64() * 3000,
				},
				Bulls: MetaInfo{
					Bets: 10,
					Bank: rand.Float64() * 3000,
				},
			},
			Type:  1,
			Asset: Assets[rand.Intn(len(Assets))],
			Bet:   rand.Float64() * 100,
			Bets: []*Bet{
				{
					Id:         bson.NewObjectId(),
					Bet:        rand.Float64() * 100,
					AuthorId:   bson.NewObjectId(),
					Payout:     rand.Float64() * 1000,
					Commission: rand.Float64() * 20,
					Created:    time.Now(),
				},
				{
					Id:         bson.NewObjectId(),
					Bet:        rand.Float64() * 100,
					AuthorId:   bson.NewObjectId(),
					Payout:     rand.Float64() * 1000,
					Commission: rand.Float64() * 20,
					Created:    time.Now(),
				},
				{
					Id:         bson.NewObjectId(),
					Bet:        rand.Float64() * 100,
					AuthorId:   bson.NewObjectId(),
					Payout:     rand.Float64() * 1000,
					Commission: rand.Float64() * 20,
					Created:    time.Now(),
				},
				{
					Id:         bson.NewObjectId(),
					Bet:        rand.Float64() * 100,
					AuthorId:   bson.NewObjectId(),
					Payout:     rand.Float64() * 1000,
					Commission: rand.Float64() * 20,
					Created:    time.Now(),
				},
				{
					Id:         bson.NewObjectId(),
					Bet:        rand.Float64() * 100,
					AuthorId:   bson.NewObjectId(),
					Payout:     rand.Float64() * 1000,
					Commission: rand.Float64() * 20,
					Created:    time.Now(),
				},
				{
					Id:         bson.NewObjectId(),
					Bet:        rand.Float64() * 100,
					AuthorId:   bson.NewObjectId(),
					Payout:     rand.Float64() * 1000,
					Commission: rand.Float64() * 20,
					Created:    time.Now(),
				},
			},
			FullMembers: []*PublicUser{
				{
					Id:       bson.NewObjectId(),
					Username: "player" + strconv.Itoa(i),
					Avatar:   bson.NewObjectId(),
				},
				{
					Id:       bson.NewObjectId(),
					Username: "player" + strconv.Itoa(i),
					Avatar:   bson.NewObjectId(),
				},
				{
					Id:       bson.NewObjectId(),
					Username: "player" + strconv.Itoa(i),
					Avatar:   bson.NewObjectId(),
				},
				{
					Id:       bson.NewObjectId(),
					Username: "player" + strconv.Itoa(i),
					Avatar:   bson.NewObjectId(),
				},
			},
			FistAsset: AssetValue{
				Value:     rand.Float64() * 200,
				Timestamp: time.Now(),
			},
			LastAsset: AssetValue{
				Value:     rand.Float64() * 200,
				Timestamp: time.Now(),
			},
			AuthorId:   bson.NewObjectId(),
			Status:     1,
			GameStatus: 2,
			Created:    time.Now(),
			Deadline:   time.Now(),
			Expired:    time.Now(),
			BotOnly:    false,
		}
	}
	return rooms
}
