package class

import (
    "errors"
)

var (
    ErrNoSpellSlots = errors.New("No spell slots of that level")
    ErrMaxSpellSlots = errors.New("Max spell slots of that level")
)
