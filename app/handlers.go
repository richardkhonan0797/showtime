package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if tz, ok := q["tz"]; ok {
		var res []map[string]time.Time
		tz := strings.Split(tz[0], ",")

		if len(tz[0]) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("please input time zones"))
			log.Println("please input time zones")
		} else {
			for _, val := range tz {
				loc, err := time.LoadLocation(val)

				if err != nil {
					w.WriteHeader(http.StatusNotFound)
					w.Write([]byte(fmt.Sprintf("invalid timezone %s", val)))
					log.Println(fmt.Sprintf("invalid timezone %s", val))
					return
				}

				time := map[string]time.Time{val: time.Now().In(loc)}
				res = append(res, time)
			}

			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
		}
	} else {
		res := map[string]time.Time{"current_time": time.Now().UTC()}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
