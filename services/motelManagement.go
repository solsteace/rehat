package services

import (
	"github.com/solsteace/rest/models"
	"github.com/solsteace/rest/repositories"
)

type MotelManagement struct {
	repositories.Motel
	repositories.MotelAdmin
}

// Checks whether an admin have permission to certain motel
func (m MotelManagement) checkPermission(UserId, motelId int64) error {
	if _, err := m.MotelAdmin.GetByUserAndMotelId(UserId, motelId); err != nil {
		if _, ok := err.(*repositories.ErrRecordNotFound); ok {
			return &ErrNoResourcePermission{}
		}
		return err
	}
	return nil
}

// Inserts a record of motel and admin associated with it to the database. Upon
// successful insertion, both motel and admin would be supplied with its new id
func (m MotelManagement) Add(admin *models.MotelAdmin, motel *models.Motel) error {
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
func (m MotelManagement) EditById(userId, motelId int64, motel *models.Motel) error {
	if err := m.checkPermission(userId, motelId); err != nil {
		return err
	}

	if err := m.Motel.EditById(motelId, *motel); err != nil {
		return err
	}
	return nil
}

// Deletes motel record with certain id
func (m MotelManagement) DeleteById(userId, motelId int64) error {
	if err := m.checkPermission(userId, motelId); err != nil {
		return err
	}

	// TODO: Apply transaction
	if err := m.MotelAdmin.DeleteByMotelId(motelId); err != nil {
		return err
	}

	if err := m.Motel.DeleteById(motelId); err != nil {
		return err
	}
	return nil
}
