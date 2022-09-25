package entity

type Weather struct {
	Wind  string `db:"wind"`
	Water string `db:"water"`
}

type StatusReport struct {
	Status  string `db:"report"`
	Weather Weather
}
