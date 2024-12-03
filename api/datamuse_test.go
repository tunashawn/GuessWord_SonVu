package api

import (
	"testing"
)

func Test_fetchWordsFromDatamuse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Test 1",
			args:    args{"h?llo"},
			want:    "hello",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchWordsFromDatamuse(tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchWordsFromDatamuse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 || got[0].Word != tt.want {
				t.Errorf("fetchWordsFromDatamuse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
