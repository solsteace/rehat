package services

import (
	"github.com/solsteace/rest/models"
	"github.com/solsteace/rest/repositories"
)

type MotelManagement struct {
	repositories.Motel
	repositories.MotelAdmin
}

// Inserts a record of motel and admin associated with it to the database. Upon
// successful insertion, both motel and admin would be supplied with its new id
func (m MotelManagement) AddMotel(
	motel *models.Motel,
	admin *models.MotelAdmin,
) error {
	// TODO: Apply transaction
	var newMotelId, newMotelAdminId int64

	newMotelId, err := m.Motel.Save(*motel)
	if err != nil {
		return err
	}

	motel.MotelID = newMotelId
	admin.MotelID = newMotelId
	newMotelAdminId, err = m.MotelAdmin.Save(*admin)
	if err != nil {
		return err
	}

	admin.AdminID = newMotelAdminId
	return nil
}

// Edits motel record with certain id
func (m MotelManagement) EditMotelById(id int64, motel *models.Motel) error {
	if err := m.Motel.EditById(id, *motel); err != nil {
		return err
	}

	return nil
}

// Deletes motel record with certain id
func (m MotelManagement) DeleteMotelById(id int64) error {
	// TODO: Apply transaction
	if err := m.MotelAdmin.DeleteByMotelId(id); err != nil {
		return err
	}

	if err := m.Motel.DeleteById(id); err != nil {
		return err
	}

	return nil
}
