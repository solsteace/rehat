package services

import "github.com/solsteace/rest/repositories"

type Reservation struct {
	repositories.Room
	repositories.Class
}
