package usecases

import (
	"fmt"
	"reflect"
	"test/gormtransactionerr/app/domains"
	"test/gormtransactionerr/app/repositories"

	"gorm.io/gorm"
)

type InterfaceDummyUsecase interface {
	WithTrx(*gorm.DB) InterfaceDummyUsecase

	FindAll() ([]domains.Dummy, error)
	FindById(id uint32) (domains.Dummy, error)
	FindByConditions(conditions map[string]interface{}) (domains.Dummy, error)
	FindByEmail(email string) (domains.Dummy, error)
	Create(model *domains.Dummy) error
	UpdateDummy(user_id uint32, email string) error
	Update(model *domains.Dummy) error
	Delete(id uint32) error
}

type dummyUseCase struct {
	db        *gorm.DB
	dummyRepo repositories.InterfaceDummyRepository
}

func NewDummyUseCase(db *gorm.DB, repo repositories.InterfaceDummyRepository) InterfaceDummyUsecase {
	return &dummyUseCase{
		db:        db,
		dummyRepo: repo,
	}
}

func (m *dummyUseCase) WithTrx(trxHandle *gorm.DB) InterfaceDummyUsecase {
	m.dummyRepo = m.dummyRepo.WithTrx(trxHandle)
	return m
}

func (m *dummyUseCase) FindAll() ([]domains.Dummy, error) {
	newssubscription_list, err := m.dummyRepo.FindAll()
	return newssubscription_list, err
}

func (m *dummyUseCase) FindById(id uint32) (domains.Dummy, error) {
	dummy, err := m.dummyRepo.FindById(id)
	return dummy, err
}

func (m *dummyUseCase) FindByConditions(conditions map[string]interface{}) (domains.Dummy, error) {
	dummy, err := m.dummyRepo.FindByConditions(conditions)
	return dummy, err
}

func (m *dummyUseCase) FindByEmail(email string) (domains.Dummy, error) {
	dummy, err := m.dummyRepo.FindByEmail(email)
	if (reflect.DeepEqual(dummy, domains.Dummy{})) {
		return dummy, domains.ErrNotFound
	}
	return dummy, err
}

func (m *dummyUseCase) Create(model *domains.Dummy) error {
	err := m.dummyRepo.Create(model)
	return err
}

func (m *dummyUseCase) UpdateDummy(user_id uint32, email string) error {
	tx := m.db.Begin()
	if err := tx.Error; err != nil {
		fmt.Println("transaction error;")
		return err
	}
	dummy, err := m.dummyRepo.WithTrx(tx).FindById(user_id)
	if err != nil {
		fmt.Println("Rollback: query error")
		tx.Rollback()
		return err
	}

	dummy.Email = email
	err = m.dummyRepo.WithTrx(tx).Update(&dummy)
	if err != nil {
		fmt.Println("Rollback: update error")
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		fmt.Println("Rollback: commit error")
		tx.Rollback()
		return err
	}
	fmt.Println("Commit")
	return nil
}

func (m *dummyUseCase) Update(model *domains.Dummy) error {
	tx := m.db.Begin()
	if err := tx.Error; err != nil {
		fmt.Println("tx.Error;")
		return err
	}
	// ....
	// other usecase code
	// ....
	err := m.dummyRepo.WithTrx(tx).Update(model)
	if err != nil {
		tx.Rollback()
		return err
	}
	// ....
	// other usecase code
	// ....
	err = tx.Commit().Error
	if err != nil {
		fmt.Println("Rollback: sync error 2")
		tx.Rollback()
		return err
	}
	fmt.Println("Commit")
	return err
}

func (m *dummyUseCase) Delete(id uint32) error {
	res, _ := m.dummyRepo.FindById(id)

	if (reflect.DeepEqual(res, domains.Dummy{})) {
		return domains.ErrNotFound
	}

	err := m.dummyRepo.Delete(id)
	return err
}
