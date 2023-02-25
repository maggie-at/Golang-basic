package go_gorm

import "gorm.io/gorm"

// OneToMany 一个Host可以拥有多个Slave, 一个Slave必须属于一个Host
func OneToMany(db *gorm.DB) {
	type CreditCard struct { // Slave
		gorm.Model
		Number     string
		EmployeeID uint
	}
	type Employee struct { // Host
		gorm.Model
		CreditCards []CreditCard
	}
	db.AutoMigrate(&Employee{}, &CreditCard{})
	/*
		CREATE TABLE `credit_cards` (
		  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
		  `created_at` datetime(3) DEFAULT NULL,
		  `updated_at` datetime(3) DEFAULT NULL,
		  `deleted_at` datetime(3) DEFAULT NULL,
		  `number` longtext,
		  `employee_id` bigint unsigned DEFAULT NULL,
		  PRIMARY KEY (`id`),
		  KEY `idx_credit_cards_deleted_at` (`deleted_at`),
		  KEY `fk_employees_credit_cards` (`employee_id`),
		  CONSTRAINT `fk_employees_credit_cards` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |
	*/
}
