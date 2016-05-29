// Copyright (c) 2016 David Lu
// See License.txt

package store

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"github.com/davidlu1997/gogogo/model"
)

type SqlPlayerStore struct {
	*SqlStore
}