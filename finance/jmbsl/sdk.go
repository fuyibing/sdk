// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package jmbsl

import (
	"github.com/fuyibing/sdk"
)

const (
	Name = "java.mbsl"
)

// Create bill.
func Create(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/insure/policy/paging").SetBody(body).Run(ctx)
}

func buildBody(ctx interface{}, bod interface{}) *sdk.ClientResponse {
	return nil
}

//**
//* 批量发消息
//* @param array $body 入参类型
//* @return ResponseInterface
//*/
//public function batch($body)
//{
//$body = $this->buildBody($body);
//return $this->restful("POST", "/topic/batch", $body);
//}
//
///**
// * 发布单条消息
// * @param array $body 入参类型
// * @return ResponseInterface
// */
//public function publish($body)
//{
//$body = $this->buildBody($body);
//return $this->restful("POST", "/topic/publish", $body);
//}
//
///**
// * @param array|object $body
// * @return array
// */
//private function buildBody($body)
//{
//// 1. 原始入参
//if (!is_array($body)) {
//if (is_object($body)) {
//if (method_exists($body, 'toArray')) {
//$body = $body->toArray();
//} else {
//$body = json_decode(json_encode($body, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE), true);
//}
//} else {
//$body = [];
//}
//}
//// 2. 转成Java入参
//$uuid = md5(json_encode($body, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE));
//$body['tag'] = isset($body['topicTag']) && is_string($body['topicTag']) && $body['topicTag'] !== '' ? $body['topicTag'] : '';
//$body['topic'] = isset($body['topicName']) && is_string($body['topicName']) && $body['topicName'] !== '' ? $body['topicName'] : '';
//$body['reqNo'] = $uuid;
//$body['msgKey'] = isset($body['msgKey']) && is_string($body['msgKey']) && $body['msgKey'] !== '' ? $body['msgKey'] : '';
//return $body;
//}
