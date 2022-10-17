// Package cassandra holds the cassandra repository implementation
package cassandra

import (
	"crypto/tls"
	"fmt"
	"genericsapi/internal/cryptography"
	"genericsapi/internal/genericsapiv1"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gocql/gocql"
)

// Cluster defines the connection config params
type Cluster struct {
	ConnTimeout time.Duration `json:"connTimeout" yaml:"connTimeout"`
	Timeout     time.Duration `json:"timeout" yaml:"timeout"`
}

// TLS defines TLS config params
type TLS struct {
	Enabled                  bool `json:"enabled" yaml:"enabled"`
	DisableInitialHostLookup bool `json:"disableInitialHostLookup" yaml:"disableInitialHostLookup"`
	EnableHostVerification   bool `json:"enableHostVerification" yaml:"enableHostVerification"`
	SkipCertVerify           bool `json:"skipCertVerify" yaml:"skipCertVerify"`
}

// Config defines cassandra configuration
type Config struct {
	Address  string  `json:"address" yaml:"address"`
	Port     int     `json:"port" yaml:"port"`
	User     string  `json:"user" yaml:"username"`
	Password string  `json:"password" yaml:"password"`
	Limit    uint    `json:"limit" yaml:"limit"`
	Cluster  Cluster `json:"cluster" yaml:"cluster"`
	TLS      TLS     `json:"tls" yaml:"tls"`
}

// GetCryptography creates a new Cryptography with a specific secret
func GetCryptography() cryptography.Cryptography {
	return cryptography.New("b09e58536e4df2a4fc6dd3c9773e4f3d")
}

// GetSession builds the cassandra connection
func GetSession(keyspace string, conf *Config) (*gocql.Session, error) {
	host := fmt.Sprintf("%s:%d",
		conf.Address,
		conf.Port,
	)
	cluster := gocql.NewCluster(host)
	cluster.ConnectTimeout = conf.Cluster.ConnTimeout * time.Second
	cluster.Timeout = conf.Cluster.Timeout * time.Second
	cluster.Port = conf.Port
	cluster.Consistency = gocql.Quorum
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 3}
	if conf.TLS.Enabled {
		cluster.DisableInitialHostLookup = conf.TLS.DisableInitialHostLookup
		cluster.SslOpts = &gocql.SslOptions{
			EnableHostVerification: conf.TLS.EnableHostVerification,
			Config: &tls.Config{
				InsecureSkipVerify: conf.TLS.SkipCertVerify, //nolint:gosec
				MinVersion:         tls.VersionTLS12,
				Renegotiation:      tls.RenegotiateFreelyAsClient,
			},
		}
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: conf.User,
			Password: conf.Password,
		}
	} else {
		cluster.Keyspace = "system"
		sysSession, _ := cluster.CreateSession()
		if err := sysSession.Query(fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS  %s
			WITH replication = {
				'class' : 'SimpleStrategy',
				'replication_factor' : 1
			}`, keyspace)).Exec(); err != nil {
			log.Error().Msgf("error creating keyspace %v", err)
		}
	}
	cluster.Keyspace = keyspace
	session, err := cluster.CreateSession()
	return session, err
}

func buildFilter(filter []*genericsapiv1.Filter) string {
	const bitSize = 64
	filterExp := make([]string, len(filter))
	for i := range filter {
		f := filter[i]
		pattern := "%s"
		if _, err := strconv.ParseFloat(f.StringValue, bitSize); err != nil {
			if _, err := strconv.ParseBool(f.StringValue); err != nil {
				pattern = "'%s'"
			}
		}
		filterExp[i] = fmt.Sprintf("%s %s %s", f.ColumnName, operator(f.Operator), fmt.Sprintf(pattern, f.StringValue))
	}
	return strings.Join(filterExp, " and ")
}

func operator(op genericsapiv1.FilterOperator) string {
	return func() string {
		switch op {
		case genericsapiv1.FilterOperator_EQUAL:
			return "="
		case genericsapiv1.FilterOperator_NOTEQUAL:
			return "!="
		case genericsapiv1.FilterOperator_GREATERTHAN:
			return ">"
		case genericsapiv1.FilterOperator_GREATERTHANOREQUAL:
			return ">="
		case genericsapiv1.FilterOperator_LESSTHAN:
			return "<"
		case genericsapiv1.FilterOperator_LESSTHANOREQUAL:
			return "<="
		}
		return ""
	}()
}
