package interfaces

import "../../models"
type SedesDao interface {
	FindBySedes (idSedes int) (models.Sedes,error)
}
