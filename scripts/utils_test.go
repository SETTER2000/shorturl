package scripts

import (
	"github.com/SETTER2000/shorturl/config"
	"regexp"
	"testing"
)

func TestCheckEnvironFlag(t *testing.T) {
	type args struct {
		environName string
		flagName    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "positive test #1",
			args: args{
				environName: "environName",
				flagName:    "flagName",
			},
			want: true,
		}, {
			name: "positive test #2 flag variable set",
			args: args{
				environName: "",
				flagName:    "flagName",
			},
			want: true,
		}, {
			name: "positive test #3 same naming of environment variable and flag",
			args: args{
				environName: "environName",
				flagName:    "DATABASE_DSN",
			},
			want: true,
		}, {
			name: "negative test #1",
			args: args{
				environName: "",
				flagName:    "",
			},
			want: false,
		}, {
			name: "negative test #2",
			args: args{
				environName: "environName",
				flagName:    "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckEnvironFlag(tt.args.environName, tt.args.flagName); got != tt.want {
				t.Errorf("CheckEnvironFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHost(t *testing.T) {
	type args struct {
		cfg      config.HTTP
		shorturl string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "positive test #1",
			args: args{
				cfg:      config.HTTP{BaseURL: "http://localhost:8080"},
				shorturl: "wEuWothteri_t23",
			},
			want: "http://localhost:8080/wEuWothteri_t23",
		}, {
			name: "negative test #1 extra closing slash",
			args: args{
				cfg:      config.HTTP{BaseURL: "http://localhost:8080/"},
				shorturl: "wEuWothteri_t23",
			},
			want: "http://localhost:8080//wEuWothteri_t23",
		}, {
			name: "negative test #2 missing protocol",
			args: args{
				cfg:      config.HTTP{BaseURL: "localhost:8080"},
				shorturl: "wEuWothteri_t23",
			},
			want: "localhost:8080/wEuWothteri_t23",
		}, {
			name: "negative test #3 missing protocol and port",
			args: args{
				cfg:      config.HTTP{BaseURL: "localhost"},
				shorturl: "wEuWothteri_t23",
			},
			want: "localhost/wEuWothteri_t23",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHost(tt.args.cfg, tt.args.shorturl); got != tt.want {
				t.Errorf("GetHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateString(t *testing.T) {
	//type args struct {
	//	n int
	//}
	tests := []struct {
		name    string
		pattern string
		n       int
		want    bool
	}{
		{
			name:    "positive test #1",
			n:       10,
			pattern: `[_0-9a-zA-Z]`,
			want:    true,
		}, {
			name:    "positive test #2",
			n:       1,
			pattern: `[_0-9a-zA-Z]{3}`,
			want:    true,
		}, {
			name:    "positive test #3",
			n:       5,
			pattern: `[_0-9a-zA-Z]{5}`,
			want:    true,
		}, {
			name:    "positive test #4",
			n:       0,
			pattern: `[_0-9a-zA-Z]{3}`,
			want:    true,
		}, {
			name:    "negative test #1",
			n:       -1,
			pattern: `[_0-9a-zA-Z]`,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateString(tt.n)
			if match, _ := regexp.MatchString(tt.pattern, got); !match {
				t.Errorf("GenerateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFibo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateString(3)
	}
}
