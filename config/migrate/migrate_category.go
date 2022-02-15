package migrate

import "gorm.io/gorm"

func MigrateCategory(db *gorm.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS category (
	  id varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
	  name varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
	  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	  updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	  deleted_at DATETIME(3) DEFAULT NULL,
	  PRIMARY KEY (id),
	  UNIQUE KEY idx_category_id (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_general_ci`)
}
