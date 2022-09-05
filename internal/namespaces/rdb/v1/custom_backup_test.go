package rdb

import (
	"os"
	"testing"
	"time"

	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Test_CreateBackup(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			createInstance(engine),
			// We opened an internal issue about the fact that the instance is considered ready even if rdb is not yet available.
			core.BeforeFuncWhenUpdatingCassette(
				func(ctx *core.BeforeFuncCtx) error {
					time.Sleep(1 * time.Minute)
					return nil
				},
			),
		),
		Cmd:           "scw rdb backup create name=foobar expires-at=2999-01-02T15:04:05-07:00 instance-id={{ .Instance.ID }} database-name=rdb --wait",
		Check:         core.TestCheckGolden(),
		AfterFunc:     deleteInstance(),
		DefaultRegion: scw.RegionNlAms,
	}))
}

func Test_RestoreBackup(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			createInstance(engine),
			// We opened an internal issue about the fact that the instance is considered ready even if rdb is not yet available.
			core.BeforeFuncWhenUpdatingCassette(
				func(ctx *core.BeforeFuncCtx) error {
					time.Sleep(1 * time.Minute)
					return nil
				},
			),
			core.ExecStoreBeforeCmd(
				"Backup",
				"scw rdb backup create name=foobar expires-at=2999-01-02T15:04:05-07:00 instance-id={{ .Instance.ID }} database-name=rdb --wait",
			),
		),
		Cmd: "scw rdb backup restore {{ .Backup.ID }} instance-id={{ .Instance.ID }} --wait",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: deleteInstance(),
	}))
}

func Test_ExportBackup(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			createInstance(engine),
			// We opened an internal issue about the fact that the instance is considered ready even if rdb is not yet available.
			core.BeforeFuncWhenUpdatingCassette(
				func(ctx *core.BeforeFuncCtx) error {
					time.Sleep(1 * time.Minute)
					return nil
				},
			),
			core.ExecStoreBeforeCmd(
				"Backup",
				"scw rdb backup create name=foobar expires-at=2999-01-02T15:04:05-07:00 instance-id={{ .Instance.ID }} database-name=rdb --wait",
			),
		),
		Cmd: "scw rdb backup export {{ .Backup.ID }} --wait",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: deleteInstance(),
	}))
}

func Test_DownloadBackup(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			createInstance(engine),
			core.ExecStoreBeforeCmd(
				"Backup",
				"scw rdb backup create name=foobar expires-at=2999-01-02T15:04:05-07:00 instance-id={{ .Instance.ID }} database-name=rdb --wait",
			),
			core.ExecStoreBeforeCmd(
				"BackupExport",
				"scw rdb backup export {{ .Backup.ID }} --wait",
			),
		),
		Cmd: "scw rdb backup download {{ .Backup.ID }} output=dump",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteInstance(),
			func(ctx *core.AfterFuncCtx) error {
				err := os.Remove("dump")
				return err
			},
		),
		DefaultRegion: scw.RegionNlAms,
		TmpHomeDir:    true,
	}))
}

// If ran please update the cassette by changing 'download_url_expires_at'
// when it's not null to a much later date
// E.g. from "2022-09-05T13:14:54.437192Z" to "2999-09-05T13:14:54.437192Z"
func Test_ListBackup(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			createInstance(engine),
			core.ExecStoreBeforeCmd(
				"BackupA",
				"scw rdb backup create name=will_be_exported expires-at=2999-01-02T15:04:05-07:00 instance-id={{ .Instance.ID }} database-name=rdb --wait",
			),
			core.ExecStoreBeforeCmd(
				"BackupB",
				"scw rdb backup create name=will_not_be_exported expires-at=2999-01-02T15:04:05-07:00 instance-id={{ .Instance.ID }} database-name=rdb --wait",
			),
			core.ExecStoreBeforeCmd(
				"BackupExport",
				"scw rdb backup export {{ .BackupA.ID }} --wait",
			),
		),
		Cmd: "scw rdb backup list instance-id={{ .Instance.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: deleteInstance(),
	}))
}
