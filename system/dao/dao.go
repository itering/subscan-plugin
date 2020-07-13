package dao

import (
	"github.com/itering/subscan-plugin/system/model"
	"github.com/itering/subscan-plugin/tools"
	"github.com/itering/substrate-api-rpc"
	"github.com/itering/substrate-api-rpc/metadata"
	"github.com/jinzhu/gorm"
	"strings"
)

func CreateExtrinsicError(db *gorm.DB, hash string, moduleError *substrate.MetadataModuleError) error {
	if moduleError == nil {
		return nil
	}
	query := db.Create(&model.ExtrinsicError{
		ExtrinsicHash: tools.AddHex(hash),
		Module:        moduleError.Module,
		Name:          moduleError.Name,
		Doc:           strings.Join(moduleError.Doc, ","),
	})
	return query.Error
}

func ExtrinsicError(db *gorm.DB, hash string) *model.ExtrinsicError {
	var e model.ExtrinsicError
	db.Where("extrinsic_hash = ?", hash).Find(&e)
	return &e
}

func CheckExtrinsicError(spec int, raw string, moduleIndex, errorIndex int) *substrate.MetadataModuleError {

	modules := metadata.Process(&metadata.RuntimeRaw{Raw: raw, Spec: spec})

	if moduleIndex >= len(modules.Metadata.Modules) {
		return nil
	}

	module := modules.Metadata.Modules[moduleIndex]
	if errorIndex >= len(module.Errors) {
		return nil
	}

	err := module.Errors[errorIndex]
	return &substrate.MetadataModuleError{
		Module: module.Name,
		Name:   err.Name,
		Doc:    err.Doc,
	}
}
