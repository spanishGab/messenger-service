package repositories

import (
	"encoding/json"
	"fmt"
	"spanishGab/aula_camada_model/src/db"
	"spanishGab/aula_camada_model/src/models"
	"time"

	_ "embed"

	"github.com/google/uuid"
)

const dateLayout = "2006-01-02"

type PersonDBRegistry struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Document  string    `json:"document"`
	BirthDate string    `json:"birth_date"`
}

func (p *PersonDBRegistry) ToModel() *models.Person {
	birthDate, _ := time.Parse(dateLayout, p.BirthDate)
	return &models.Person{
		Id:        p.Id,
		Name:      p.Name,
		Document:  p.Document,
		BirthDate: birthDate,
	}
}

// pagination.go
// package repositories
type Pagination struct {
	Limit uint8 `json:"limit"`
	Page  uint8 `json:"page"`
}

func (p *Pagination) Paginate(items any) any {
	return nil
}

// END

type PaginatedPersons struct {
	Content    []*models.Person `json:"content"`
	Pagination `json:"pagination"`
}

type PersonRepository struct {
	dbConnection db.FileHandler
}

func NewPersonRepository(dbConnection db.FileHandler) *PersonRepository {
	return &PersonRepository{
		dbConnection: dbConnection,
	}
}

func (p *PersonRepository) GetById(id uuid.UUID) (*models.Person, error) {
	var persons []PersonDBRegistry

	personsDBTable, err := p.dbConnection.Read()
	if err != nil {
		fmt.Printf("Error on PersonRepository.getById: %s\n", err)
		return nil, err
	}
	err = json.Unmarshal(personsDBTable, &persons)

	if err != nil {
		fmt.Printf("Error on PersonRepository.getById: %s\n", err)
		return nil, err
	}
	for _, person := range persons {
		if person.Id == id {
			return person.ToModel(), nil
		}
	}
	return nil, fmt.Errorf("person with id '%s' not found", id)
}

func (p *PersonRepository) GetAll(limit uint8, offset uint8) (*PaginatedPersons, error) {
	var persons []PersonDBRegistry

	personsDBTable, err := p.dbConnection.Read()
	if err != nil {
		fmt.Printf("Error on PersonRepository.GetAll: %s\n", err)
		return nil, err
	}
	err = json.Unmarshal(personsDBTable, &persons)

	if err != nil {
		fmt.Printf("Error on PersonRepository.GetAll: %s\n", err)
		return nil, err
	}

	start := int(limit * offset)
	end := start + int(limit)
	var personModels []*models.Person
	for i := start; i < end && i < len(persons); i++ {
		personModels = append(personModels, persons[i].ToModel())
	}

	pagination := Pagination{
		Limit: limit,
		Page:  offset,
	}
	a := &PaginatedPersons{
		Content:    pagination.Paginate(personModels).([]*models.Person),
		Pagination: pagination,
	}
	return a, nil
}
