package repository

import (
  "web-golang/entity"
  "context"
  "database/sql"
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
