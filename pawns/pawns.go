package pawns

import (
	"errors"
	"strconv"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Pawner interface {
	Create(*gorm.DB) error
	Read(*gorm.DB) error
	Update(*gorm.DB) error
	Delete(*gorm.DB) error
	HardDelete(*gorm.DB) error
	List(*gorm.DB) ([]Pawn, error)
	AttackTo(*gorm.DB, uint) error
	MoveTo(*gorm.DB, uint16, uint16) error
	IsRouteValid(*gorm.DB, uint16, uint16) bool
	InitiatePawn(*gorm.DB, uint16, uint16) error
}

type Pawn struct {
	gorm.Model
	UserID    uint
	GameID    uint
	LocationX uint16
	LocationY uint16
	Health    int16
	Defense   int16
	Attack    int16
	Speed     int16
	Affect    int16
	Range     int8
	Type      string
}

func (p *Pawn) InitiatePawn() error {
	switch p.Type {
	case "cavalry":
		p.Health = 100
		p.Defense = 50
		p.Attack = 60
		p.Speed = 6
		p.Affect = 1
		p.Range = 1
		return nil
	default:
		return errors.New("unknown pawn type")
	}
}
func (p *Pawn) AttackTo(db *gorm.DB, pawnID uint) error {
	log.Info("Pawn " + strconv.Itoa(int(p.ID)) + " attacking to Pawn " + strconv.Itoa(int(pawnID)))
	defenderPawn := Pawn{Model: gorm.Model{ID: pawnID}}
	err := defenderPawn.Read(db)
	if err != nil {
		return err
	}
	defenderPawn.Health = defenderPawn.Health - (p.Attack - defenderPawn.Defense)
	log.Info("Pawn " + strconv.Itoa(int(defenderPawn.ID)) + " took " + strconv.Itoa(int(p.Attack-defenderPawn.Defense)) + " damage.Now its health point is: " + strconv.Itoa(int(defenderPawn.Health)))
	p.Health = p.Health - (defenderPawn.Attack - p.Defense)
	log.Info("Pawn " + strconv.Itoa(int(p.ID)) + " took " + strconv.Itoa(int(defenderPawn.Attack-p.Defense)) + " damage.Now its health point is: " + strconv.Itoa(int(p.Health)))
	err = p.Update(db)
	if err != nil {
		return err
	}
	err = defenderPawn.Update(db)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pawn) MoveTo(db *gorm.DB, X uint16, Y uint16) error {
	err := p.Read(db)
	if err != nil {
		return err
	}
	if p.IsRouteValid(db, X, Y) {
		p.LocationX = X
		p.LocationY = Y
		err := p.Update(db)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("this move is not permitted")
	}

}

func (p *Pawn) IsRouteValid(db *gorm.DB, X uint16, Y uint16) bool {

	return true
}
