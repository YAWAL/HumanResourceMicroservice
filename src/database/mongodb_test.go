package database

import "testing"

func Test_mongoURI(t *testing.T) {

	tests := []struct {
		name string
		host string
		port string
		want string
	}{
		{
			name: "Mongo URI with default port",
			host: "localhost",
			port: "27017",
			want: "mongodb://localhost:27017",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mongoURI(tt.host, tt.port); got != tt.want {
				t.Errorf("mongoURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
