package category

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	Create(ctx context.Context, request CreateRequest) (Response, error)
	Update(ctx context.Context, request UpdateRequest) (Response, error)
	Delete(ctx context.Context, categoryId int) error
	FindById(ctx context.Context, categoryId int) (Response, error)
	FindAll(ctx context.Context) ([]Response, error)
}

type service struct {
	Repository Repository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewService(repository Repository, db *sql.DB, validate *validator.Validate) Service {
	return &service{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *service) Create(ctx context.Context, request CreateRequest) (resp Response, err error) {
	err = service.Validate.Struct(request)
	if err != nil {
		return
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	category := Category{
		Name: request.Name,
	}

	category, err = service.Repository.Save(ctx, tx, category)
	if err != nil {
		return
	}

	resp = ToResponse(category)
	return
}

func (service *service) Update(ctx context.Context, request UpdateRequest) (resp Response, err error) {
	err = service.Validate.Struct(request)
	if err != nil {
		return
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
		err = tx.Commit()
	}()

	category, err := service.Repository.FindById(ctx, tx, request.Id)
	if err != nil {
		return
	}

	category.Name = request.Name

	category, err = service.Repository.Update(ctx, tx, category)
	if err != nil {
		return
	}

	return ToResponse(category), err
}

func (service *service) Delete(ctx context.Context, categoryId int) (err error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
		err = tx.Commit()
	}()

	category, err := service.Repository.FindById(ctx, tx, categoryId)
	if err != nil {
		return
	}

	err = service.Repository.Delete(ctx, tx, category)
	if err != nil {
		return
	}

	return
}

func (service *service) FindById(ctx context.Context, categoryId int) (resp Response, err error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
		err = tx.Commit()
	}()

	category, err := service.Repository.FindById(ctx, tx, categoryId)
	if err != nil {
		return
	}

	resp = ToResponse(category)
	return
}

func (service *service) FindAll(ctx context.Context) (resp []Response, err error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	categoryList, err := service.Repository.FindAll(ctx, tx)
	if err != nil {
		return
	}

	resp = ToResponseList(categoryList)
	return
}
