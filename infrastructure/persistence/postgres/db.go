package postgres

import (
	"chopp-reitistom-backend/config"
	"chopp-reitistom-backend/domain/repository"
	"chopp-reitistom-backend/infrastructure/persistence/model"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repositories struct {
	User       repository.UserRepositoryInterface
	Address    repository.AddressRepositoryInterface
	Suggestion repository.SuggestionRepositoryInterface
	Category   repository.CategoryRepositoryInterface
	Rating     repository.RatingRepositoryInterface
	Product    repository.ProductRepositoryInterface
	Ingredient repository.IngredientRepositoryInterface
	db         *gorm.DB
	sqlDB      *sql.DB
}

func NewDB(config config.DB) (*Repositories, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)

	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", config.Host, config.Port, config.User, config.Name, config.Password)
	dbPostgres, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := dbPostgres.DB()

	return &Repositories{
		User:       NewUserRepository(dbPostgres),
		Address:    NewAddressRepository(dbPostgres),
		Suggestion: NewSuggestionRepository(dbPostgres),
		Category:   NewCategoryRepository(dbPostgres),
		Rating:     NewRatingRepository(dbPostgres),
		Product:    NewProductRepository(dbPostgres),
		Ingredient: NewIngredientRepository(dbPostgres),
		db:         dbPostgres,
		sqlDB:      sqlDB,
	}, nil
}

func (r *Repositories) Close() error {
	return r.sqlDB.Close()
}

func (r *Repositories) Automigrate() error {
	return r.db.AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.Suggestion{},
		&model.Category{},
		&model.Rating{},
		&model.Product{},
		&model.Ingredient{},
	)
}
