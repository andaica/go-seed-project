package home

import (
	"errors"

	log "../../logger"
	"../../ormbase"
	"github.com/jinzhu/gorm"
)

type PageModel struct {
	ID    uint    `gorm:"primary_key;auto_increment" json:"id"`
	Name  string  `gorm:"column:name;unique" json:"name"`
	Link  string  `gorm:"column:link;size:255" json:"link"`
	Image *string `gorm:"column:image" json:"image"`
}

func (PageModel) TableName() string {
	return "page"
}

var db *gorm.DB
var connectErr error

// auto run, initialize db connection
func init() {
	db, connectErr = ormbase.Connect("root", "anmap1234", "@tcp(127.0.0.1:3306)", "demoAPI")
	if connectErr != nil {
		log.Error(connectErr)
	}
	log.Log("Connect DB success!")

	if db.HasTable(&PageModel{}) == false {
		db.AutoMigrate(&PageModel{})
		log.Log("Created table page!")
	}
}

func allPages() (all []PageModel, err error) {
	errs := db.Find(&all).GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return
	}

	log.Log("all Pages: ", all)
	return
}

func insertPage(page PageModel) (newPage PageModel, err error) {
	var checkPage PageModel
	db.Where("name = ?", page.Name).Find(&checkPage)
	if checkPage.ID > 0 {
		err = errors.New("Page already exitst!")
		return
	}

	errs := db.Create(&page).GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return
	}

	log.Log("created Page: ", page)
	newPage = page
	return
}

func deletePage(page PageModel) (result bool, err error) {
	errs := db.Where("id = ?", page.ID).Find(&page).GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return false, err
	}

	errs = db.Delete(&page).GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return false, err
	}

	log.Log("deleted Page: ", page)
	return true, err
}

func updatePage(newPage PageModel) (page PageModel, err error) {
	errs := db.Where("id = ?", newPage.ID).Find(&page).GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return page, err
	}

	errs = db.Model(&page).Updates(newPage).GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return
	}

	log.Log("updated Page: ", page)
	return
}
