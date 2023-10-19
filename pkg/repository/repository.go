package repository

type Authrization interface{

}

type DutyList interface{

}

type Repository struct{
	Authrization
	DutyList
}

func NewRepository() *Repository{
	return &Repository{}
}