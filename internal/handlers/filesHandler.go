package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
)

func (s *server) handlePhotoUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseMultipartForm(10 << 20)

		file, handler, err := r.FormFile("newsPhoto")
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		defer file.Close()

		dst, err := os.Create("./static/news/" + handler.Filename)

		defer dst.Close()

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		filePath := path.Join("./static/news/" + handler.Filename)

		if _, err := io.Copy(dst, file); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusCreated, map[string]interface{}{"status": "successfully create", "file path": filePath})
	}
}

func (s *server) handlePhotoDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileName := &FileName{}

		if err := json.NewDecoder(r.Body).Decode(fileName); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		filePath := "./static/news/" + fileName.Name

		err := os.Remove(filePath)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, map[string]interface{}{"status": "successfully delete", "file path": filePath})
	}
}
