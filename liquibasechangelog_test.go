package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

//var testData = []byte(`
//databaseChangeLog:
//  - changeSet:
//      author: "cool"
//      changes:
//        sqlFile:
//          dbms: postgresql
//          encoding: utf-8
//          path: postgres/v0-artikeltable_2020-06-02.sql
//          relativeToChangelogFile: true
//      id: test1
//  - changeSet:
//      author: Johanna Nolte
//      id: v1-bestellungtable_2020-06-17
//      changes:
//        sqlFile:
//          dbms: "foo"
//          encoding: utf-8
//          path: "cool"
//          relativeToChangelogFile: true
//`)

func TestValidate(t *testing.T) {
	testcases := []struct {
		testName    string
		testYaml    []byte
		diagnostics Diagnostics
	}{
		{
			testName: "invalid yaml",
			testYaml: []byte(`not a valid yaml`),
			diagnostics: Diagnostics{
				{
					Summary: "invalid yaml",
					Detail:  "failed to parse yaml",
					Line:    0,
				},
			},
		},
		{
			testName: "invalid databaseChangeLog",
			testYaml: []byte(`databaseChangeLog: wrong`),
			diagnostics: Diagnostics{
				{
					Summary: "invalid databaseChangeLog",
					Detail:  "failed to parse databaseChangeLog",
					Line:    1,
				},
			},
		},
		{
			testName: "empty author",
			testYaml: []byte(`
databaseChangeLog:
  - changeSet:
      id: test
      changes:
        sqlFile:
          dbms: postgresql
          encoding: utf-8
          path: postgres/sample.sql
          relativeToChangelogFile: true
`),
			diagnostics: Diagnostics{
				{
					Summary: "empty author",
					Detail:  "author cannot be empty",
					Line:    4,
				},
			},
		},
		{
			testName: "empty id",
			testYaml: []byte(`
databaseChangeLog:
  - changeSet:
      author: test author
      changes:
        sqlFile:
          dbms: postgresql
          encoding: utf-8
          path: postgres/sample.sql
          relativeToChangelogFile: true
`),
			diagnostics: Diagnostics{
				{
					Summary: "empty id",
					Detail:  "id cannot be empty",
					Line:    4,
				},
			},
		},
		{
			testName: "missing sqlFile dbms",
			testYaml: []byte(`
databaseChangeLog:
  - changeSet:
      id: test
      author: test author
      changes:
        sqlFile:
          encoding: utf-8
          path: postgres/sample.sql
          relativeToChangelogFile: true
`),
			diagnostics: Diagnostics{
				{
					Summary: "empty dbms",
					Detail:  "sqlfile dbms path cannot be empty",
					Line:    8,
				},
			},
		},
		{
			testName: "missing sqlFile encoding",
			testYaml: []byte(`
databaseChangeLog:
  - changeSet:
      id: test
      author: test author
      changes:
        sqlFile:
          dbms: postgresql
          path: postgres/sample.sql
          relativeToChangelogFile: true
`),
			diagnostics: Diagnostics{
				{
					Summary: "empty encoding",
					Detail:  "sqlfile encoding path cannot be empty",
					Line:    8,
				},
			},
		},
		{
			testName: "missing sqlFile path",
			testYaml: []byte(`
databaseChangeLog:
  - changeSet:
      id: test
      author: test author
      changes:
        sqlFile:
          dbms: postgresql
          encoding: utf-8
          relativeToChangelogFile: true
`),
			diagnostics: Diagnostics{
				{
					Summary: "empty path",
					Detail:  "sqlfile path cannot be empty",
					Line:    8,
				},
			},
		},
		{
			testName: "missing sqlFile relativeToChangelogFile",
			testYaml: []byte(`
databaseChangeLog:
  - changeSet:
      id: test
      author: test author
      changes:
        sqlFile:
          dbms: postgresql
          encoding: utf-8
          path: postgres/sample.sql
`),
			diagnostics: Diagnostics{
				{
					Summary: "empty relativeToChangelogFile",
					Detail:  "sqlfile relativeToChangelogFile cannot be empty",
					Line:    8,
				},
			},
		},
		{
			testName: "valid",
			testYaml: []byte(`
databaseChangeLog:
  - changeSet:
      id: test
      author: test author
      changes:
        sqlFile:
          dbms: postgresql
          encoding: utf-8
          path: postgres/sample.sql
          relativeToChangelogFile: true
`),
			diagnostics: Diagnostics{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			diags := Validate(tc.testYaml)
			assert.Equal(t, tc.diagnostics, diags)
		})
	}
}
