package idgen

import (
	"git.querycap.com/practice/srv-demo-suns/pkg/constants/types"
	"time"

	"github.com/go-courier/snowflakeid"
)

const (
	// 时间可以占用的bit数为42位
	workerBit   = 12
	sequenceBit = 10
)

var startTime, _ = time.Parse(time.RFC3339, "2010-01-01T00:00:00Z")
var sff = snowflakeid.NewSnowflakeFactory(workerBit, sequenceBit, 1, startTime)

func MustNewIDGen() IDGen {
	idgen, err := newIDGen()
	if err != nil {
		panic(err)
	}
	return idgen
}

func newIDGen() (IDGen, error) {
	sf, err := sff.NewSnowflake(WorkerIDFromIP(ResolveLocalIP(), workerBit))
	if err != nil {
		return nil, err
	}
	return &IDGenImpl{
		sf: sf,
	}, nil
}

type IDGenImpl struct {
	sf *snowflakeid.Snowflake
}

func (idGen *IDGenImpl) GenerateID() (types.SFID, error) {
	id, err := idGen.sf.ID()
	return types.SFID(id), err
}
func (idGen *IDGenImpl) GenerateIDs(n int) (types.SFIDs, error) {
	var ids types.SFIDs
	for i := 0; i < n; i++ {
		id, err := idGen.sf.ID()
		if err != nil {
			return nil, err
		}
		ids = append(ids, types.SFID(id))
	}
	return ids, nil
}

type IDGen interface {
	GenerateID() (types.SFID, error)
	GenerateIDs(n int) (types.SFIDs, error)
}
