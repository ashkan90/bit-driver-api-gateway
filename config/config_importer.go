package config

type Importer interface {
	ImportInto(fPath string)
}
