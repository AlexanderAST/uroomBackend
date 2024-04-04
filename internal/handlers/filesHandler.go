package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func deleteSpaces(str string) string {
	result := strings.ReplaceAll(str, " ", "")
	return result
}

func (s *server) handlePhotoUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseMultipartForm(10 << 20)

		file, handler, err := r.FormFile("newsPhoto")
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		defer file.Close()

		newName := deleteSpaces(handler.Filename)

		dst, err := os.Create("./static/news/" + newName)

		defer dst.Close()

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		filePath := path.Join("./static/news/" + newName)

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

func (s *server) handleGetPhoto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileName := r.URL.Query().Get("filename")
		if fileName == "" {
			s.error(w, r, http.StatusBadRequest, errors.New("filename parameter is required"))
			return
		}

		file, err := os.Open(fileName)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Type", "image/png")

		if _, err := io.Copy(w, file); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
	}
}
