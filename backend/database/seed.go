package database

import (
	"fmt"
	"reflect"
	"time"

	"github.com/TimiBolu/lema-ai-users-service/models"
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to count users: %w", err)
	}
	if count > 0 {
		fmt.Println("Database already seeded")
		return nil
	}

	// Begin a transaction
	tx := db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	// Seed 50 users
	fkUUID := faker.UUID{}
	users := make([]models.User, 50)
	for i := range users {
		user := models.User{}
		uuidStr, err := fkUUID.Hyphenated(reflect.Value{})
		if err != nil {
			return fmt.Errorf("failed to generate UUID for user: %w", err)
		}
		user.ID = uuidStr.(string)
		user.Name = fmt.Sprintf("%s %s", faker.FirstName(), faker.LastName())
		user.Username = faker.Username()
		user.Phone = faker.Phonenumber()
		user.Email = faker.Email()
		users[i] = user
	}

	// Seed addresses for each user
	addresses := make([]models.Address, 50)
	for i, user := range users {
		realAddress := faker.GetRealAddress()
		uuidStr, err := fkUUID.Hyphenated(reflect.Value{})
		if err != nil {
			return fmt.Errorf("failed to generate UUID for address: %w", err)
		}

		address := models.Address{
			ID:      uuidStr.(string),
			UserID:  user.ID,
			Street:  realAddress.Address,
			City:    realAddress.City,
			State:   realAddress.State,
			ZipCode: realAddress.PostalCode,
		}
		addresses[i] = address
	}

	// Seed 4 posts per user
	posts := make([]models.Post, 0)
	fkLorem := faker.Lorem{}
	for _, user := range users {
		for i := 0; i < 4; i++ {
			uuidStr, err := fkUUID.Hyphenated(reflect.Value{})
			if err != nil {
				return fmt.Errorf("failed to generate UUID for post: %w", err)
			}
			lorem, err := fkLorem.Paragraph(reflect.Value{})
			if err != nil {
				return fmt.Errorf("failed to generate lorem paragraph for post: %w", err)
			}
			post := models.Post{
				ID:        uuidStr.(string),
				UserID:    user.ID,
				Title:     faker.Sentence(),
				Body:      lorem.(string),
				CreatedAt: time.Now(),
			}
			posts = append(posts, post)
		}
	}

	// Perform bulk inserts within the transaction
	if err := tx.Create(&users).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert users: %w", err)
	}
	if err := tx.Create(&addresses).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert addresses: %w", err)
	}
	if err := tx.Create(&posts).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert posts: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Database seeded successfully!")
	return nil
}
