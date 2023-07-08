package reposirories

import (
	"github.com/go-redis/redismock/v9"
	"github.com/pkg/errors"
	"testing"
	"time"
)

func TestRedisRepository_normalizeKey(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "positive normal en",
			args: struct{ query string }{query: "Limfield"},
			want: "limfield",
		},
		{
			name: "positive normal utf",
			args: struct{ query string }{query: "Пано Полемидия Πάνω Πολεμίδια"},
			want: "пано полемидия πάνω πολεμίδια",
		},
		{
			name: "positive strip",
			args: struct{ query string }{query: " Арбат "},
			want: "арбат",
		},
		{
			name: "positive strip",
			args: struct{ query string }{query: " Арбат "},
			want: "арбат",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _ := redismock.NewClientMock()
			r := &RedisRepository{
				db: db,
			}
			if got := r.normalizeKey(tt.args.query); got != tt.want {
				t.Errorf("normalizeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisRepository_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "positive",
			args: args{
				key: "арбат",
			},
			want: "OK",
		},
		{
			name: "negative",
			args: args{
				key: "",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := redismock.NewClientMock()
			r := &RedisRepository{
				db: db,
			}

			if tt.wantErr {
				mock.ExpectGet(tt.args.key).SetErr(errors.New("expected error"))
			} else {
				mock.ExpectGet(tt.args.key).SetVal(tt.want)
			}

			got, err := r.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisRepository_Set(t *testing.T) {
	type args struct {
		key   string
		value string
		ttl   time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "positive",
			args: args{
				key:   "арбат",
				value: "{\"street\": \"арбат\"}",
				ttl:   time.Hour * 30,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := redismock.NewClientMock()
			r := &RedisRepository{
				db: db,
			}

			mock.ExpectSet(tt.args.key, tt.args.value, tt.args.ttl).SetVal("OK")

			if err := r.Set(tt.args.key, tt.args.value, tt.args.ttl); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisRepository_Delete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "positive",
			args: args{
				key: "арбат",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := redismock.NewClientMock()
			r := &RedisRepository{
				db: db,
			}

			mock.ExpectDel(tt.args.key).SetVal(0)

			if err := r.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
