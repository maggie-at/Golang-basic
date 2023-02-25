package go_gorm

import "gorm.io/gorm"

// ManyToOne 多对一关系
func ManyToOne(db *gorm.DB) {
	// 定义多对一关系
	type Company struct { // One
		ID   int
		Name string
	}
	type Employee struct { // Many
		gorm.Model
		Name      string
		CompanyID int
		Company   Company
	}
	// 建表
	db.AutoMigrate(&Company{}, &Employee{})
	/*
		CREATE TABLE `employees` (
			`id` bigint unsigned NOT NULL AUTO_INCREMENT,
			`created_at` datetime(3) DEFAULT NULL,
			`updated_at` datetime(3) DEFAULT NULL,
			`deleted_at` datetime(3) DEFAULT NULL,
			`name` longtext,
			`company_id` bigint DEFAULT NULL,				// 外键名称
			PRIMARY KEY (`id`),
			KEY `idx_employees_deleted_at` (`deleted_at`),
			KEY `fk_employees_company` (`company_id`),		// 外键
			CONSTRAINT `fk_employees_company` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |
	*/
}

func ManyToOneRenamed(db *gorm.DB) {
	// 定义多对一关系, 重写外键名称
	type Company struct { // One
		ID   int
		Name string
	}
	type Employee struct { // Many
		gorm.Model
		Name         string
		CompanyRefer int
		Company      Company `gorm:"foreignKey:CompanyRefer"`
	}
	db.AutoMigrate(&Company{}, &Employee{})
	/*
		CREATE TABLE `employees` (
		  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
		  `created_at` datetime(3) DEFAULT NULL,
		  `updated_at` datetime(3) DEFAULT NULL,
		  `deleted_at` datetime(3) DEFAULT NULL,
		  `name` longtext,
		  `company_refer` bigint DEFAULT NULL,				// 外键名称
		  PRIMARY KEY (`id`),
		  KEY `idx_employees_deleted_at` (`deleted_at`),
		  KEY `fk_employees_company` (`company_refer`),		// 外键
		  CONSTRAINT `fk_employees_company` FOREIGN KEY (`company_refer`) REFERENCES `companies` (`id`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |
	*/
}

func ManyToOneRewrite(db *gorm.DB) {
	// 定义多对一关系
	type Company struct { // One
		ID   int
		Code string // 想要的外键
		Name string
	}
	type Employee struct { // Many
		gorm.Model
		Name      string
		CompanyID string
		Company   Company `gorm:"references:Code"` // 注意这里不是foreignKey
	}
	db.AutoMigrate(&Company{}, &Employee{})
}
