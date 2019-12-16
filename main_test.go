package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test first input an arbitary slice of numbers can be flatten",
			args: args{
				input: "[[1,2,3],4]",
			},
			want: []string{"1", "2", "3", "4"},
		},
		{
			name: "Test second input an arbitary slice of numbers can be flatten",
			args: args{
				input: "[[1,2,[3]],4]",
			},
			want: []string{"1", "2", "3", "4"},
		},
		{
			name: "Test third input an arbitary slice of numbers can be flatten",
			args: args{
				input: "[[1,[[2]],[3]],4]",
			},
			want: []string{"1", "2", "3", "4"},
		},
		{
			name: "Test input an arbitary slice of strings can be flatten",
			args: args{
				input: "[one,two,three,[two]]",
			},
			want: []string{"one", "two", "three", "two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatten(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_requestAPI(t *testing.T) {
	url := "http://localhost:8080/flat?key=[1,2,3,[2]]"

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
