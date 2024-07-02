package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// dsn := "Server=hcdb.c5is28gamk81.ap-south-1.rds.amazonaws.com;;Database=SMART_OFFICE;User Id=hc;Password=Hps202132"
	dsn := "Server=103.83.146.30;;Database=OMHARI;User Id=hc;Password=Hps202132"

	// var dataMap = map[string]func(gorm.ColumnType) (dataType string){
	// 	"bigint": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "int64"
	// 	},
	// 	"int": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "int32"
	// 	},
	// 	"smallint": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "int16"
	// 	},
	// 	"tinyint": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "uint8"
	// 	},
	// 	"bit": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "bool"
	// 	},
	// 	"decimal": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "*big.Rat" // Or another type that can handle fixed-point numbers
	// 	},
	// 	"numeric": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "*big.Rat" // Same as decimal
	// 	},
	// 	"money": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "float64" // Consider precision for monetary types
	// 	},
	// 	"smallmoney": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "float64" // Consider precision for monetary types
	// 	},
	// 	"float": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "float64"
	// 	},
	// 	"real": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "float32"
	// 	},
	// 	"datetime": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "time.Time"
	// 	},
	// 	"smalldatetime": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "time.Time"
	// 	},
	// 	"char": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "string"
	// 	},
	// 	"varchar": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "string"
	// 	},
	// 	"text": func(columnType gorm.ColumnType) (dataType string) {
	// 		return "string"
	// 	},
	// 	// Add mappings for other data types as needed
	// }

	g := gen.NewGenerator((gen.Config{
		OutPath:           "../pkg/db/query", // output directory, default value is ./query
		Mode:              gen.WithoutContext,
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	}))

	// g.WithDataTypeMap(dataMap)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	g.UseDB(db)

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)

	// Execute the generator
	g.Execute()
}
