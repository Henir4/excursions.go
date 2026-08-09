package handler

import (
	"excursion.com/config"
	"gorm.io/gorm"
)

var (
  logger *config.Logger
  db *gorm.DB
)

func InitalizeHandler() {
  logger = config.GetLogger("handler")
  db = config.GetSQLite()
}