package plantower

import (
	"bytes"
	"encoding/hex"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	t.Run("takes a reading", func(t *testing.T) {
		s := FileStream(t, "testdata/hex.txt")
		r, err := ReadNext(s)
		require.NoError(t, err)
		require.Equal(t, &Reading{
			pm1_lab:   4,
			pm2_5_lab: 12,
			pm10_lab:  14,
			pm1_atm:   4,
			pm2_5_atm: 12,
			pm10_atm:  14,
			pc_0_3:    498,
			pc_0_5:    448,
			pc_1:      70,
			pc_2_5:    0,
			pc_5:      0,
			pc_10:     0,
		}, r)
	})

	t.Run("takes a reading when it finds the start byte", func(t *testing.T) {
		s := FileStream(t, "testdata/hex_invalid_start.txt")
		r, err := ReadNext(s)
		require.NoError(t, err)
		require.Equal(t, &Reading{
			pm1_lab:   4,
			pm2_5_lab: 12,
			pm10_lab:  14,
			pm1_atm:   4,
			pm2_5_atm: 12,
			pm10_atm:  14,
			pc_0_3:    496,
			pc_0_5:    446,
			pc_1:      70,
			pc_2_5:    0,
			pc_5:      0,
			pc_10:     0,
		}, r)
	})
}

func FileStream(t *testing.T, f string) io.Reader {
	t.Helper()

	raw, err := os.ReadFile(f)
	require.NoError(t, err)

	b := bytes.Replace(raw, []byte(" "), []byte(""), -1)
	b = bytes.TrimSuffix(b, []byte("\n"))

	stream, err := hex.DecodeString(string(b))
	require.NoError(t, err)

	buf := bytes.NewBuffer(stream)

	return buf
}
