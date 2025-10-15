package models

type Position string

const (
	PositionArquero                Position = "Arquero"
	PositionLateralDerecho         Position = "Lateral Derecho"
	PositionDefensorCentral        Position = "Defensor Central"
	PositionLateralIzquierdo       Position = "Lateral Izquierdo"
	PositionMediocampistaDefensivo Position = "Mediocampista Defensivo"
	PositionMediocampistaCentral   Position = "Mediocampista Central"
	PositionMediocampistaOfensivo  Position = "Mediocampista Ofensivo"
	PositionMediocampistaDerecho   Position = "Mediocampista Por Derecha"
	PositionMediocampistaIzquierdo Position = "Mediocampista Por Izquierda"
	PositionExtremoIzquierdo       Position = "Extremo Izquierdo"
	PositionExtremoDerecho         Position = "Extremo Derecho"
	PositionDelanteroCentro        Position = "Delantero Centro"
)

type Player struct {
	ID          uint `gorm:"primarykey"`
	TeamID      *uint
	Team        Team     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Name        string   `gorm:"size:255;not null"`
	Nationality string   `gorm:"size:100;not null"`
	MarketValue float64  `gorm:"not null"`
	Rating      float64  `gorm:"not null"`
	PhotoUrl    string   `gorm:"size:255"`
	Position    Position `gorm:"type:varchar(50);not null"`
}

func GetAvailablePositions() []Position {
	return []Position{
		PositionArquero,
		PositionLateralDerecho,
		PositionDefensorCentral,
		PositionLateralIzquierdo,
		PositionMediocampistaDefensivo,
		PositionMediocampistaCentral,
		PositionMediocampistaOfensivo,
		PositionMediocampistaDerecho,
		PositionMediocampistaIzquierdo,
		PositionExtremoIzquierdo,
		PositionExtremoDerecho,
		PositionDelanteroCentro,
	}
}
