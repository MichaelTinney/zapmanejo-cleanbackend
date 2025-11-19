package database

import "github.com/MichaelTinney/zapmanejo-cleanbackend/internal/models"

func SeedLifetimeSlots() {
	var count int64
	DB.Model(&models.LifetimeSlot{}).Count(&count)
	if count > 0 {
		return // already seeded
	}

	for i := 1; i <= 200; i++ {
		slot := models.LifetimeSlot{
			Slot:     i,
			Occupied: false,
		}
		DB.Create(&slot)
	}
}
