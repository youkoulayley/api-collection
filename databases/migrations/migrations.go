package migrations

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/inflection"
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/models"
)

// LaunchMigrations launch all migrations that you want
func LaunchMigrations() {
	log.Info("Migrations - start")
	bootstrap.Db().AutoMigrate(&models.Role{})
	bootstrap.Db().AutoMigrate(&models.User{})
	bootstrap.Db().AutoMigrate(&models.Channel{})

	log.Info("Migrations - add foreign keys")
	BelongsToFIndex(bootstrap.Db(), &models.User{}, &models.Role{})
	BelongsToFIndex(bootstrap.Db(), &models.Channel{}, &models.User{})
	Many2ManyFIndex(bootstrap.Db(), &models.Channel{}, &models.User{})

	log.Info("Migrations - end")
}

// BelongsToFIndex create the foreign key for a belongs to relation
func BelongsToFIndex(db *gorm.DB, parentModel interface{}, childModel interface{}) {
	// Get name of the model
	tableParentAccessor := ReduceModelToName(parentModel)
	tableChildAccessor := ReduceModelToName(childModel)

	// Set model name to plural
	tableParentName := inflection.Plural(tableParentAccessor)
	tableChildName := inflection.Plural(tableChildAccessor)

	db.Table(tableParentName).AddForeignKey(tableChildAccessor+"_id", tableChildName+"(id)", "CASCADE", "CASCADE")
}

// Many2ManyFIndex create the index for many to many relations
func Many2ManyFIndex(db *gorm.DB, parentModel interface{}, childModel interface{}) {
	// Get name of the model
	table1Accessor := ReduceModelToName(parentModel)
	table2Accessor := ReduceModelToName(childModel)

	// Set model name to plural
	table1Name := inflection.Plural(table1Accessor)
	table2Name := inflection.Plural(table2Accessor)

	// Join the tables
	joinTable := fmt.Sprintf("%s_%s", table1Accessor, table2Accessor)

	// Add foreign Key and unique index
	db.Table(joinTable).AddForeignKey(table1Accessor+"_id", table1Name+"(id)", "CASCADE", "CASCADE")
	db.Table(joinTable).AddForeignKey(table2Accessor+"_id", table2Name+"(id)", "CASCADE", "CASCADE")
	db.Table(joinTable).AddUniqueIndex(joinTable+"_unique", table1Accessor+"_id", table2Accessor+"_id")
}

// ReduceModelToName get the name of the model by its reference
func ReduceModelToName(model interface{}) string {
	value := reflect.ValueOf(model)
	if value.Kind() != reflect.Ptr {
		return ""
	}

	elem := value.Elem()
	t := elem.Type()
	rawName := t.Name()
	return gorm.ToDBName(rawName)
}
