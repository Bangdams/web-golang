package repository

import (
  "web-golang/entity"
  "context"
  "database/sql"
)

type divisiRepositoryImpl struct {
  DB *sql.DB
}

func NewDivisi(db *sql.DB) DivisiRepository {
  return &divisiRepositoryImpl{DB: db}
}


func (repository *divisiRepositoryImpl) FindAll(ctx context.Context) ([]entity.Divisi, error) {
  sql := "SELECT kelas FROM divisi"
  rows, err := repository.DB.QueryContext(ctx, sql)
  if err != nil {
    panic(err)
  }

  defer rows.Close()
  var divisis []entity.Divisi
  for rows.Next() {
    divisi := entity.Divisi{}
    rows.Scan(&divisi.Kelas)
    divisis = append(divisis, divisi)
  }

  return divisis, nil
}

func (repository *divisiRepositoryImpl) FindAnggotaDiv(ctx context.Context) ([]entity.Divisi, error) {
  sql := "SELECT divisi.kelas, siswa.name FROM divisi JOIN siswa ON divisi.id=siswa.kelas"
  rows, err := repository.DB.QueryContext(ctx, sql)
  if err != nil {
    panic(err)
  }

  defer rows.Close()
  var divisis []entity.Divisi
  for rows.Next() {
    divisi := entity.Divisi{}
    rows.Scan(&divisi.Kelas, &divisi.Name)
    divisis = append(divisis, divisi)
  }

  return divisis, nil
}
