package premise

import (
	"github.com/PillLabs/plugman-monorepo/metadata/src/dao"
	"os"
)

var RuntimeVars *Vars

type Vars struct {
	Dao        *dao.PlugmanDao
	PrivateKey string
}

func init() {
	privateKey := os.Getenv("PRIVATE_KEY")
	d := dao.NewPlugmanDao()
	RuntimeVars = &Vars{
		Dao:        d,
		PrivateKey: privateKey,
	}
}
