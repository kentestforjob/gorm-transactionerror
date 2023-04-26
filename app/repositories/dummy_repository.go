package repositories

import (
	"log"
	"test/gormtransactionerr/app/domains"

	"gorm.io/gorm"
)

type InterfaceDummyRepository interface {
	WithTrx(*gorm.DB) InterfaceDummyRepository

	FindAll() ([]domains.Dummy, error)
	FindById(id uint32) (domains.Dummy, error)
	FindByConditions(conditions map[string]interface{}) (domains.Dummy, error)
	FindByEmail(email string) (domains.Dummy, error)
	Create(model *domains.Dummy) error
	BatchCreate(models *[]domains.Dummy) error
	Update(model *domains.Dummy) error
	UpdateModel(model *domains.Dummy) error
	Delete(id uint32) error
}

type dummyRepository struct {
	Conn *gorm.DB
}

// GORM - delcare repository
func NewDummy(conn *gorm.DB) InterfaceDummyRepository {
	return &dummyRepository{
		Conn: conn,
	}
}

func (m *dummyRepository) WithTrx(trxHandle *gorm.DB) InterfaceDummyRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return m
	}
	m.Conn = trxHandle
	return m
}

func (m *dummyRepository) FindAll() ([]domains.Dummy, error) {
	var dummy_list []domains.Dummy
	err := m.Conn.Find(&dummy_list).Error

	return dummy_list, err
}

func (m *dummyRepository) FindById(id uint32) (domains.Dummy, error) {
	var dummy domains.Dummy
	err := m.Conn.First(&dummy, id).Error

	return dummy, err
}

func (m *dummyRepository) FindByConditions(conditions map[string]interface{}) (domains.Dummy, error) {
	dummy := domains.Dummy{}
	error := m.Conn.Model(&domains.Dummy{}).Where(conditions).First(&dummy).Error

	return dummy, error
}

func (m *dummyRepository) FindByEmail(email string) (domains.Dummy, error) {
	var dummy domains.Dummy
	err := m.Conn.First(&dummy, "email = ? ", email).Error
	return dummy, err
}

func (m *dummyRepository) Create(model *domains.Dummy) error {
	err := m.Conn.Create(&model).Error
	return err
}
func (m *dummyRepository) BatchCreate(models *[]domains.Dummy) error {
	err := m.Conn.Create(&models).Error
	return err
}

func (m *dummyRepository) Update(model *domains.Dummy) error {
	err := m.Conn.Omit("id").Updates(&model).Error
	return err
}

func (m *dummyRepository) UpdateModel(model *domains.Dummy) error {
	err := m.Conn.Debug().Model(&model).Omit("id").Updates(&model).Error
	return err
}

func (m *dummyRepository) Delete(id uint32) error {
	err := m.Conn.Delete(&domains.Dummy{}, id).Error
	return err
}
