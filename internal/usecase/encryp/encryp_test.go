package encryp

import (
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
	"github.com/SETTER2000/shorturl/config"
	"github.com/stretchr/testify/require"
	"reflect"
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
		wantErr interface{}
	}{
		{
			name: "positive test #1 should return string",
			args: args{
				secretKey: secretSecret,
			},
			wantErr: nil,
		}, {
			name: "negative test #2 should return not a string",
			args: args{
				secretKey: "",
			},
			wantErr: ErrEncryptToken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encrypt{}
			got, err := e.EncryptToken(tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				//t.Errorf("EncryptToken() error = %v, wantErr %v", err, tt.wantErr)
				require.EqualValues(t, err, tt.wantErr, "Error type must match.")
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

func TestEncrypt_cipher(t *testing.T) {
	type fields struct {
		cfg *config.Cookie
	}
	type args struct {
		key [32]byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    cipher.Block
		wantErr interface{}
	}{
		{
			name:    "positive test #1",
			args:    args{key: sha256.Sum256([]byte(secretSecret))},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encrypt{
				cfg: tt.fields.cfg,
			}
			got, err := e.cipher(tt.args.key)
			fmt.Println(err)
			fmt.Println(tt.wantErr)

			if (err != nil) != tt.wantErr {
				require.IsType(t, err, tt.wantErr, "These two arguments must be of the same type.")
				//t.Errorf("cipher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cipher() got = %v, want %v", got, tt.want)
			}
		})
	}
}
