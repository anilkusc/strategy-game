package pawns

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

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
	InitiatePawn(*gorm.DB, uint16, uint16) error
}

type Pawn struct {
	gorm.Model
	UserID    uint
	GameID    uint
	BoardID   uint
	Direction uint8 // 1 right , 2 up , 3 left , 4 down
	X         int16
	Y         int16
	Health    int16
	Defense   int16
	Attack    int16
	Speed     int16
	Affect    int16
	Range     int8 // damageable area
	Agility   int8 // it decides who will attack first
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
		p.Agility = 1
		return nil
	default:
		return errors.New("unknown pawn type")
	}
}

func (p *Pawn) AttackTo(db *gorm.DB, pawnID uint) error {
	defenderPawn := Pawn{Model: gorm.Model{ID: pawnID}}
	err := defenderPawn.Read(db)
	if err != nil {
		return err
	}
	err = p.Read(db)
	if err != nil {
		return err
	}
	if defenderPawn.UserID == p.UserID {
		return nil
	}
	log.Info("Pawn " + strconv.Itoa(int(p.ID)) + " attacking to Pawn " + strconv.Itoa(int(pawnID)))

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

func (p *Pawn) ShufflePawns(db *gorm.DB, pawnlist []uint) ([]Pawn, error) {
	var pawns []Pawn
	var intervalPawns []Pawn
	var shuffledPawns []Pawn
	for _, pawnid := range pawnlist {
		pawn := Pawn{
			Model: gorm.Model{
				ID: pawnid,
			},
		}
		err := pawn.Read(db)
		if err != nil {
			return pawns, err
		}
		pawns = append(pawns, pawn)
		for i := 0; i < int(pawn.Agility); i++ {
			intervalPawns = append(intervalPawns, pawn)
		}
	}
	for len(intervalPawns) > 0 {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := len(intervalPawns) - 1
		poppedPawn := rand.Intn(max-min+1) + min
		shuffledPawns = append(shuffledPawns, intervalPawns[poppedPawn])
		for i := 0; i < len(intervalPawns); i++ {
			if intervalPawns[i].ID == intervalPawns[poppedPawn].ID {
				intervalPawns = append(intervalPawns[:i], intervalPawns[i+1:]...)
			}
		}
	}

	return shuffledPawns, nil

}
