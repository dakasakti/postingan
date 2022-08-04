package services

import (
	"strconv"

	"github.com/dakasakti/postingan/entities"
	tm "github.com/dakasakti/postingan/repositories/posttype"
)

type postTypeService struct {
	Tm tm.PostTypeModel
}

func NewPostTypeService(tm tm.PostTypeModel) *postTypeService {
	return &postTypeService{Tm: tm}
}

func (ts *postTypeService) Register(user_id uint, data entities.PostTypeRequest) error {
	dataPost := entities.PostType{
		UserID: user_id,
		Jenis:  data.Jenis,
		Type:   data.Type,
	}

	err := ts.Tm.Insert(dataPost)
	if err != nil {
		return err
	}

	return nil
}

func (ts *postTypeService) GetAll() ([]entities.PostType, error) {
	data, err := ts.Tm.Gets()
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	return data, nil
}

func (ts *postTypeService) CheckParamId(id string) (uint, error) {
	idConv, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(idConv), nil
}

func (ts *postTypeService) GetById(id uint) (entities.PostType, error) {
	data, err := ts.Tm.Get(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (ts *postTypeService) UpdateById(id uint, data entities.PostTypeUpdateRequest) error {
	dataPost := entities.PostType{
		Jenis: data.Jenis,
		Type:  data.Type,
	}

	err := ts.Tm.Update(id, dataPost)
	if err != nil {
		return err
	}

	return nil
}

func (ts *postTypeService) DeleteById(id uint) error {
	err := ts.Tm.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
