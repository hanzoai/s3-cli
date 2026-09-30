// Copyright (c) 2015-2022 Hanzo AI, Inc.
//
// This file is part of Hanzo S3 stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import "testing"

// Tests valid host URL functionality.
func TestParseEnvURLStr(t *testing.T) {
	testCases := []struct {
		hostURL      string
		accessKey    string
		secretKey    string
		sessionToken string
		hostname     string
		port         string
	}{
		{
			hostURL:   "https://s3user:s3secret1#23@localhost:9000",
			accessKey: "s3user",
			secretKey: "s3secret1#23",
			hostname:  "localhost",
			port:      "9000",
		},
		{
			hostURL:   "https://s3user:s3secret123@@localhost:9000",
			accessKey: "s3user",
			secretKey: "s3secret123@",
			hostname:  "localhost",
			port:      "9000",
		},
		{
			hostURL:   "https://s3user:s3user@123@@localhost:9000",
			accessKey: "s3user",
			secretKey: "s3user@123@",
			hostname:  "localhost",
			port:      "9000",
		},
		{
			hostURL:   "https://localhost:9000",
			accessKey: "",
			secretKey: "",
			hostname:  "localhost",
			port:      "9000",
		},
		{
			hostURL:      "https://s3user:s3secret123:token@localhost:9000",
			accessKey:    "s3user",
			secretKey:    "s3secret123",
			sessionToken: "token",
			hostname:     "localhost",
			port:         "9000",
		},
		{
			hostURL:      "https://s3user:s3user@123:token@@localhost:9000",
			accessKey:    "s3user",
			secretKey:    "s3user@123",
			sessionToken: "token@",
			hostname:     "localhost",
			port:         "9000",
		},
		{
			hostURL:   "https://s3user@localhost:9000",
			accessKey: "s3user",
			hostname:  "localhost",
			port:      "9000",
		},
		{
			hostURL:   "https://s3user:@localhost:9000",
			accessKey: "s3user",
			hostname:  "localhost",
			port:      "9000",
		},
		{
			hostURL:  "https://:@localhost:9000",
			hostname: "localhost",
			port:     "9000",
		},
		{
			hostURL:   "https://:s3secret123@localhost:9000",
			hostname:  "localhost",
			secretKey: "s3secret123",
			port:      "9000",
		},
		{
			hostURL:      "https://:s3secret123:token@localhost:9000",
			hostname:     "localhost",
			secretKey:    "s3secret123",
			sessionToken: "token",
			port:         "9000",
		},
		{
			hostURL:   "https://:s3secret123:@localhost:9000",
			hostname:  "localhost",
			secretKey: "s3secret123",
			port:      "9000",
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			url, ak, sk, token, err := parseEnvURLStr(testCase.hostURL)
			if testCase.accessKey != ak {
				t.Fatalf("Expected %s, got %s", testCase.accessKey, ak)
			}
			if testCase.secretKey != sk {
				t.Fatalf("Expected %s, got %s", testCase.secretKey, sk)
			}
			if testCase.sessionToken != token {
				t.Fatalf("Expected %s, got %s", testCase.sessionToken, token)
			}
			if testCase.hostname != url.Hostname() {
				t.Fatalf("Expected %s, got %s", testCase.hostname, url.Hostname())
			}
			if testCase.port != url.Port() {
				t.Fatalf("Expected %s, got %s", testCase.port, url.Port())
			}
			if err != nil {
				t.Fatalf("Expected test to pass. Failed with err %s", err)
			}
		})
	}
}

func TestParseEnvURLStrInvalid(t *testing.T) {
	_, _, _, _, err := parseEnvURLStr("")
	if err == nil {
		t.Fatalf("Expected failure")
	}
}
