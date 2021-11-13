package handler

import (
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
)

const redirect_to = "https://abranhe.com"

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		w.Header().Set("Location", redirect_to)
		w.WriteHeader(http.StatusMovedPermanently)
		return
	}

	if path[0] == '/' {
		path = path[1:]
	}

	conn, err := redis.DialURL(os.Getenv("REDIS_CONN_URL"))

	if err != nil {
		w.Header().Set("Location", redirect_to)
		w.WriteHeader(http.StatusMovedPermanently)
		return
	}

	defer conn.Close()

	domain, err := redis.String(conn.Do("GET", "SHORT_URL:"+path))

	if err != nil {
		w.Header().Set("Location", redirect_to)
		w.WriteHeader(http.StatusMovedPermanently)
		return
	}

	w.Header().Set("Location", domain)
	w.WriteHeader(http.StatusMovedPermanently)
}
