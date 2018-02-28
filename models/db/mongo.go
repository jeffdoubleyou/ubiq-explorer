package db

// TODO - Add session pooling

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"log"
	"math/rand"
	"time"
)

var sessions map[int]*mgo.Session
var sessionNum = 0

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
