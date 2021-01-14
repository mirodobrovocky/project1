package config

type Properties struct {
	Mongo struct {
		Uri			string
		Database 	string
		Collection 	string
	}
	Postgres struct{
		Dns 			string
		SingularTable 	bool	`yaml:"singularTable"`
	}
}
