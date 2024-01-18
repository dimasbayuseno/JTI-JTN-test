package sampleservice

import (
	"context"
	"fmt"
	"initial/domain/sample/samplemodel"
	"initial/domain/sample/samplerepository"
	"initial/infrastructure/shared/response"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateData(ctx context.Context, model samplemodel.SampleDataCreateModel) (res samplemodel.SampleDataCreateModel, err error)
	UpdateData(ctx context.Context, faqListModel samplemodel.SampleDataUpdateModel, id string) (res samplemodel.SampleDataUpdateModel, err error)
	DeleteData(ctx context.Context, id string) (err error)
	GetDataById(ctx context.Context, id string) (res samplemodel.SampleDataModel, err error)
	GetAllData(ctx context.Context) (res []samplemodel.SampleDataModel, err error)
}

type service struct {
	sampleRepository samplerepository.Repository
}

func New(sampleRepository samplerepository.Repository) *service {
	return &service{sampleRepository: sampleRepository}
	//return &service{repo: *repo, Cache: cache}
}

func (s *service) CreateData(ctx context.Context, sampleDataListModel samplemodel.SampleDataCreateModel) (res samplemodel.SampleDataCreateModel, err error) {
	err = response.Validate(sampleDataListModel)
	if err != nil {
		err = fmt.Errorf("sampleservice: error validating request, %w", err)
		return res, err
	}
	sampleDataList := samplemodel.Sample{
		CreatedAt: time.Now(),
	}
	_, err = s.sampleRepository.InsertData(ctx, sampleDataList)
	if err != nil {
		err = fmt.Errorf("sampleservice: error inserting data, %w", err)
		return res, err
	}

	sampleDataListModel = samplemodel.SampleDataCreateModel{
		ID:        sampleDataList.ID,
		CreatedAt: sampleDataList.CreatedAt,
	}
	return sampleDataListModel, nil
}

func (s *service) UpdateData(ctx context.Context, sampleDataListModel samplemodel.SampleDataUpdateModel, id string) (res samplemodel.SampleDataUpdateModel, err error) {
	err = response.Validate(sampleDataListModel)
	if err != nil {
		err = fmt.Errorf("sampleservice: error validating request, %w", err)
		return res, err
	}

	_, err = s.sampleRepository.GetDataById(ctx, id)
	if err != nil {
		err = fmt.Errorf("sampleservice: error getting data, %w", err)
		return res, err
	}

	sampleDataList := samplemodel.Sample{
		ID:        uuid.MustParse(id),
		UpdatedAt: time.Now(),
	}

	result, err := s.sampleRepository.UpdateData(ctx, sampleDataList)
	res = samplemodel.SampleDataUpdateModel{
		ID:        result.ID,
		UpdatedAt: result.UpdatedAt,
	}
	return res, nil
}

func (s *service) DeleteData(ctx context.Context, id string) (err error) {
	sampleDataList, err := s.sampleRepository.GetDataById(ctx, id)
	if err != nil {
		err = fmt.Errorf("sampleservice: error getting data, %w", err)
		return err
	}
	err = s.sampleRepository.DeleteData(ctx, sampleDataList)
	if err != nil {
		err = fmt.Errorf("sampleservice: error deleting data, %w", err)
		return err
	}

	return nil
}

func (s *service) GetDataById(ctx context.Context, id string) (res samplemodel.SampleDataModel, err error) {
	getById, err := s.sampleRepository.GetDataById(ctx, id)
	if err != nil {
		err = fmt.Errorf("sampleservice: error getting data, %w", err)
		return res, err
	}
	return samplemodel.SampleDataModel{
		ID:        getById.ID,
		CreatedAt: getById.CreatedAt,
		UpdatedAt: getById.UpdatedAt,
	}, nil
}

func (s *service) GetAllData(ctx context.Context) (responses []samplemodel.SampleDataModel, err error) {
	sampleDataList, err := s.sampleRepository.GetAllData(ctx)
	for _, sampleDataLists := range sampleDataList {
		responses = append(responses, samplemodel.SampleDataModel{
			ID:        sampleDataLists.ID,
			CreatedAt: sampleDataLists.CreatedAt,
			UpdatedAt: sampleDataLists.UpdatedAt,
		})
	}
	if len(sampleDataList) == 0 {
		return []samplemodel.SampleDataModel{}, nil
	}
	return responses, nil
}

func (s *service) TxExample(ctx context.Context) (err error) {
	tx, _ := s.sampleRepository.Begin(ctx)
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	// do something

	return nil
}
