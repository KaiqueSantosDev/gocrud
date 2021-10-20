package price

import (
	"gorm.io/gorm"
)

type UseCase interface {
	GetAll() ([]*Product, error)
	Get(ID int64) (*Product, error)
	Create(p *Product) error
	Update(p *Product) error
	Delete(ID int64) error
}

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() ([]*Product, error) {
	p := []*Product{}
	err := s.DB.Debug().Model(&Product{}).Find(&p).Error
	if err != nil {
		return []*Product{}, err
	}
	return p, err
}

func (s *Service) Get(ID int64) (*Product, error) {
	p := &Product{}
	err := s.DB.Debug().Model(Product{}).Where("id = ?", ID).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}

	return &Product{p.ID, p.Name, p.Description, p.Price}, nil

}

func (s *Service) Create(p *Product) error {
	err := s.DB.Debug().Create(&p).Error
	if err != nil {
		return err
	}
	return nil

}

func (s *Service) Update(p *Product) error {
	err := s.DB.Debug().Model(&Product{}).Where("id = ?", p.ID).Updates(Product{p.ID, p.Name, p.Description, p.Price}).Error
	if err != nil {
		return err
	}
	return nil

}

func (s *Service) Delete(ID int64) error {
	err := s.DB.Debug().Where("id = ?", ID).Delete(&Product{}).Error
	if err != nil {
		return err
	}
	return nil
}
