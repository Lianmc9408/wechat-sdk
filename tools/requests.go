package tools

import (
	"bytes"
	"crypto/tls"
	"encoding/pem"
	"errors"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"net/http"
)

const contentType = "application/xml; charset=utf-8"

func Post(url string, data []byte) ([]byte, error) {
	response, err := http.DefaultClient.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if body != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}

func PostWithCert(url, mchID string, data []byte, certPath string) ([]byte, error) {
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, errors.New("读取证书失败")
	}
	// 将pkcs12证书转成pem
	cert, err := pkcs12ToPem(certData, mchID)
	if err != nil {
		return nil, errors.New("pkcs12证书转成pem失败" + err.Error())
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{*cert},
	}
	transport := &http.Transport{
		TLSClientConfig:    config,
		DisableCompression: true,
	}
	client := &http.Client{Transport: transport}
	response, err := client.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if body != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}

func pkcs12ToPem(p12 []byte, password string) (*tls.Certificate, error) {
	blocks, err := pkcs12.ToPEM(p12, password)
	if err != nil {
		return nil, err
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

func Get(url string) ([]byte, error) {
	response, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if body != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}