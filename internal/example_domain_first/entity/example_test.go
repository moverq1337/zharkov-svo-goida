package entity

import (
	"reflect"
	"testing"
)

// more testing tools here
// "github.com/stretchr/testify/assert"

func TestExampleUser_Birthday(t *testing.T) {
	type fields struct {
		Username string
		Age      int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "birthday success",
			fields: fields{
				Username: "Gosha",
				Age:      15,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExampleUser{
				Username: tt.fields.Username,
				Age:      tt.fields.Age,
			}
			if err := e.Birthday(); (err != nil) != tt.wantErr {
				t.Errorf("Birthday() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExampleUser_Rename(t *testing.T) {
	type fields struct {
		Username string
		Age      int
	}
	type args struct {
		newName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "rename success",
			fields: fields{
				Username: "Gosha",
				Age:      15,
			},
			args: args{
				newName: "Gosha",
			},
			wantErr: false,
		},
		{
			name: "rename failed",
			fields: fields{
				Username: "Gosha",
				Age:      15,
			},
			args: args{
				newName: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExampleUser{
				Username: tt.fields.Username,
				Age:      tt.fields.Age,
			}
			if err := e.Rename(tt.args.newName); (err != nil) != tt.wantErr {
				t.Errorf("Rename() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewExampleUser(t *testing.T) {
	type args struct {
		username string
		age      int
	}
	tests := []struct {
		name    string
		args    args
		want    *ExampleUser
		wantErr bool
	}{
		{
			name: "success new user",
			args: args{
				username: "Gosha",
				age:      15,
			},
			want: &ExampleUser{
				Username: "Gosha",
				Age:      15,
			},
			wantErr: false,
		},
		{
			name: "failed new user drops on username",
			args: args{
				username: "",
				age:      15,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed new user drops on zero age",
			args: args{
				username: "Gosha",
				age:      0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed new user drops on negative age",
			args: args{
				username: "Gosha",
				age:      -10,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewExampleUser(tt.args.username, tt.args.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewExampleUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExampleUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
