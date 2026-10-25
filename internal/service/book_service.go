package service

import (
	"API/internal/model"
	"API/internal/store"
)

type Service struct{
	store store.Store
}

func New(s store.Store) *Service{
	return&Service{
		store: s,
	}
}

func(s *Service) ObtenTodosLosLibros()([]*model.Libro,error){
	return s.store.GetAll()
}

func (s *Service) ObtenLibroPorID(id int)(*model.Libro,error){
	return s.store.GetByID(id)
}

func (s *Service) CrearLibro(libro model.Libro)(*model.Libro,error){
return s.store.Create(&libro)
}

func (s *Service) UpdateAllLibro(id int, libro model.Libro) (*model.Libro,error){
	return s.store.Update(id,&libro)
}

func (s *Service) BorrarLibro(id int) error{
	return s.store.Delete(id)
}