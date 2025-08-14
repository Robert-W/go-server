package sample

import (
	"context"
	"testing"
)

// These tests are placeholders as the service is currently hardcoded to return
// a list of samples, once we use a real database, this will need to change
func TestSampleService(t *testing.T) {
	ctx := context.Background()
	service := sampleService{}

	t.Run("should return a list of samples", func(t *testing.T) {
		samples, _ := service.listAllSamples(ctx)

		if len(*samples) != 3 {
			t.Errorf("Expected three samples, got %d", len(*samples))
		}
	})

	t.Run("should create a set of samples", func(t *testing.T) {
		samples, _ := service.createSamples(ctx)

		if len(*samples) != 1 {
			t.Errorf("Expected one sample, got %d", len(*samples))
		}
	})

	t.Run("should read a single sample", func(t *testing.T) {
		sample, _ := service.getSampleById(ctx)

		if sample.Id != "123" {
			t.Errorf("Expected sample with ID 123, got %s", sample.Id)
		}
	})

	t.Run("should update a single sample", func(t *testing.T) {
		sample, _ := service.updateSampleById(ctx)

		if sample.Id != "321" {
			t.Errorf("Expected sample with ID 321, got %s", sample.Id)
		}
	})

	t.Run("should delete a single sample", func(t *testing.T) {
		sample, _ := service.deleteSampleById(ctx)

		if sample.Id != "321" {
			t.Errorf("Expected sample with ID 321, got %s", sample.Id)
		}
	})
}
