package db

// TODO - Add session pooling

import (
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
	"time"
)

type UpsertCollection struct {
	Find *bson.M
	Data interface{}
}

var poolSize int = 25
var sessions map[int]*mgo.Session
var sessionNum = 0
var pool = make(map[string]chan interface{}, poolSize)
var upsertPool = make(map[string]chan *UpsertCollection, poolSize)

func Conn() *mgo.Session {
	return sessions[rand.Intn(len(sessions))].Clone()
}

func init() {
	rand.Seed(time.Now().Unix())
	url := beego.AppConfig.String("mongodb::url")
	maxConnections, err := beego.AppConfig.Int("mongodb::max_connections")
	if err != nil {
		log.Printf("mongodb::max_connections not set - defaulting to 5 connections")
		maxConnections = 5
	}
	sessions = make(map[int]*mgo.Session)
	for i := 0; i < maxConnections; i++ {
		sess, err := GetConnection(url)
		if err != nil {
			panic(err)
		}
		sess.SetMode(mgo.Monotonic, true)
		sessions[i] = sess
	}
}

func GetConnection(url string) (*mgo.Session, error) {
	return mgo.Dial(url)
}

func Insert(collection string, data interface{}) error {
	if _, ok := pool[collection]; !ok {
		fmt.Printf("Create pool for %s\n", collection)
		pool[collection] = make(chan interface{}, poolSize)
	}
	pool[collection] <- data
	go SyncInsert(collection)
	return nil
}

func Upsert(collection string, find *bson.M, data interface{}) error {
	if _, ok := upsertPool[collection]; !ok {
		fmt.Printf("Create upsert pool for %s\n", collection)
		upsertPool[collection] = make(chan *UpsertCollection, poolSize)
	}
	upsertPool[collection] <- &UpsertCollection{find, data}
	go SyncUpsert(collection)
	return nil
}

func SyncInsert(collection string) {
	conn := Conn()
	defer conn.Close()
	//fmt.Printf("Have %d items in the %s pool\n", len(pool[collection]), collection)
	for i := 0; i < len(pool[collection]); i++ {
		data := <-pool[collection]
		err := conn.DB("").C(collection).Insert(data)
		if err != nil {
			panic("Failed to insert to DB from pool")
		}
	}
}

func SyncUpsert(collection string) {
	conn := Conn()
	defer conn.Close()

	//fmt.Printf("Have %d items in the %s upsert pool\n", len(upsertPool[collection]), collection)
	for i := 0; i < len(upsertPool[collection]); i++ {
		data := <-upsertPool[collection]
		_, err := conn.DB("").C(collection).Upsert(data.Find, data.Data)
		if err != nil {
			panic("Failed to upsert to DB from pool")
		}
	}
}
