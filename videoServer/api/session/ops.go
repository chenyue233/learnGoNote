package session

import (
	"sync"
	"video_sever/videoServer/apieoServer/api/dbops"
	"video_sever/videoServer/apieoServer/api/defs"
	"video_sever/videoServer/apieoServer/api/utils"
	"time"
)

var sessionMap  *sync.Map

func init(){
	sessionMap = &sync.Map{}
}
func nowInMilli() int64 {
	return time.Now().UnixNano()/1000000
}

func deleteExpiredSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
func LoadSessionFromDB()  {
	r, err := dbops.RetrieveAllSession()
	if err != nil{
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k,ss)
		return true
	})

}

func GenerateNewSessionId(un string) string {
	id ,_ :=utils.NewUUID()
	ct := nowInMilli()
	ttl := ct * 30 * 60 *1000  // session本地过期时间

	ss := &defs.SimpleSession{Username:un,TTL:ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)

	return id
}

func IsSessionExpired(sid string) (string,bool) {
	ss,ok := sessionMap.Load(sid)
	if ok{
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			// 删除数据库中的session
			deleteExpiredSession(sid)
			return "",true
		}
		return ss.(*defs.SimpleSession).Username,false
	}
	return "",true
}
