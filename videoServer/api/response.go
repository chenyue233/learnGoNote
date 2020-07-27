package main

import (
	"net/http"
	"video_sever/videoServer/apieoServer/api/defs"
	"encoding/json"
	"io"
)

func sendErrorResponse(w http.ResponseWriter,errResp defs.ErrResponse)  {
	w.WriteHeader(errResp.HttpSC)

	resStr,_ := json.Marshal(&errResp.Error)
	io.WriteString(w,string(resStr))
}

func sendNormalResponse(w http.ResponseWriter,resp string,sc int)  {
	w.WriteHeader(sc)
	io.WriteString(w,resp)
}
