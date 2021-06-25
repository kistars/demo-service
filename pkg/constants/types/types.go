package types

import (
	toolsdatatype "git.querycap.com/tools/datatypes"
	"github.com/go-courier/sqlx/v2/datatypes"
)

type Bool = datatypes.Bool
type Timestamp = datatypes.Timestamp
type SFID = toolsdatatype.SFID
type SFIDs = toolsdatatype.SFIDs

var (
	BOOL_TRUE  = datatypes.BOOL_TRUE
	BOOL_FALSE = datatypes.BOOL_FALSE
)
