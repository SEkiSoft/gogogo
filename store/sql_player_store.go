// Copyright (c) 2016 David Lu
// See License.txt

package store

import (
	"database/sql"
	"fmt"
	"github.com/davidlu1997/gogogo/model"
	"strconv"
	"strings"
)

type SqlPlayerStore struct {
	*SqlStore
}
