package repository

import (
  "web-golang/entity"
  "context"
  "database/sql"
  _"errors"
)

type absenRepositoryImp struct {
  DB *sql.DB
}

func NewAbsen(db *sql.DB) AbsenRepository {
  return &absenRepositoryImp{DB: db}
}

func (repository *absenRepositoryImp) InsertData(ctx context.Context, absen entity.Absen) (entity.Absen, error) {
  data := "INSERT INTO absen(idSiswa, status, Tanggal) VALUES(? ,?, ?)"
  result, err := repository.DB.ExecContext(ctx, data, absen.IdSiswa, absen.Status, absen.Date)
  if err != nil {
    return absen, err
  }

  // AI Buat Id
  id, err := result.LastInsertId()
  if err != nil {
    return absen, err
  }

  absen.Id = int32(id)

  return absen, nil
}

func (repository *absenRepositoryImp) FindById(ctx context.Context, kelas int32) ([]entity.User, error) {
  data := "SELECT id, name, alamat, jk FROM siswa WHERE kelas=?"
  rows, err := repository.DB.QueryContext(ctx, data, kelas)
  if err != nil {
    panic(err)
  }

  defer rows.Close()
  var users []entity.User

  for rows.Next() {
    user := entity.User{}
    rows.Scan(&user.Id, &user.Name, &user.Alamat, &user.Jk)
    users = append(users, user)
  }

  // if rows.Next() {
  //   user := entity.User{}
  //   rows.Scan(&user.Id, &user.Name, &user.Alamat, &user.Jk)
  //   users = append(users, user)
  // }else {
  //   return users, errors.New("Data tidak ada")
  // }
  return users, nil
}
