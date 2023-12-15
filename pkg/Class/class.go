package class

import (
	"crypto/rand"
	"math/big"
)

func (c *Class) GetLevel() uint {
	return c.level
}

func (c *Class) LevelUp() {
	c.level++
}

func (c *Class) GetHitDie() uint {
	return c.hitDie
}

func (c *Class) GetHPRoll() uint {
	roll, err := rand.Int(rand.Reader, big.NewInt(int64(c.hitDie)))
	if err != nil {
		panic(err)
	}
	return uint(roll.Int64() + 1)
}

func (c *Class) GetHPAverage() uint {
	return c.hitDie/2 + 1
}
