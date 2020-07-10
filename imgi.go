package imgi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	config Config
}

func NewServer(config Config) *Server {
	srv := &Server{
		config: config,
	}
	return srv
}

func (srv *Server) Start() {
	config := srv.config
	addr := ":" + strconv.Itoa(config.General.Port)
	log.Infof("imgi start listening on %s", addr)
	http.ListenAndServe(addr, createHandler(config))
}

func createHandler(config Config) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil { // revover from unknown panicking and response with 500 error
				log.Error(p)
				replyError(w, r, errors.New("unknown error"), http.StatusInternalServerError)
			}
		}()

		img, err := fetchImage(config, r)
		if err != nil {
			replyError(w, r, err, http.StatusNotFound)
			return
		}

		query := r.URL.Query().Get("imgi")

		ops, err := parseOperations(query)
		if err != nil {
			replyError(w, r, err, http.StatusBadRequest)
		}
		if len(ops) == 0 {
			replyImage(w, r, img)
			return
		}
		for _, op := range ops {
			img, err = op.Action.Act(img, op.Options)
			if err != nil {
				replyError(w, r, err, http.StatusInternalServerError)
			}
		}
		replyImage(w, r, img)
	})
	return mux
}

func replyImage(w http.ResponseWriter, r *http.Request, img Image) {
	w.Header().Set("Content-Type", img.Mime)
	w.Header().Set("Content-Length", strconv.Itoa(len(img.Buf)))
	w.Write(img.Buf)
}

func fetchImage(config Config, r *http.Request) (Image, error) {
	img := Image{}

	path := r.URL.Path
	var file string
	for k, v := range config.Locations {
		if strings.HasPrefix(path, k) {
			path = strings.TrimPrefix(path, k)
			file = filepath.Join(v, path)
			break
		}
	}
	if file == "" {
		return img, ErrNotFound
	}

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return img, ErrNotFound
	}
	img.Buf = buf
	img.Mime = http.DetectContentType(buf)
	return img, nil
}
