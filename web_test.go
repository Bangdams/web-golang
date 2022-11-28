package web_golang

import (
  "github.com/julienschmidt/httprouter"
  "testing"
  "net/http"
  "fmt"
  repository "web-golang/repository"
  "web-golang/entity"
  "context"
  _ "github.com/go-sql-driver/mysql"
  "strconv"
  "strings"
  "time"
  "encoding/json"
  "html/template"
  "bytes"
)

func templateMenu(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
  myTemplates.ExecuteTemplate(writer, "body", nil)
}

func templateAbout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
  myTemplates.ExecuteTemplate(writer, "about", nil)
}

func templateTambah(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
  myTemplates.ExecuteTemplate(writer, "tambahSiswa", nil)
}

func templateSadam(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
  myTemplates.ExecuteTemplate(writer, "sadam", nil)
}

func formPost(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
  err := request.ParseForm()
  if err != nil {
    panic(err)
  }

  name := request.PostForm.Get("name")
  alamat := request.PostForm.Get("alamat")
  jk := request.PostForm.Get("jk")
  kelas := request.PostForm.Get("divisi")
  
  getKelas, _ := strconv.Atoi(kelas)
  id32 := int32(getKelas)

  user := entity.User{
    Name: name,
    Alamat: alamat,
    Jk: jk,
    Kelas: id32,
  }

  db := repository.NewSiswa(repository.GetConnection())
  ctx := context.Background()

  _, err = db.InsertUser(ctx, user)
  if err != nil {
    ResponseErr(writer, http.StatusInternalServerError, err.Error())
  }

  data := map[string]interface{}{
    "message": "Berhasil",
    "data": template.HTML(GetData()),
  }

  ResponseJson(writer, http.StatusOK, data)
  http.Redirect(writer, request, "/siswa", http.StatusFound)
}

func ResponseErr(writer http.ResponseWriter, code int, message string)  {
  ResponseJson(writer, code, map[string]string{"error": message})
}

func ResponseJson(writer http.ResponseWriter, code int, pyload interface{})  {
  response, _ := json.Marshal(pyload)
  writer.Header().Set("Content-Type", "application/json")
  writer.WriteHeader(code)
  writer.Write(response)
}

func DeleteSiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
  id := params.ByName("id")

  getId, _ := strconv.Atoi(id)
  id32 := int32(getId)

  db := repository.NewSiswa(repository.GetConnection())
  ctx := context.Background()
  _, err := db.DeleteUser(ctx, id32)
  if err != nil {
    panic(err)
  }

  http.Redirect(writer, request,  "/", http.StatusFound)
}

func EditSiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
  db := repository.NewSiswa(repository.GetConnection())
  ctx := context.Background()

  id := params.ByName("id")
  getId, _ := strconv.Atoi(id)
  id32 := int32(getId)

  result, err := db.FindById(ctx, id32)
  if err != nil {
    panic(err)
  }

  myTemplates.ExecuteTemplate(writer, "editSiswa", map[string]interface{}{
    "Data": result,
  })
}

func UpdateSiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
  err := request.ParseForm()
  if err != nil {
    panic(err)
  }

  id := request.PostForm.Get("id")
  name := request.PostForm.Get("name")
  alamat := request.PostForm.Get("alamat")
  jk := request.PostForm.Get("jk")

  getId, _ := strconv.Atoi(id)
  id32 := int32(getId)

  user := entity.User{
    Id: id32,
    Name: name,
    Alamat: alamat,
    Jk: jk,
  }

  db := repository.NewSiswa(repository.GetConnection())
  ctx := context.Background()

  _, err = db.UpdateUser(ctx, user)
  if err != nil {
    panic(err)
  }

  http.Redirect(writer, request,  "/about", http.StatusFound)
}

func templateAbsen(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
  db := repository.NewSiswa(repository.GetConnection())
  ctx := context.Background()

  siswa, err := db.FindAll(ctx)
  if err != nil {
    panic(err)
  }

  now := time.Now()
  myTemplates.ExecuteTemplate(writer, "absen", map[string]interface{}{
    "Siswa": siswa,
    "Tanggal": now.Format("Jan 2, 2006"),
  })
}

func templateSiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
  myTemplates.ExecuteTemplate(writer,"siswa", map[string]interface{}{
    "Data": template.HTML(GetData()),
  })
}

func GetData() string {
  buffer := &bytes.Buffer{}
  db := repository.NewSiswa(repository.GetConnection())
  ctx := context.Background()

  siswa, err := db.FindAll(ctx)
  if err != nil {
    panic(err)
  }

  myTemplates.ExecuteTemplate(buffer,"siswaData", map[string]interface{}{
    "Siswa": siswa,
  })

  return buffer.String()
}

func absenPost(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
  err := request.ParseForm()
  if err != nil {
    panic(err)
  }

  cek := request.Form
  for key, _ := range cek {
    trim := strings.Trim(key, "status[ ]")

    getId, _ := strconv.Atoi(trim)
    id32 := int32(getId)

    now := time.Now()
    status := request.PostForm.Get(key)

    absen := entity.Absen{
      IdSiswa: id32,
      Status: status,
      Date: now.String(),
    }

    db := repository.NewAbsen(repository.GetConnection())
    ctx := context.Background()

    _, err = db.InsertData(ctx, absen)
    if err != nil {
      panic(err)
    }
  }
  return
  fmt.Fprint(writer, cek)
}

func templateDivisi(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
  db := repository.NewDivisi(repository.GetConnection())
  dbS := repository.NewSiswa(repository.GetConnection())
  ctx := context.Background()

  divisi, err := db.FindAll(ctx)
  siswa, err := dbS.FindAll(ctx)

  if err != nil {
    panic(err)
  }

  myTemplates.ExecuteTemplate(writer, "devisi", map[string]interface{}{
    "Kelas": divisi,
    "Siswa": siswa,
  })
}

func FindAbsen(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
  db := repository.NewAbsen(repository.GetConnection())
  ctx := context.Background()

  id := params.ByName("id")
  getId, _ := strconv.Atoi(id)
  id32 := int32(getId)

  result, err := db.FindById(ctx, id32)
  if err != nil {
    panic(err)
  }

  now := time.Now()
  myTemplates.ExecuteTemplate(writer, "absen", map[string]interface{}{
    "Siswa": result,
    "Tanggal": now.Format("Jan 2, 2006"),
  })
}

func cek(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
  db := repository.NewDivisi(repository.GetConnection())
  
  ctx := context.Background()
  divisi, err := db.FindAll(ctx)
  if err != nil {
    panic(err)
  }

  myTemplates.ExecuteTemplate(writer, "modal", map[string]interface{}{
    "Divisi": divisi,
  })
}

func rr(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
  myTemplates.ExecuteTemplate(writer, "Test", nil)
}


func TestRouterR(t *testing.T) {

  router := httprouter.New()
  // Custom method Method Not Allowed
  // router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request)  {
  //   http.Redirect(writer, request, "/siswa", http.StatusTemporaryRedirect)
  // })

  router.NotFound = http.StripPrefix("/static", http.FileServer(http.Dir("assets")))
  router.GET("/", templateMenu)
  router.GET("/sadam", templateSadam)
  router.GET("/rr", rr)
  router.GET("/about", templateAbout)
  router.GET("/siswa", templateSiswa)
  router.GET("/tambah", templateTambah)
  router.POST("/post", formPost)
  router.POST("/postt", UpdateSiswa)
  router.GET("/delete/:id", DeleteSiswa)
  router.GET("/edit/:id", EditSiswa)
  // router.GET("/absen", templateAbsen)
  router.GET("/absenCek/:id", FindAbsen)
  router.POST("/absenPost", absenPost)
  router.GET("/divisi", templateDivisi)
  router.GET("/cek", cek)

  server := http.Server{
    Handler: router,
    Addr: "localhost:3030",
  }

  err := server.ListenAndServe()
  if err != nil {
    panic(err)
  }
}
