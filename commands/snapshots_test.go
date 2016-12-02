/*
Copyright 2016 The Doctl Authors All rights reserved.
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

package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnapshotCommand(t *testing.T) {
	cmd := Snapshot()
	assert.NotNil(t, cmd)
	assertCommandNames(t, cmd, "delete", "list", "get")
}

func TestSnapshotList(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.snapshots.On("List").Return(testSnapshotList, nil)

		err := RunSnapshotList(config)
		assert.NoError(t, err)
	})
}

func TestSnapshotDelete(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.snapshots.On("Delete", "1").Return(nil)

		config.Args = append(config.Args, "1")
		//config.Doit.Set(config.NS, doctl.ArgDeleteForce, true)

		err := RunSnapshotDelete(config)
		assert.Error(t, err)

	})
}
