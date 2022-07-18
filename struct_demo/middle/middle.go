package middle

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.Header)
	fmt.Fprintln(w, string(data))
}
