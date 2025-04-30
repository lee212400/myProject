package saas

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	hm "github.com/lee212400/myProject/mock/httpmock"
	"github.com/stretchr/testify/require"
)

func TestGetUser(t *testing.T) {
	httpmock.Activate()                 // httpmock活性化
	defer httpmock.DeactivateAndReset() // テスト後リセット
	client = http.DefaultClient
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   int
		process func()
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			args{"dummy_id"},
			"testName",
			30,
			func() {
				hm.GetUser("dummy_id", "testName", 30, hm.Response200)
			},
			false,
		},
		{
			"fail",
			args{"dummy_id"},
			"testName",
			30,
			func() {
				hm.GetUser("", "", 30, hm.Response400)
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer httpmock.Reset()
			tt.process()

			got, got1, err := GetUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				require.Equal(t, tt.want, got)
				require.Equal(t, tt.want1, got1)
			}
		})
	}
}
