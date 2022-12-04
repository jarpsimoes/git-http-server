package utils

import (
	"fmt"
	githttpserver1alpha1 "github.com/jarpsimoes/git-http-server-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"log"
	"reflect"
)

var dictionaryPaths = map[string]string{
	"PathClone":   "PATH_CLONE",
	"PathPull":    "PATH_PULL",
	"PathVersion": "PATH_VERSION",
	"PathWebHook": "PATH_WEBHOOK",
	"PathHealth":  "PATH_HEALTH",
	"RepoBranch":  "REPO_BRANCH",
	"RepoTarget":  "REPO_TARGET_FOLDER",
}
var dictionaryDefaults = map[string]string{
	"PathClone":   "_clone",
	"PathPull":    "_pull",
	"PathVersion": "_version",
	"PathWebHook": "_hook",
	"PathHealth":  "_health",
	"RepoBranch":  "main",
	"RepoTarget":  "target-git",
}

// MergeConfigurationWithEnvironmentVariables merge configuration with default values
// If value not found, will be use default value
func MergeConfigurationWithEnvironmentVariables(v githttpserver1alpha1.GitHttpServer) []corev1.EnvVar {
	specs := v.Spec
	vSpecs := reflect.ValueOf(specs)
	tSpecs := vSpecs.Type()

	var envVariables []corev1.EnvVar

	for i := 0; i < vSpecs.NumField(); i++ {
		if _, ok := dictionaryPaths[tSpecs.Field(i).Name]; ok && vSpecs.Field(i).Interface() != "" {
			log.Printf("Field %s defined with value %s\n", dictionaryPaths[tSpecs.Field(i).Name], vSpecs.Field(i).Interface())
			envVariables = append(envVariables, corev1.EnvVar{
				Name:      dictionaryPaths[tSpecs.Field(i).Name],
				Value:     fmt.Sprintf("%s", vSpecs.Field(i).Interface()),
				ValueFrom: nil,
			})
		} else if _, ok := dictionaryPaths[tSpecs.Field(i).Name]; ok && vSpecs.Field(i).Interface() == "" {
			log.Printf("Field %s not defined will be used value %s\n", dictionaryPaths[tSpecs.Field(i).Name], dictionaryDefaults[tSpecs.Field(i).Name])
			envVariables = append(envVariables, corev1.EnvVar{
				Name:      dictionaryPaths[tSpecs.Field(i).Name],
				Value:     dictionaryDefaults[tSpecs.Field(i).Name],
				ValueFrom: nil,
			})
		}
	}
	// Repo URL is mandatory. Not needed check if exists
	envVariables = append(envVariables, corev1.EnvVar{
		Name:      "REPO_URL",
		Value:     specs.RepoURL,
		ValueFrom: nil,
	})

	if specs.RepoUsername != "" && specs.RepoPassword != "" {
		envVariables = append(envVariables, corev1.EnvVar{
			Name:      "REPO_USERNAME",
			Value:     specs.RepoUsername,
			ValueFrom: nil,
		})

		envVariables = append(envVariables, corev1.EnvVar{
			Name:      "REPO_PASSWORD",
			Value:     specs.RepoPassword,
			ValueFrom: nil,
		})
	}

	if specs.HttpPort != 0 {
		envVariables = append(envVariables, corev1.EnvVar{
			Name:      "HTTP_PORT",
			ValueFrom: nil,
			Value:     portString(specs.HttpPort),
		})
	}
	return envVariables
}
func GetProbe(v githttpserver1alpha1.GitHttpServer) corev1.Probe {

	livenessProbe := dictionaryDefaults["PathHealth"]
	port := int32(8081)

	if v.Spec.PathHealth != "" {
		livenessProbe = v.Spec.PathHealth
	}

	if v.Spec.HttpPort != 0 {
		port = v.Spec.HttpPort
	}
	var mProbe = corev1.Probe{
		ProbeHandler: corev1.ProbeHandler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: fmt.Sprintf("/%s", livenessProbe),
				Port: intstr.FromInt(int(port)),
			},
		},
		PeriodSeconds:    10,
		FailureThreshold: 3,
	}

	return mProbe

}
func portString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
