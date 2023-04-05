package serializer

import (
	"fmt"
	"net/http"
)

func SerializeError(w http.ResponseWriter, err error) {
	str := fmt.Sprintf("{\"error\":\"%v\"}", err)
	w.Write([]byte(str))
}
