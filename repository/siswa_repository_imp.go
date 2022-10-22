package repository

import (
  "web-golang/entity"
  "context"
  "database/sql"
  "errors"
)

type siswaRepositoryImp struct {
  DB *sql.DB
}

func NewSiswa(db *sql.DB) SiswaRepository {
  return &siswaRepositoryImp{DB: db}
}

func (repository *siswaRepositoryImp) FindAll(ctx context.Context) ([]entity.User, error) {
  data := "SELECT id, name, alamat, jk FROM siswa"
  rows, err := repository.DB.QueryContext(ctx, data)
  if err != nil {
    panic(err)
  }

  defer rows.Close()
  var siswas []entity.User
  for rows.Next() {
    siswa := entity.User{}
    rows.Scan(&siswa.Id, &siswa.Name, &siswa.Alamat, &siswa.Jk)
    siswas = append(siswas, siswa)
  }

  return siswas, nil
}

func (repository *siswaRepositoryImp) InsertUser(ctx context.Context, user entity.User) (entity.User, error) {
  data := "INSERT INTO siswa(name, alamat, jk) VALUES(? ,? ,?)"
  result, err := repository.DB.ExecContext(ctx, data, user.Name, user.Alamat, user.Jk)
  if err != nil {
    return user, err
  }

  // AI Buat Id
  id, err := result.LastInsertId()
  if err != nil {
    return user, err
  }

  user.Id = int32(id)

  return user, nil
}

func (repository *siswaRepositoryImp) DeleteUser(ctx context.Context, id int32) (entity.User, error) {
  data := "SELECT id FROM siswa WHERE id=? LIMIT 1"
  rows, err := repository.DB.QueryContext(ctx, data, id)
  siswa := entity.User{}
  if err != nil {
    return siswa, err
  }
  defer rows.Close()

  sql := "DELETE FROM siswa WHERE id=?"
  _, err = repository.DB.ExecContext(ctx, sql, id)
  if err != nil {
    return siswa, err
  }

  return siswa, nil
}

func (repository *siswaRepositoryImp) FindById(ctx context.Context, id int32) (entity.User, error) {
  data := "SELECT id, name, alamat, jk FROM siswa WHERE id=?"
  rows, err := repository.DB.QueryContext(ctx, data, id)
  user := entity.User{}
  if err != nil {
    return user, err
  }

  defer rows.Close()
  if rows.Next() {
    rows.Scan(&user.Id, &user.Name, &user.Alamat, &user.Jk)
    return user, nil
  }else {
    return user, errors.New("Data tidak ada")
  }
}

func (repository *siswaRepositoryImp) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
  data := "UPDATE siswa SET name=?, alamat=?, jk=? WHERE id=?"
  _, err := repository.DB.ExecContext(ctx, data, user.Name, user.Alamat, user.Jk, user.Id)
  if err != nil {
    return user, err
  }

  return user, nil
}
