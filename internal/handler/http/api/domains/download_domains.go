package domains

import (
	"go-project/internal/common"
	"net/http"
	"os"
	"path"
	"runtime"
)

func (a *Api) Download(w http.ResponseWriter, r *http.Request) {
	_, b, _, _ := runtime.Caller(0)
	d1 := path.Join(path.Dir(b))
	file, errOpen := a.storage.Get(d1 + "/../../../../var/api.csv")
	if errOpen != nil {
		common.SendErrorResponse(w, errOpen.Error())
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	common.SendFile(w, r, file)
}
