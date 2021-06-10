package alibabaopen

import (
	"testing"
)

func TestClient_handleURI(t *testing.T) {
	type fields struct {
		AppKey      string
		AppSecret   string
		AccessToken string
	}
	type args struct {
		uri string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{name: "test1", fields: fields{
			AppKey:      "1994",
			AppSecret:   "alibaba",
			AccessToken: "5427dd95-18dd-4474-a2b5-dbbf60af5a31",
		}, args: args{uri: "com.alibaba.p4p:alibaba.cps.op.searchCybOffers-1"}, want: "param2/1/com.alibaba.p4p/alibaba.cps.op.searchCybOffers/1994"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				AppKey:      tt.fields.AppKey,
				AppSecret:   tt.fields.AppSecret,
				AccessToken: tt.fields.AccessToken,
			}
			if got := c.handleURI(tt.args.uri); got != tt.want {
				t.Errorf("handleURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSHA1(t *testing.T) {
	type args struct {
		key  string
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{
			key:  "alibaba",
			data: "alibaba",
		}, want: "9a3291da5b3f46c4c853ed511a6072b8e91c85e8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacSHA1(tt.args.key, tt.args.data); got != tt.want {
				t.Errorf("HmacSHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}
