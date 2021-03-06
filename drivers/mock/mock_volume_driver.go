package mock

import (
	"github.com/emccode/rexray/core"
	"github.com/emccode/rexray/core/errors"
)

const mockVolDriverName = "mockVolumeDriver"

type mockVolDriver struct {
	name string
}

type badMockVolDriver struct {
	mockVolDriver
}

func newVolDriver() core.Driver {
	var d core.VolumeDriver = &mockVolDriver{mockVolDriverName}
	return d
}

func newBadVolDriver() core.Driver {
	var d core.VolumeDriver = &badMockVolDriver{
		mockVolDriver{BadMockVolDriverName}}
	return d
}

func (m *mockVolDriver) Init(r *core.RexRay) error {
	return nil
}

func (m *badMockVolDriver) Init(r *core.RexRay) error {
	return errors.New("init error")
}

func (m *mockVolDriver) Name() string {
	return m.name
}

func (m *mockVolDriver) Mount(
	volumeName, volumeID string,
	overwriteFs bool, newFsType string) (string, error) {
	return "", nil
}

func (m *mockVolDriver) Unmount(volumeName, volumeID string) error {
	return nil
}

func (m *mockVolDriver) Path(volumeName, volumeID string) (string, error) {
	return "", nil
}

func (m *mockVolDriver) Create(volumeName string, opts core.VolumeOpts) error {
	return nil
}

func (m *mockVolDriver) Remove(volumeName string) error {
	return nil
}

func (m *mockVolDriver) Attach(volumeName, instanceID string) (string, error) {
	return "", nil
}

func (m *mockVolDriver) Detach(volumeName, instanceID string) error {
	return nil
}

func (m *mockVolDriver) NetworkName(
	volumeName, instanceID string) (string, error) {
	return "", nil
}
