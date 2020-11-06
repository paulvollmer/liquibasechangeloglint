package main

import (
  "gopkg.in/yaml.v3"
)

func Validate(data []byte) Diagnostics {
	diags := make(Diagnostics, 0)

	var liquibaseChangelog LiquibaseChangelog

	err := yaml.Unmarshal(data, &liquibaseChangelog)
	if err != nil {
		diags = diags.Append(&Diagnostic{"invalid yaml", "failed to parse yaml", 0})
		return diags
	}

	changelogs := []Changelog{}
	decErr := liquibaseChangelog.DatabaseChangeLog.Decode(&changelogs)
	if decErr != nil {
		diags = diags.Append(&Diagnostic{"invalid databaseChangeLog", "failed to parse databaseChangeLog", liquibaseChangelog.DatabaseChangeLog.Line})
		return diags
	}

	for _, changelog := range changelogs {
		changeset := ChangeSet{}
		decErr := changelog.ChangeSet.Decode(&changeset)
		if decErr != nil {
			diags = diags.Append(&Diagnostic{"Failed to parse changeset", decErr.Error(), changelog.ChangeSet.Line})
			return diags
		}

		if changeset.Author.Value == "" {
			diags = diags.Append(&Diagnostic{"empty author", "author cannot be empty", changelog.ChangeSet.Line})
		}

		if changeset.ID.Value == "" {
			diags = diags.Append(&Diagnostic{"empty id", "id cannot be empty", changelog.ChangeSet.Line})
		}

		if len(changeset.Changes.Content) == 0 {
			diags = diags.Append(&Diagnostic{"missing changes block", "changes block must be defined", changelog.ChangeSet.Line})
			return diags
		}

		changes := Changes{}
		decErr = changeset.Changes.Decode(&changes)
		if decErr != nil {
			diags = diags.Append(&Diagnostic{"Failed to parse changes", err.Error(), changeset.Changes.Line})
			return diags
		}

		if len(changes.SQLFile.Content) == 0 {
			diags = diags.Append(&Diagnostic{"missing sqlfile block", "sqlfile block must be defined", changeset.Changes.Line})
			return diags
		}

		sqlfile := SQLFile{}
		decErr = changes.SQLFile.Decode(&sqlfile)
		if decErr != nil {
			diags = diags.Append(&Diagnostic{"Failed to parse sqlfile", err.Error(), changes.SQLFile.Line})
			return diags
		}

		if sqlfile.Dbms.Value == "" {
			diags = diags.Append(&Diagnostic{"empty dbms", "sqlfile dbms path cannot be empty", changes.SQLFile.Line})
		}

		if sqlfile.Encoding.Value == "" {
			diags = diags.Append(&Diagnostic{"empty encoding", "sqlfile encoding path cannot be empty", changes.SQLFile.Line})
		}

		if sqlfile.Path.Value == "" {
			diags = diags.Append(&Diagnostic{"empty path", "sqlfile path cannot be empty", changes.SQLFile.Line})
		}

		if sqlfile.RelativeToChangelogFile.Value == "" {
			diags = diags.Append(&Diagnostic{"empty relativeToChangelogFile", "sqlfile relativeToChangelogFile cannot be empty", changes.SQLFile.Line})
		}
	}

	return diags
}

type LiquibaseChangelog struct {
	DatabaseChangeLog yaml.Node `yaml:"databaseChangeLog"`
}

type Changelog struct {
	ChangeSet yaml.Node `yaml:"changeSet"`
}

type ChangeSet struct {
	ID      yaml.Node `yaml:"id"`
	Author  yaml.Node `yaml:"author"`
	Changes yaml.Node `yaml:"changes"`
}

type Changes struct {
	SQLFile yaml.Node `yaml:"sqlFile"`
}

type SQLFile struct {
	Dbms                    yaml.Node `yaml:"dbms"`
	Encoding                yaml.Node `yaml:"encoding"`
	Path                    yaml.Node `yaml:"path"`
	RelativeToChangelogFile yaml.Node `yaml:"relativeToChangelogFile"`
}
