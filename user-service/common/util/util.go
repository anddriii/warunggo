package util

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func BindFromJson(dest any, filename, path string) error {
	v := viper.New()

	v.SetConfigType("json")

	configName := strings.TrimSuffix(filename, filepath.Ext(filename))
	v.SetConfigName(configName)

	// cek apakah path kosong
	if path != "" {
		// path ada
		v.AddConfigPath(path)
	} else {
		// pakai path saat ini jika path kosong
		v.AddConfigPath(".")
	}

	//debug
	fmt.Println("")

	if err := v.ReadInConfig(); err != nil {
		logrus.Errorf("Failed to read config file %v", err)
		return err
	}

	return nil
}

// SetEnvFromConsulKV mengatur environment variables berdasarkan konfigurasi dari Consul KV.
func SetEnvFromConsulKV(v *viper.Viper) error {
	env := make(map[string]any)

	err := v.Unmarshal(&env)
	if err != nil {
		logrus.Errorf("Failed to unmarshal: %v", err)
		return err
	}

	for k, v := range env {
		var (
			valeOf = reflect.ValueOf(v)
			val    string
		)

		switch valeOf.Kind() {
		case reflect.String:
			val = valeOf.String()
		case reflect.Int:
			val = strconv.Itoa(int(valeOf.Int()))
		case reflect.Uint:
			val = strconv.Itoa(int(valeOf.Uint()))
		case reflect.Float32:
			val = strconv.Itoa(int(valeOf.Float()))
		case reflect.Float64:
			val = strconv.Itoa(int(valeOf.Float()))
		case reflect.Bool:
			val = strconv.FormatBool(valeOf.Bool())
		}

		err = os.Setenv(k, val)
		if err != nil {
			logrus.Errorf("Failed to set env: %v", err)
			return err
		}
	}
	return nil
}

func BindFromConsul(dest any, endpoint, path string) error {
	v := viper.New()

	v.SetConfigType("json")
	err := v.AddRemoteProvider("consul", endpoint, path)
	if err != nil {
		logrus.Errorf("failed to add remote provider: %v", err)
		return err
	}

	err = v.ReadRemoteConfig()
	if err != nil {
		logrus.Errorf("Failed to read config: %v", err)
		return err
	}

	err = v.Unmarshal(&dest)
	if err != nil {
		logrus.Errorf("Failed to unmarshal: %v", err)
		return err
	}

	err = SetEnvFromConsulKV(v)
	if err != nil {
		logrus.Errorf("Failed to unmarshal: %v", err)
		return err
	}

	return nil
}
