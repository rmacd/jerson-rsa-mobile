package rsa

import (
	"fmt"
	"testing"
)

func TestFastRSA_DecryptOAEP(t *testing.T) {

	instance := NewFastRSA()
	output, err := instance.DecryptOAEP(cipherTextOAEP, "", "sha256", privateKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("output:", output)
}

var testCipher_gh56 = `QK5pP83ZaaHF+HOr8XAZ6kXyFvSCQWEKlGgEdEQbCvpEa4rq57poO38+gBYXOtrsvTtxOKAbW96FgybrVRdrW95IkyqUN121FuiHBCkhSUNav188KYKqmZEj4Kf+V7h9pljtQrcxaVNp9nOSfGyO70v18ATOuSlUEP0QiJVsynnuf7t+SQTNuSHqSjAnoQcu9/KsZFakfRSETg6Z1FuNeex3hkOhRl2K3fVhhZsQsEAW1KKrP2ONHEtoSlV09CWpaTlx9nEnmNovh/1FAuHZ7+Nu2UqQKKT+rpa8kaij2Ks7Fb87uhZRnA8gU+APcMnHBfca+GHMvMX7jJQuOujSiA==`
var testKey_gh56 = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAqINTkDRjJ3dtq28jbWJaUhIWN4OKaLj50D6V0bDkspXDgviLf2ZvEfmg6DlrplVl4156lDojylFvpuNDEC9YItcuVD43095mvmho/LOypLcwx1RKN6etY5RYX1YZC0si9qsDMgElZUrTEwXeEYaYLS5VG+uOixxa94f1seEy9usLd8ISMzcb0sEvurZTx6hBwE61R9slJTlL7paWvJ6vbnaUEDyo7uLNqsyl0158wfHRFJJ9itA689Yk3cEZL50IBMpShmu7flqxXbKQGSu8Vcw1zLIrSruvUuUD+w2SZVY3Kj3W9bQm+OW6DY2llB7/mTApo2yNoS3mS8hOeZhVIQIDAQABAoIBAAdqii+vQore3G5+TGiYusFgZuZoDVwZAfOKgZTyLEp4wVs+/YKsmyhHHHAD8Osn2H1Wb94cpe0WQuvtFgs9DiuwXStbdqEGKW7pUO+IKA1MO6MWIFMNN9oJUg9WDnGbC62f7prRhpgL4Di22iRa0FEyuA+rwQD5zT4WF3fN4ayQrrmQaiLAYw0QvYZQDl7JK4ocQbVjgB7cHmPjB8G/zFKZhHwgLrYwDEnL7mz5f2t2aMLrURovc1Opnwgk14SqdUGpJ1XMPE6TIzxaTPdqncGBpwTTspr5iJqvriYW2GjAT72HIXy3t3BpBOJt4P5A7D0p9jTHJQRY+1QFgcr9Ec0CgYEA5J2HsYV1HPKoolvYhvFYXkTLTqDn9rKnHJP4MxTTQcSstjCgzAuDN7TlrXo5csNC6t7AElO22wPWwiEL9dM/O5VEQ0V/3+veABp6dIRh+TTKR6miNQ9UNrUjCdhrOW40FX23oEZIFMpstId8/yA81oXVenda//uxhHDe1nJac/UCgYEAvLLDRNxPLr4ICQcfPiw/kplgUDSn0z0vIZLl9bQ3Q0pX9eoF5gqxORbViBmZzzfsz2oo9LN3aaV9bUaarINmWTH0slm+vLk2b2eyCGPlulS9RhNgmrrCazdqssWu1yBMIxBdvRmxEP6GZdpROn014cuZDFi/zaQ+TXwiFzdLTP0CgYAEVukoeXwLnJ+O1Wd6yEIBKBUj3PIKQMHjTPu9HHwWF4Gfw7SJqv5GpGxnqpZEk2hFxQyoTGaAKcZ90NrFQ8lDfEXbcQpIWdXQ8q+4XnrtnA7q5VFq6GuUzkNoAG+om2rprYU6yZq8qkr98kRxI0+EUu4GcRWNHl30QaA5Odp1sQKBgQCK+e4vbUM0Xel1HLW6CMTZp/TznZRtVAa+Z37Os+hvuvWFvNKTVxSnw1WJY7GQmNPk/38imns8aBI0xWdt32kmEFD0enysaozZCDprS4gK8BZm4iaoTxyZ8rq26DmZX8Qznv4rJBzxM0SxB1YECewBXP5fxY2eW3U3hFFnX+Yp0QKBgQDbY41wFkswdz9BnzRF25eMlBQnAjbXnO93FWWtexPXzLmqXJOt4lgF2YTv6YlosS5zrAGhaDHCZTiPrajqILeheS4RmYyzm6UvI2w6QJ9NbxR/6plxTdCuhuL8JTBIakjRNAWWQa2l71rDvOeMvkh8ES00We9PRLTyG3iMruasEA==\n-----END RSA PRIVATE KEY-----"

func TestFastRSA_DecryptOAEP_02(t *testing.T) {
	instance := NewFastRSA()
	output, err := instance.DecryptOAEP(testCipher_gh56, "", "sha256", testKey_gh56)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%x\n", output)
}

func TestFastRSA_DecryptOAEPBytesBigMessage(t *testing.T) {

	instance := NewFastRSA()
	inputMessage := getBigInputMessage()

	cipherTextOAEP, err := instance.EncryptOAEPBytes([]byte(inputMessage), "", "sha256", publicKey)
	if err != nil {
		t.Fatal(err)
	}

	output, err := instance.DecryptOAEPBytes(cipherTextOAEP, "", "sha256", privateKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("output:", string(output))
}
