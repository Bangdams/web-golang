package repository

import (
  "context"
  "web-golang/entity"
)

type SiswaRepository interface {
  FindAll(ctx context.Context) ([]entity.User, error)
  InsertUser(ctx context.Context, user entity.User) (entity.User, error)
  DeleteUser(ctx context.Context, id int32) (entity.User, error)
  UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
  FindById(ctx context.Context, id int32) (entity.User, error)
}

type AbsenRepository interface {
  InsertData(ctx context.Context, absen entity.Absen) (entity.Absen, error)
  FindById(ctx context.Context, kelas int32) ([]entity.User, error)
}

type DivisiRepository interface {
  FindAll(ctx context.Context) ([]entity.Divisi, error)
  FindAnggotaDiv(ctx context.Context) ([]entity.Divisi, error)
  // InsertDivisi(ctx context.Context, divisi entity.Divisi) (entity.Divisi, error)
  // UpdateDivisi(ctx context.Context, divisi entity.Divisi) (entity.Divisi, error)
}
