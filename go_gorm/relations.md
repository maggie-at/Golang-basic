### 1. Many to One (Belong To)

> 比如员工与公司, 可以默认使用`CompanyID`作为外键, 可以重写外键名称, 或者定义其它列作为外键

```GO
// ManyToOne 多对一关系
func ManyToOne(db *gorm.DB) {
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

```

```GO
// 重写外键名称
type Company struct { // One
    ID   int
    Name string
}
type Employee struct { // Many
    gorm.Model
    Name            string
    CompanyRefer    int
    Company         Company `gorm:foreignKey:CompanyRefer`
}
db.AutoMigrate(&Company{}, &Employee{})
```

```GO
// 指定其它列作为外键
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
```


### 2. One to One (Has One)
```GO
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
```


### 3. One to Many (Has Many)
```GO
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
		  KEY `fk_employees_credit_cards` (`employee_id`),  // Host's ID
		  CONSTRAINT `fk_employees_credit_cards` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`)  // 外键: 标识属于谁
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |
	*/
}
```


### 4. Many To Many

```GO
func ManyToMany(db *gorm.DB) {
	type Language struct {
		gorm.Model
		Name string
	}
	type People struct {
		gorm.Model
		Languages []Language `gorm:"many2many:people_languages"`
	}
	//db.AutoMigrate(&People{}, &Language{})
	/*
		CREATE TABLE `people_languages` (
		  `people_id` bigint unsigned NOT NULL,
		  `language_id` bigint unsigned NOT NULL,
		  PRIMARY KEY (`people_id`,`language_id`),
		  KEY `fk_people_languages_language` (`language_id`),
		  CONSTRAINT `fk_people_languages_language` FOREIGN KEY (`language_id`) REFERENCES `languages` (`id`),
		  CONSTRAINT `fk_people_languages_people` FOREIGN KEY (`people_id`) REFERENCES `peoples` (`id`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |
	*/

	lEng := Language{
		Name: "English22",
	}
	lCn := Language{
		Name: "Chinese11",
	}
	p := People{
		Languages: []Language{lEng, lCn},
	}
	db.Create(&p)
}
```
