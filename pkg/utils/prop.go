package utils

import (
	"github.com/magiconair/properties"
)

func LoadProperty(path ...string) (*properties.Properties, error) {
	return properties.LoadAll(path, properties.UTF8, true)
}
