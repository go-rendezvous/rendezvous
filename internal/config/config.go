package config

type Conf struct {
	DBConf
}

type DBConf struct {
	DiverType   string
	Destination string
}
