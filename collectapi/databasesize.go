/*
 Copyright 2016 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package collectapi

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetDatabaseSizeMetrics(logger *log.Logger, dbs []string, HOSTNAME string, dbConn *sql.DB) []Metric {
	logger.Println("get database size metrics")

	var metrics = make([]Metric, 0)

	for i := 0; i < len(dbs); i++ {
		metric := Metric{}

		var dbsize int64
		err := dbConn.QueryRow("select pg_database_size('" + dbs[i] + "') / 1024 / 1024").Scan(&dbsize)
		if err != nil {
			logger.Println("error: " + err.Error())
			return metrics
		}

		metric.Hostname = HOSTNAME
		metric.MetricName = "databasesize"
		metric.Units = "megabytes"
		metric.Value = dbsize
		metric.DatabaseName = dbs[i]
		metrics = append(metrics, metric)
	}

	return metrics

}