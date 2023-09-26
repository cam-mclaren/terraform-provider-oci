// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AwrDatabaseSnapshotCollection The list of AWR snapshots for one database.
type AwrDatabaseSnapshotCollection struct {

	// The name of the query result.
	Name *string `mandatory:"true" json:"name"`

	// The version of the query result.
	Version *string `mandatory:"false" json:"version"`

	// The time taken to query the database tier (in seconds).
	DbQueryTimeInSecs *float64 `mandatory:"false" json:"dbQueryTimeInSecs"`

	// A list of AWR snapshot summary data.
	Items []AwrDatabaseSnapshotSummary `mandatory:"false" json:"items"`
}

// GetName returns Name
func (m AwrDatabaseSnapshotCollection) GetName() *string {
	return m.Name
}

// GetVersion returns Version
func (m AwrDatabaseSnapshotCollection) GetVersion() *string {
	return m.Version
}

// GetDbQueryTimeInSecs returns DbQueryTimeInSecs
func (m AwrDatabaseSnapshotCollection) GetDbQueryTimeInSecs() *float64 {
	return m.DbQueryTimeInSecs
}

func (m AwrDatabaseSnapshotCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDatabaseSnapshotCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AwrDatabaseSnapshotCollection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwrDatabaseSnapshotCollection AwrDatabaseSnapshotCollection
	s := struct {
		DiscriminatorParam string `json:"awrResultType"`
		MarshalTypeAwrDatabaseSnapshotCollection
	}{
		"AWRDB_SNAPSHOT_SET",
		(MarshalTypeAwrDatabaseSnapshotCollection)(m),
	}

	return json.Marshal(&s)
}
