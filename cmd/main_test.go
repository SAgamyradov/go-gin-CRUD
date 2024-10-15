package main

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// )

// func handleRequest(w *httptest.ResponseRecorder, r *http.Request) {
// 	router := GetRouters()
// 	router.ServeHTTP(w, r)
// }

// func TestAlbumList(t *testing.T) {
// 	request, _ := http.NewRequest("GET", "/albums", strings.NewReader(""))
// 	w := httptest.NewRecorder()
// 	handleRequest(w, request)
// 	if w.Code != 200 {
// 		t.Fatal("status no ok")
// 	}
// }

// func TestAlbumDetail(t *testing.T) {
// 	albumId := "1"
// 	request, _ := http.NewRequest("GET", "/albums/"+albumId, strings.NewReader(""))
// 	w := httptest.NewRecorder()
// 	handleRequest(w, request)
// 	if w.Code != 200 {
// 		t.Fatal("status no ok")
// 	}
// }

// func TestAlbumNotFound(t *testing.T) {
// 	albumId := "9999"
// 	request, _ := http.NewRequest("GET", "/albums/"+albumId, strings.NewReader(""))
// 	w := httptest.NewRecorder()
// 	handleRequest(w, request)
// 	if w.Code != http.StatusNotFound {
// 		t.Fatal("status no ok")
// 	}
// }

// func TestDelete(t *testing.T) {
// 	albumId := "1"
// 	request, _ := http.NewRequest("DELETE", "/albums/"+albumId, strings.NewReader(""))
// 	w := httptest.NewRecorder()
// 	handleRequest(w, request)
// 	if w.Code != http.StatusNoContent {
// 		t.Fatal("status could not delete")
// 	}
// }

// func TestDeleteNotContent(t *testing.T) {
// 	albumId := "999"
// 	request, _ := http.NewRequest("DELETE", "/albums/"+albumId, strings.NewReader(""))
// 	w := httptest.NewRecorder()
// 	handleRequest(w, request)
// 	if w.Code != http.StatusNoContent {
// 		t.Fatal("not found")
// 	}
// }
