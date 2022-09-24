package entity

type Weather struct {
	Wind  int8 `db:"wind"`
	Water int8 `db:"water"`
}

type StatusReport struct {
	Status  string `db:"report"`
	Weather Weather
}
