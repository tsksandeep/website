package download

import (
	"net/http"
	"path/filepath"

	"github.com/website/handlers"
	"github.com/website/httputils"

	log "github.com/sirupsen/logrus"
)

type downloadHanlder struct{}

// New creates a new instance of Contact Handler
func New() handlers.DownloadHandler {
	return &downloadHanlder{}
}

func (ch *downloadHanlder) GetResume(w http.ResponseWriter, r *http.Request) {

	absPath, err := filepath.Abs("./static/sandeep_resume.pdf")
	if err != nil {
		log.Error("error resolving relative path for sandeep resume")
		handlers.WriteHandlerError(err, http.StatusInternalServerError, httputils.UnexpectedError, w, r)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(absPath))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, absPath)
}
