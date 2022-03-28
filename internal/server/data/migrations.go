package data

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"github.com/infrahq/infra/internal/server/models"
)

func migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// rename grants.identity -> grants.subject
		{
			ID: "202203231621", // current date
			Migrate: func(tx *gorm.DB) error {
				// it's a good practice to copy any used structs inside the function,
				// so side-effects are prevented if the original struct changes

				if tx.Migrator().HasColumn(&models.Grant{}, "identity") {
					return tx.Migrator().RenameColumn(&models.Grant{}, "identity", "subject")
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().RenameColumn(&models.Grant{}, "subject", "identity")
			},
		},
		// next one here
	})

	m.InitSchema(automigrate)

	if err := m.Migrate(); err != nil {
		return err
	}

	// automigrate again, so that for simple things like adding db fields we don't necessarily need to do a migration
	return automigrate(db)
}

func automigrate(db *gorm.DB) error {
	tables := []interface{}{
		&models.User{},
		&models.Machine{},
		&models.Group{},
		&models.Grant{},
		&models.Provider{},
		&models.ProviderToken{},
		&models.Destination{},
		&models.AccessKey{},
		&models.Settings{},
		&models.EncryptionKey{},
		&models.TrustedCertificate{},
		&models.RootCertificate{},
		&models.Credential{},
	}

	for _, table := range tables {
		if err := db.AutoMigrate(table); err != nil {
			return err
		}
	}

	return nil
}