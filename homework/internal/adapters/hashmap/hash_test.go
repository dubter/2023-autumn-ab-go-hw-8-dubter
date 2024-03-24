package hashmap

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"homework/internal/devices"
)

const (
	testSeqNum1 = "test 1"
	testSeqNum2 = "test 2"
	testSeqNum3 = "test 3"

	testIP1 = "test ip 1"
	testIP2 = "test ip 2"
	testIP3 = "test ip 3"

	testModel1 = "test model 1"
	testModel2 = "test model 2"
	testModel3 = "test model 3"
)

type hashTestSuite struct {
	suite.Suite
	hashRepo *hash
	values   []*devices.Device
}

func TestHashRun(t *testing.T) {
	suite.Run(t, new(hashTestSuite))
}

func (d *hashTestSuite) SetupTest() {
	d.hashRepo = NewHash().(*hash)

	d.hashRepo.hashTable[testSeqNum1] = &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}

	d.values = append(d.values, d.hashRepo.hashTable[testSeqNum1])

	d.hashRepo.hashTable[testSeqNum2] = &devices.Device{
		SerialNum: testSeqNum2,
		IP:        testIP2,
		Model:     testModel2,
	}

	d.values = append(d.values, d.hashRepo.hashTable[testSeqNum2])
}

func (d *hashTestSuite) TestGet() {
	actual, err := d.hashRepo.Get(testSeqNum1)

	require.NoError(d.T(), err)
	require.NotNil(d.T(), actual)
	require.Equal(d.T(), actual, d.values[0])
}

func (d *hashTestSuite) TestGetError() {
	actual, err := d.hashRepo.Get("")

	require.Error(d.T(), err)
	require.Nil(d.T(), actual)
}

func (d *hashTestSuite) TestCreate() {
	device := &devices.Device{
		SerialNum: testSeqNum3,
		IP:        testIP3,
		Model:     testModel3,
	}

	err := d.hashRepo.Create(device)

	require.NoError(d.T(), err)
}

func (d *hashTestSuite) TestCreateError() {
	device := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}

	err := d.hashRepo.Create(device)

	require.Error(d.T(), err)
}

func (d *hashTestSuite) TestDelete() {
	err := d.hashRepo.Delete(testSeqNum1)

	require.NoError(d.T(), err)
}

func (d *hashTestSuite) TestDeleteError() {
	err := d.hashRepo.Delete(testSeqNum3)

	require.Error(d.T(), err)
}

func (d *hashTestSuite) TestUpdate() {
	device := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP3,
		Model:     testModel3,
	}

	err := d.hashRepo.Update(device)

	require.NoError(d.T(), err)

	actual := d.hashRepo.hashTable[testSeqNum1]

	require.Equal(d.T(), actual, device)
}

func (d *hashTestSuite) TestUpdateError() {
	device := &devices.Device{
		SerialNum: testSeqNum3,
		IP:        testIP1,
		Model:     testModel1,
	}

	err := d.hashRepo.Update(device)

	require.Error(d.T(), err)
}

// tests for checking speed processing
func BenchmarkRepoRun(b *testing.B) {
	b.Run("Get device", BenchmarkGet)
	b.Run("Create device", BenchmarkCreate)
	b.Run("Update device", BenchmarkUpdate)
	b.Run("Delete device", BenchmarkDelete)
}

func BenchmarkGet(b *testing.B) {
	hash := NewHash()
	device := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}

	_ = hash.Create(device)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = hash.Get(testSeqNum1)
	}
}

func BenchmarkCreate(b *testing.B) {
	hash := NewHash()
	device := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = hash.Create(device)
	}
}

func BenchmarkUpdate(b *testing.B) {
	hash := NewHash()
	device1 := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}

	device2 := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP2,
		Model:     testModel2,
	}

	_ = hash.Create(device1)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = hash.Update(device2)
	}
}

func BenchmarkDelete(b *testing.B) {
	hash := NewHash()
	device := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}

	_ = hash.Create(device)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = hash.Delete(testSeqNum1)
	}
}

// Fuzz tests
func FuzzGetDevice(f *testing.F) {
	hash := NewHash()

	// Test for correct ID.
	f.Fuzz(func(t *testing.T, seqNum string) {
		expect := &devices.Device{
			SerialNum: testSeqNum1,
			IP:        testIP1,
			Model:     testModel1,
		}

		err := hash.Create(expect)

		require.NoError(t, err)

		actual, err := hash.Get(testSeqNum1)

		require.Equal(t, actual, expect)
		require.NoError(t, err)

		err = hash.Delete(testSeqNum1)

		require.NoError(t, err)

		actual, err = hash.Get(testSeqNum1)

		require.Nil(t, actual)
		require.Error(t, err)
	})
}
