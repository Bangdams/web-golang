package repository

import (
  "testing"
  "fmt"
  "context"
  _ "github.com/go-sql-driver/mysql"
  _"web-golang/entity"
  _"time"
)

// func TestCek(t *testing.T) {
//   divisi := entity.Divisi{
//     Id: 1,
//     Kelas: "A",
//     User: entity.User{
//       Name: "Sadam",
//       Alamat: "Tasik",
//       Jk: "L",
//     },
//   }
//
//   fmt.Println(divisi)
// }

func TestQuery(t *testing.T) {
  db := NewDivisi(GetConnection())
  ctx := context.Background()
  // divisi := entity.Divisi{
  //   Kelas: "Sofware Enginer",
  // }

  divisi, err := db.FindAnggotaDiv(ctx)
  if err != nil {
    panic(err)
  }
  fmt.Println(divisi)
}

// func TestDelete(t *testing.T) {
//   db := NewSiswa(GetConnection())
//   ctx := context.Background()
//   _, err := db.DeleteUser(ctx, 11)
//   if err != nil {
//     panic(err)
//   }
//   fmt.Println("Berhasil")
// }
//
// func TestUpdate(t *testing.T) {
//   db := NewSiswa(GetConnection())
//   ctx := context.Background()
//   user := entity.User{
//     Id: 8,
//     Name: "Agus",
//     Alamat: "Bandung",
//     Jk: "laki-laki",
//   }
//
//   _, err := db.UpdateUser(ctx, user)
//   if err != nil {
//     panic(err)
//   }
//
//   fmt.Println("Berhasil")
// }
//
// func TestEdit(t *testing.T) {
//   db := NewSiswa(GetConnection())
//   ctx := context.Background()
//
//   result, err := db.FindById(ctx, 9)
//   if err != nil {
//     panic(err)
//   }
//
//   fmt.Println(result)
// }
