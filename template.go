package web_golang

import (
  "web-golang/entity"
  "context"
  _ "github.com/go-sql-driver/mysql"
  "testing"
  _"fmt"
  "net/http"
  _"net/http/httptest"
  "embed"
  golangRepository "web-golang/repository"
  "html/template"
  _"io"
)

//go:embed template/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "template/*.gohtml"))

func TemplateMenu(writer http.ResponseWriter, request *http.Request) {
  myTemplates.ExecuteTemplate(writer, "body", nil)
}

func TemplateAbout(writer http.ResponseWriter, request *http.Request) {
  myTemplates.ExecuteTemplate(writer, "about", nil)
}

func TemplateSiswa(writer http.ResponseWriter, request *http.Request) {
  db := golangRepository.NewSiswa(golangRepository.GetConnection())
  ctx := context.Background()

  siswa, err := db.FindAll(ctx)
  if err != nil {
    panic(err)
  }

  myTemplates.ExecuteTemplate(writer, "siswa", map[string]interface{}{
    "Siswa": siswa,
  })
}

func FormPost(writer http.ResponseWriter, request *http.Request)  {
  cek := request.Method
  if cek != "POST" {
    http.Redirect(writer, request, "/siswa", http.StatusTemporaryRedirect)
  }else {
    // parsing data
    err := request.ParseForm()
    if err != nil {
      panic(err)
    }

    // request.PostFormValue("firstname") cara singkat tanpa parsing

    name := request.PostForm.Get("name")
    alamat := request.PostForm.Get("alamat")
    jk := request.PostForm.Get("jk")

    user := entity.User{
      Name: name,
      Alamat: alamat,
      Jk: jk,
    }

    db := golangRepository.NewSiswa(golangRepository.GetConnection())
    ctx := context.Background()

    _, err = db.InsertUser(ctx, user)
    if err != nil {
      panic(err)
    }

    http.Redirect(writer, request, "/siswa", http.StatusTemporaryRedirect)
  }
}

func Tambah(writer http.ResponseWriter, request *http.Request) {
  myTemplates.ExecuteTemplate(writer, "tambahSiswa", nil)
}

func TestTemplatess(t *testing.T) {
  mux := http.NewServeMux()
  mux.HandleFunc("/", TemplateMenu)
  mux.HandleFunc("/about", TemplateAbout)
  mux.HandleFunc("/siswa", TemplateSiswa)
  mux.HandleFunc("/post", FormPost)
  mux.HandleFunc("/tambah", Tambah)

  server := http.Server{
    Addr: "localhost:8080",
    Handler: mux,
  }

  err := server.ListenAndServe()
  if err != nil {
    panic(err)
  }
}
