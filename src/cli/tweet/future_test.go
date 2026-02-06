package tweet

import (
	"testing"
)


func TestHandleTweetTime(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantHours   int
		wantMinutes int
		wantErr     bool
	}{
		{
			name:        "hours and minutes",
			input:       "2h3m",
			wantHours:   2,
			wantMinutes: 3,
			wantErr:     false,
		},
		{
			name:        "only hours",
			input:       "5h",
			wantHours:   5,
			wantMinutes: 0,
			wantErr:     false,
		},
		{
			name:        "only minutes",
			input:       "30m",
			wantHours:   0,
			wantMinutes: 30,
			wantErr:     false,
		},
		{
			name:        "empty string",
			input:       "",
			wantHours:   0,
			wantMinutes: 0,
			wantErr:     true,
		},
		{
			name:        "invalid format garbage",
			input:       "abc",
			wantHours:   0,
			wantMinutes: 0,
			wantErr:     true,
		},
		{
			name:        "missing suffix",
			input:       "2h3",
			wantHours:   0,
			wantMinutes: 0,
			wantErr:     true,
		},
		{
			name:        "invalid number",
			input:       "xh",
			wantHours:   0,
			wantMinutes: 0,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHours, gotMinutes, err := handleTweetTime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleTweetTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHours != tt.wantHours {
				t.Errorf("handleTweetTime() gotHours = %v, want %v", gotHours, tt.wantHours)
			}
			if gotMinutes != tt.wantMinutes {
				t.Errorf("handleTweetTime() gotMinutes = %v, want %v", gotMinutes, tt.wantMinutes)
			}
		})
	}
}

func TestHandleFutureTweetArgs(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantTweet string
		wantTime  string
		wantErr   bool
	}{
		{
			name:      "valid args",
			args:      []string{"x", "f", "hello world", "2h"},
			wantTweet: "hello world",
			wantTime:  "2h",
			wantErr:   false,
		},
		{
			name:      "not enough args",
			args:      []string{"x", "f"},
			wantTweet: "",
			wantTime:  "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTweet, gotTime, err := handleFutureTweetArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleFutureTweetArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTweet != tt.wantTweet {
				t.Errorf("handleFutureTweetArgs() gotTweet = %v, want %v", gotTweet, tt.wantTweet)
			}
			if gotTime != tt.wantTime {
				t.Errorf("handleFutureTweetArgs() gotTime = %v, want %v", gotTime, tt.wantTime)
			}
		})
	}
}
