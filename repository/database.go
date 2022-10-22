package repository

import (
  "database/sql"
  "time"
)

func GetConnection() *sql.DB  {
  db, err := sql.Open("mysql", "root@tcp(localhost:3306)/siswa_golang?parseTime=true")
  if err != nil {
    panic(err)
  }

  // Pooling
  db.SetMaxIdleConns(10)
  db.SetMaxOpenConns(100)
  db.SetConnMaxIdleTime(5 * time.Minute)
  db.SetConnMaxLifetime(60 * time.Minute)

  return db
}
