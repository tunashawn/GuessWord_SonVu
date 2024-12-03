package api

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGuessWord(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []VoteeResult
		wantErr bool
	}{
		{
			name: "Guess Word - Test connection",
			args: args{"https://wordle.votee.dev:8000/word/ant?guess=ant"},
			want: []VoteeResult{
				{
					Slot:   0,
					Guess:  "a",
					Result: "correct",
				},
				{
					Slot:   1,
					Guess:  "n",
					Result: "correct",
				},
				{
					Slot:   2,
					Guess:  "t",
					Result: "correct",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := guessWord(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("guessWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("guessWord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetGuessType(t *testing.T) {
	word := "abcd"
	fmt.Println(word[:3] + "x" + word[3+1:])
}
