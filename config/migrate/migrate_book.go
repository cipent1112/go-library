package migrate

import "gorm.io/gorm"

func MigrateBook(db *gorm.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS book (
  		id varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
	  	category_id varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
	  	author_id varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
	  	author_name varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
	  	title varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
	  	description text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
	  	price double NOT NULL DEFAULT '0',
	  	stock int(20) NOT NULL DEFAULT '0',
	  	release_date date DEFAULT NULL,
	  	created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	  	updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	  	deleted_at DATETIME(3) DEFAULT NULL,
	  	PRIMARY KEY (id),
	  	UNIQUE KEY idx_book_id (id),
	  	KEY idx_book_category_id (category_id),
	  	KEY idx_book_author_id (author_id),
	  	CONSTRAINT fk_book_author FOREIGN KEY (author_id) REFERENCES author (id),
	  	CONSTRAINT fk_book_category FOREIGN KEY (category_id) REFERENCES category (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_general_ci`)
}
