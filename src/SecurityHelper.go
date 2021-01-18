package main

import (
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"strings"
)

func decryptValue(value string, session *session.Session)  string{

	securityService := kms.New(session)

	data, _ := base64.StdEncoding.DecodeString(value)

	inputDecrypt := &kms.DecryptInput{
		CiphertextBlob: data,
	}

	respDecrypt, _ := securityService.Decrypt(inputDecrypt)

	return strings.Trim(string(respDecrypt.Plaintext), " \n")
}