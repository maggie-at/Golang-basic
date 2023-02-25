package go_gorm

import "gorm.io/gorm"

func ManyToMany(db *gorm.DB) {
	type Language struct {
		gorm.Model
		Name string
	}
	type People struct {
		gorm.Model
		Languages []Language `gorm:"many2many:people_languages"`
	}
	db.AutoMigrate(&People{}, &Language{})
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
