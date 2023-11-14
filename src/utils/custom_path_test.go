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
	os.Setenv("GHS_CUSTOM_PATH_images.digilex-infordoc-images-1", "https://storage.gra.cloud.ovh.net/v1/AUTH_ed33ec9e34c64b54aca49d6fcb6dc4c8/img")
	os.Setenv("GHS_CUSTOM_PATH_/images/digilex-infordoc-images-2", "https://storage.gra.cloud.ovh.net/v1/AUTH_ed33ec9e34c64b54aca49d6fcb6dc4c8/infordoc")
	result := GetCustomPathsInstance()

	assert.True(t, len(*result) == 3)

	for _, customVar := range *result {
		log.Printf("Path: %s - Target: %s - Rewrite: %t \n", customVar.GetPath(), customVar.GetTarget(), customVar.IsRewrite())
		assert.False(t, strings.HasPrefix(customVar.path, "GHS_CUSTOM_PATH_"))
		assert.False(t, strings.HasPrefix(customVar.path, "/"))
		assert.True(t, strings.Contains(customVar.path, "/"))
	}

}
func TestFindPath(t *testing.T) {
	os.Setenv("GHS_CUSTOM_PATH_images/digilex-infordoc-images", "https://storage.gra.cloud.ovh.net/v1/AUTH_ed33ec9e34c64b54aca49d6fcb6dc4c8/infordoc-img")
	os.Setenv("GHS_CUSTOM_REWRITE_images/digilex-infordoc-images", "")
	os.Setenv("GHS_CUSTOM_PATH_images.digilex-infordoc/1", "https://storage.gra.cloud.ovh.net/v1/AUTH_ed33ec9e34c64b54aca49d6fcb6dc4c8/img")
	os.Setenv("GHS_CUSTOM_PATH_/images/digilex-infordoc-img", "https://storage.gra.cloud.ovh.net/v1/AUTH_ed33ec9e34c64b54aca49d6fcb6dc4c8/infordoc")

	result, path := FindPath("/images/digilex-infordoc-images")

	assert.True(t, result)
	assert.True(t, path.GetTarget() == "https://storage.gra.cloud.ovh.net/v1/AUTH_ed33ec9e34c64b54aca49d6fcb6dc4c8/infordoc-img")
	assert.True(t, path.IsRewrite())

	result1, path1 := FindPath("/images/digilex-infordoc/1")

	assert.True(t, result1)
	assert.True(t, path1.GetTarget() == "https://storage.gra.cloud.ovh.net/v1/AUTH_ed33ec9e34c64b54aca49d6fcb6dc4c8/img")
	assert.True(t, path1.IsRewrite() == false)

	result2, path2 := FindPath("/images/digilex-infordoc-img")

	assert.True(t, result2)
	assert.True(t, path2.GetTarget() == "https://storage.gra.cloud.ovh.net/v1/AUTH_ed33ec9e34c64b54aca49d6fcb6dc4c8/infordoc")
	assert.True(t, path2.IsRewrite() == false)

}
