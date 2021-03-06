package chart

import (
	"gorm.io/gorm"
	"net/http"
)

var GlobHub *Hub

// Framework-based operation
func Server(w http.ResponseWriter, r *http.Request, channelId string, clientId uint64, db *gorm.DB) {
	if GlobHub == nil{
		GlobHub = NewServer()
		go GlobHub.Run()
	}
	if clients == nil {
		clients = make(map[uint64]*Client)
	}
	DB = db
	RunWs(GlobHub, w, r, channelId, clientId)
}

func ChannelAllUserId(channelId string) []uint64 {
	var ids []uint64
	if list, ok := GlobHub.clients[channelId]; ok {
		for c, _ := range list {
			ids = append(ids, c.Id)
		}
	}
	return ids
}

func AddUserNum() int {
	num := 0
	for _, v := range GlobHub.clients {
		num += len(v)
	}
	return num
}
