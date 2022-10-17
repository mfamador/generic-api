// Package cassandra defines the cassandra data source
package cassandra //nolint

import (
	"fmt"
	"genericsapi/internal/cryptography"
	"genericsapi/internal/genericsapiv1"
	"genericsapi/internal/model"
	"genericsapi/internal/repository"
	"time"

	"github.com/gocql/gocql"
	"github.com/rs/zerolog/log"
)

type foo struct {
	session      *gocql.Session
	cryptography cryptography.Cryptography
}

// NewFoo handles the instantiation
func NewFoo(session *gocql.Session) repository.ListRepository[model.Foo] {
	return &foo{
		session,
		GetCryptography(),
	}
}

// List retrieves a list
func (c *foo) List(limit uint, cursor *string, filter []*genericsapiv1.Filter) ([]*model.Foo, *string, error) {
	f := buildFilter(filter)
	queryStr := fmt.Sprintf(
		`SELECT id, name, specific_foo, value, timestamp
				FROM data.Foos
				WHERE %s ALLOW FILTERING`, f)
	log.Debug().Msgf("query: %s", queryStr)
	res := make([]*model.Foo, 0)
	var (
		page  []byte
		iter  *gocql.Iter
		query *gocql.Query
	)
	query = c.session.Query(queryStr)
	if limit == 0 {
		// no pagination
		iter = query.Iter()
	} else {
		if cursor != nil {
			page, _ = c.cryptography.DecryptString(*cursor, nil)
			cursor = nil
		}
		iter = query.PageSize(int(limit)).PageState(page).Iter()
		page = iter.PageState()
		if len(page) > 0 {
			cstr, _ := c.cryptography.EncryptAsString(page, nil)
			cursor = &cstr
		}
	}
	scanner := iter.Scanner()
	for scanner.Next() {
		var (
			timestamp         time.Time
			ID                int
			value             float64
			name, specificFoo string
		)
		err := scanner.Scan(&ID, &name, &specificFoo, &value, &timestamp)
		if err != nil {
			log.Debug().Msgf("error: %v", err)
			return nil, nil, err
		}
		res = append(res, &model.Foo{
			ID:          int64(ID),
			Name:        name,
			Value:       value,
			SpecificFoo: specificFoo,
			Timestamp:   timestamp,
		})
	}
	if err := iter.Close(); err != nil {
		log.Error().Msgf("query failed : %v", err)
		return nil, nil, err
	}
	return res, cursor, nil
}
