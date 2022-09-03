package utils

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

func TestGetAllCustomPaths(t *testing.T) {
	os.Setenv("GHS_CUSTOM_PATH_images/digilex-infordoc-images", "https://storage.gra.cloud.ovh.net/v1/AUTH_ed33ec9e34c64b54aca49d6fcb6dc4c8/infordoc-img")
	os.Setenv("GHS_CUSTOM_REWRITE_images/digilex-infordoc-images", "")
	os.Setenv("GHS_CUSTOM_PATH_images/digilex-infordoc-images-1", "https://storage.gra.cloud.ovh.net/v1/AUTH_ed33ec9e34c64b54aca49d6fcb6dc4c8/img")
	os.Setenv("GHS_CUSTOM_PATH_images/digilex-infordoc-images-2", "https://storage.gra.cloud.ovh.net/v1/AUTH_ed33ec9e34c64b54aca49d6fcb6dc4c8/infordoc")
	result := GetCustomPathsInstance()

	assert.True(t, len(*result) == 3)

	for _, customVar := range *result {
		log.Printf("Path: %s - Target: %s - Rewrite: %t \n", customVar.GetPath(), customVar.GetTarget(), customVar.IsRewrite())
		assert.False(t, strings.HasPrefix(customVar.path, "GHS_CUSTOM_PATH_"))

	}

}
