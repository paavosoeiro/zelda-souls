package main

import (
	"reflect"
	"testing"
)

func TestVec2_sum(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		vec Vec2
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec2
	}{
		{
			name:   "sum two vec 1,1",
			fields: fields{1, 1},
			args:   args{Vec2{1, 1}},
			want:   Vec2{2, 2},
		},
		{
			name:   "string",
			fields: fields{2, 3},
			args:   args{Vec2{-1, 1}},
			want:   Vec2{1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Sum(tt.args.vec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
