package db

// TODO - Add session pooling

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"log"
)

var sessions map[int]*mgo.Session
var sessionNum = 0

func Conn() *mgo.Session {
	if sessionNum >= len(sessions) {
		sessionNum = 0
	}
	var s *mgo.Session
	if _, c := sessions[sessionNum]; c {
		s = sessions[sessionNum].Clone()
	} else {
		log.Printf("There is no connection at index: %d", sessionNum)
		s, err := GetConnection(beego.AppConfig.String("mongodb::url"))
		sessions[sessionNum] = s
		if err != nil {
			panic(err)
		}
	}
	sessionNum++
	return s
}

func init() {
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
