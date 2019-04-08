// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package prometheus

import (
	"context"
	"database/sql"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/utils/logger"
	"github.com/percona/pmm-managed/utils/tests"
)

var configPath = filepath.Join("..", "..", "testdata", "prometheus", "prometheus.yml")

func setup(t *testing.T) (context.Context, *reform.DB, *Service, []byte) {
	t.Helper()

	ctx := logger.Set(context.Background(), t.Name())

	sqlDB := tests.OpenTestDB(t)
	db := reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(t.Logf))

	svc, err := NewService(configPath, "promtool", db, "http://127.0.0.1:9090/prometheus/")
	require.NoError(t, err)
	require.NoError(t, svc.Check(ctx))

	original, err := ioutil.ReadFile(configPath) //nolint:gosec
	require.NoError(t, err)

	return ctx, db, svc, original
}

func teardown(t *testing.T, db *reform.DB, svc *Service, original []byte) { //nolint:golint
	t.Helper()

	assert.NoError(t, ioutil.WriteFile(configPath, original, 0644))
	assert.NoError(t, svc.reload())

	assert.NoError(t, db.DBInterface().(*sql.DB).Close())
}

func TestPrometheus(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		ctx, db, svc, original := setup(t)
		defer teardown(t, db, svc, original)

		assert.NoError(t, svc.UpdateConfiguration(ctx))

		actual, err := ioutil.ReadFile(configPath) //nolint:gosec
		require.NoError(t, err)
		assert.Equal(t, string(original), string(actual))
	})

	t.Run("Normal", func(t *testing.T) {
		ctx, db, svc, original := setup(t)
		defer teardown(t, db, svc, original)

		for _, str := range []reform.Struct{
			&models.Node{
				NodeID:       "/node_id/cc663f36-18ca-40a1-aea9-c6310bb4738d",
				NodeType:     models.GenericNodeType,
				NodeName:     "test-generic-node",
				Address:      pointer.ToString("1.2.3.4"),
				CustomLabels: []byte(`{"_node_label": "foo"}`),
			},

			&models.Service{
				ServiceID:    "/service_id/014647c3-b2f5-44eb-94f4-d943260a968c",
				ServiceType:  models.MySQLServiceType,
				ServiceName:  "test-mysql",
				NodeID:       "/node_id/cc663f36-18ca-40a1-aea9-c6310bb4738d",
				Address:      pointer.ToString("5.6.7.8"),
				CustomLabels: []byte(`{"_service_label": "bar"}`),
			},

			&models.Agent{
				AgentID:      "/agent_id/75bb30d3-ef4a-4147-97a8-621a996611dd",
				AgentType:    models.MySQLdExporterType,
				CustomLabels: []byte(`{"_agent_label": "baz"}`),
				ListenPort:   pointer.ToUint16(12345),
			},
			&models.AgentService{
				AgentID:   "/agent_id/75bb30d3-ef4a-4147-97a8-621a996611dd",
				ServiceID: "/service_id/014647c3-b2f5-44eb-94f4-d943260a968c",
			},

			&models.Service{
				ServiceID:    "/service_id/9cffbdd4-3cd2-47f8-a5f9-a749c3d5fee1",
				ServiceType:  models.PostgreSQLServiceType,
				ServiceName:  "test-postgresql",
				NodeID:       "/node_id/cc663f36-18ca-40a1-aea9-c6310bb4738d",
				Address:      pointer.ToString("5.6.7.8"),
				CustomLabels: []byte(`{"_service_label": "bar"}`),
			},

			&models.Agent{
				AgentID:      "/agent_id/29e14468-d479-4b4d-bfb7-4ac2fb865bac",
				AgentType:    models.PostgresExporterType,
				CustomLabels: []byte(`{"_agent_label": "postgres-baz"}`),
				ListenPort:   pointer.ToUint16(12345),
			},
			&models.AgentService{
				AgentID:   "/agent_id/29e14468-d479-4b4d-bfb7-4ac2fb865bac",
				ServiceID: "/service_id/9cffbdd4-3cd2-47f8-a5f9-a749c3d5fee1",
			},

			// disabled
			&models.Agent{
				AgentID:    "/agent_id/4226ddb5-8197-443c-9891-7772b38324a7",
				AgentType:  models.NodeExporterType,
				Disabled:   true,
				ListenPort: pointer.ToUint16(12345),
			},
			&models.AgentNode{
				AgentID: "/agent_id/4226ddb5-8197-443c-9891-7772b38324a7",
				NodeID:  "/node_id/cc663f36-18ca-40a1-aea9-c6310bb4738d",
			},
		} {
			require.NoError(t, db.Insert(str))
		}

		assert.NoError(t, svc.UpdateConfiguration(ctx))

		expected := `# Managed by pmm-managed. DO NOT EDIT.
---
global:
  scrape_interval: 1m
  scrape_timeout: 10s
  evaluation_interval: 1m
rule_files:
- /etc/prometheus.d/*.rules.yml
scrape_configs:
- job_name: prometheus
  scrape_interval: 1s
  scrape_timeout: 1s
  metrics_path: /prometheus/metrics
  static_configs:
  - targets:
    - 127.0.0.1:9090
    labels:
      instance: pmm-server
- job_name: grafana
  scrape_interval: 5s
  scrape_timeout: 4s
  metrics_path: /metrics
  static_configs:
  - targets:
    - 127.0.0.1:3000
    labels:
      instance: pmm-server
- job_name: pmm-managed
  scrape_interval: 10s
  scrape_timeout: 5s
  metrics_path: /debug/metrics
  static_configs:
  - targets:
    - 127.0.0.1:7773
    labels:
      instance: pmm-server
- job_name: mysqld_exporter_agent_id_75bb30d3-ef4a-4147-97a8-621a996611dd_hr
  scrape_interval: 1s
  scrape_timeout: 1s
  metrics_path: /metrics-hr
  static_configs:
  - targets:
    - 1.2.3.4:12345
    labels:
      _agent_label: baz
      _node_label: foo
      _service_label: bar
      instance: /agent_id/75bb30d3-ef4a-4147-97a8-621a996611dd
      node_id: /node_id/cc663f36-18ca-40a1-aea9-c6310bb4738d
      node_name: test-generic-node
      service_id: /service_id/014647c3-b2f5-44eb-94f4-d943260a968c
      service_name: test-mysql
- job_name: mysqld_exporter_agent_id_75bb30d3-ef4a-4147-97a8-621a996611dd_mr
  scrape_interval: 10s
  scrape_timeout: 5s
  metrics_path: /metrics-mr
  static_configs:
  - targets:
    - 1.2.3.4:12345
    labels:
      _agent_label: baz
      _node_label: foo
      _service_label: bar
      instance: /agent_id/75bb30d3-ef4a-4147-97a8-621a996611dd
      node_id: /node_id/cc663f36-18ca-40a1-aea9-c6310bb4738d
      node_name: test-generic-node
      service_id: /service_id/014647c3-b2f5-44eb-94f4-d943260a968c
      service_name: test-mysql
- job_name: mysqld_exporter_agent_id_75bb30d3-ef4a-4147-97a8-621a996611dd_lr
  scrape_interval: 1m
  scrape_timeout: 10s
  metrics_path: /metrics-lr
  static_configs:
  - targets:
    - 1.2.3.4:12345
    labels:
      _agent_label: baz
      _node_label: foo
      _service_label: bar
      instance: /agent_id/75bb30d3-ef4a-4147-97a8-621a996611dd
      node_id: /node_id/cc663f36-18ca-40a1-aea9-c6310bb4738d
      node_name: test-generic-node
      service_id: /service_id/014647c3-b2f5-44eb-94f4-d943260a968c
      service_name: test-mysql
- job_name: postgres_exporter_agent_id_29e14468-d479-4b4d-bfb7-4ac2fb865bac
  scrape_interval: 1s
  scrape_timeout: 1s
  metrics_path: /metrics
  static_configs:
  - targets:
    - 1.2.3.4:12345
    labels:
      _agent_label: postgres-baz
      _node_label: foo
      _service_label: bar
      instance: /agent_id/29e14468-d479-4b4d-bfb7-4ac2fb865bac
      node_id: /node_id/cc663f36-18ca-40a1-aea9-c6310bb4738d
      node_name: test-generic-node
      service_id: /service_id/9cffbdd4-3cd2-47f8-a5f9-a749c3d5fee1
      service_name: test-postgresql
`
		actual, err := ioutil.ReadFile(configPath) //nolint:gosec
		require.NoError(t, err)
		assert.Equal(t, expected, string(actual))
	})
}
