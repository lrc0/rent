package types

//Mysql .
type Mysql struct {
	Database Database `yaml:"database"`
}

//Database .
type Database struct {
	Type      string `yaml:"type"`
	URL       string `yaml:"url"`
	MaxIdle   int    `yaml:"maxIdle"`
	MaxActive int    `yaml:"maxActive"`
	ShowSQL   bool   `yaml:"showsql"`
}

//Prices .
type Prices struct {
	Price Price `yaml:"price"`
}

//Price .
type Price struct {
	Water    float32 `yaml:"water"`
	Electric float32 `yaml:"electric"`
	Gas      float32 `yaml:"gas"`
}
