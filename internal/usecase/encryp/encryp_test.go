package encryp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	secretSecret = "RtsynerpoGIYdab_s234r"
)

func TestEncrypt_EncryptToken(t *testing.T) {
	type args struct {
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "positive test #1 should return string",
			args: args{
				secretKey: secretSecret,
			},
			wantErr: false,
		}, {
			name: "negative test #2 should return not a string",
			args: args{
				secretKey: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encrypt{}
			got, err := e.EncryptToken(tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.IsType(t, got, tt.want, "These two arguments must be of type string.")
			//require.EqualValues(t, secretSecret, tt.args.secretKey, "Keys must be the same.")
			//if got != tt.want {
			//	t.Errorf("EncryptToken() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
