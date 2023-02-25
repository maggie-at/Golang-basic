package go_gorm

import "gorm.io/gorm"

// OneToOne 一对一关系: 一个slave必须对应一个host, host不一定有对应的slave
func OneToOne(db *gorm.DB) {
	type CreditCard struct { // Slave
		gorm.Model
		Number     string
		EmployeeID uint
	}
	type Employee struct { // Host
		gorm.Model
		CreditCard CreditCard
	}
	db.AutoMigrate(&Employee{}, &CreditCard{}) // 保证主从创建顺序

	/*
		CREATE TABLE `credit_cards` (
		  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
		  `created_at` datetime(3) DEFAULT NULL,
		  `updated_at` datetime(3) DEFAULT NULL,
		  `deleted_at` datetime(3) DEFAULT NULL,
		  `number` longtext,
		  `employee_id` bigint unsigned DEFAULT NULL,	// Host's ID
		  PRIMARY KEY (`id`),
		  KEY `idx_credit_cards_deleted_at` (`deleted_at`),
		  KEY `fk_employees_credit_card` (`employee_id`),	// 一对一外键
		  CONSTRAINT `fk_employees_credit_card` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |
	*/
}
