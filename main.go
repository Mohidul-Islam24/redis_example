package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Users struct {
	Id       int
	UserName string
}
type server struct {
	db    map[int]*Users
	cache map[int]*Users
	dbHit int
}

func NewServer() *server {
	db := make(map[int]*Users)
	for i := 1; i <= 100; i++ {
		db[i] = &Users{
			Id:       i,
			UserName: fmt.Sprintf("userName_%d", i),
		}
	}
	return &server{
		db:    db,
		cache: make(map[int]*Users),
	}
}
func (s *server) tryCache(id int) (*Users, bool) {
	user, ok := s.cache[id]
	return user, ok
}
func (s *server) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(strId)

	//first try in cache
	user, ok := s.tryCache(id)
	if ok {
		json.NewEncoder(w).Encode(user)
		return
	}

	//hit the db
	user, ok = s.db[id]
	if !ok {
		panic("user not found")
	}
	s.dbHit++
	//insert into the cache
	s.cache[id] = user

	json.NewEncoder(w).Encode(user)
}
func main() {}
