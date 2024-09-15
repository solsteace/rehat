package services

import (
	"github.com/solsteace/rest/models"
	"github.com/solsteace/rest/repositories"
)

type MotelManagement struct {
	repositories.Motel
	repositories.MotelAdmin
	repositories.Class
}

// Checks whether an admin have permission to certain motel
func (m MotelManagement) checkPermission(userId, motelId int64) error {
	if _, err := m.MotelAdmin.GetByUserAndMotelId(userId, motelId); err != nil {
		if _, ok := err.(*repositories.ErrRecordNotFound); ok {
			return &ErrNoResourcePermission{}
		}
		return err
	}
	return nil
}

// Inserts a record of motel and admin associated with it to the database
func (m MotelManagement) AddMotel(userId int64, motel *models.Motel) (models.MotelAdmin, error) {
	var admin models.MotelAdmin

	motelId, err := m.Motel.Save(*motel)
	if err != nil {
		return admin, err
	}

	admin = models.MotelAdmin{UserID: userId, MotelID: motelId}
	adminId, err := m.MotelAdmin.Save(admin)
	if err != nil {
		return admin, err
	}
	admin.AdminID = adminId

	return admin, nil
}

// Edits a motel record by saving passed `motel`
func (m MotelManagement) EditMotel(userId int64, motel *models.Motel) error {
	if _, err := m.Motel.GetById(motel.MotelID); err != nil {
		return err
	}

	if err := m.checkPermission(userId, motel.MotelID); err != nil {
		return err
	}

	if err := m.Motel.EditById(motel.MotelID, *motel); err != nil {
		return err
	}
	return nil
}

// Deletes a motel record with certain id
func (m MotelManagement) DeleteMotel(userId, motelId int64) error {
	if _, err := m.Motel.GetById(motelId); err != nil {
		return err
	}

	if err := m.checkPermission(userId, motelId); err != nil {
		return err
	}

	if err := m.Motel.DeleteById(motelId); err != nil {
		return err
	}
	return nil
}

func (m MotelManagement) AddRoom(userId int64, room models.Room) error {
	return nil
}

func (m MotelManagement) EditRoom(userId int64, room models.Room) error {
	return nil
}

func (m MotelManagement) DeleteRoom(userId, roomId int64) error {
	return nil
}
