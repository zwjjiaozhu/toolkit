/*
* @Author: 西园公子
* @File: http_test.go
* @Date: 2024/1/21 14:20
* @IDE:  GoLand
 */

package requests

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url_ := "https://cn.bing.com/"
	resp, err := Get(url_, nil, nil, 0)
	//fmt.Println(resp, err)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp.Body))
}

func TestPost(t *testing.T) {
	url_ := "https://translate.volcengine.com/crx/translate/v1"
	body := map[string]any{
		"source_language": "en",
		"target_language": "zh",
		"text":            "test your name",
	}
	header := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := Post(url_, body, header, nil, 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp.Body))
}

func TestSSE(t *testing.T) {
	url_ := "https://codeapi.fittentech.cn:13443/chat"
	body := map[string]any{
		"inputs":   "<|system|>\nYou are an AI code assistant named Fitten Code which is developed by Fitten Tech.\nThe assistant is happy to help with almost anything, and will do its best to understand exactly what is needed.\nStay focused on current developer request. Consider the possibility that there might not be a solution.\nAsk for clarification if the message does not make sense or more input is needed.\nUse the style of a documentation article.\nOmit any links.\nInclude code snippets (using Markdown) and examples where appropriate.\nPlease refuse to answer sensitive information about politics, pornography, violence, etc.\n在提问是中文时，请尽量使用中文回答。\n<|end|>\n<|user|>\n你的名字\n<|end|>\n<|assistant|>\n",
		"ft_token": "FT_7jRN3XYDJL50dDLYkEbErJpM4bhSkgqkElxYuOkxyvfQU6dI5O",
	}
	header := map[string]string{
		"Content-Type": "application/json",
	}
	callback := func(name, text, mode string) {
		fmt.Println(name, text, mode)
	}
	err := SSE(url_, body, header, nil, 0, callback, "gwe", "append")
	if err != nil {
		t.Error(err)
	}
}

func TestSSEBaichuan(t *testing.T) {
	url_ := "https://www.baichuan-ai.com/api/chat/v1/chat"
	body := map[string]any{
		"type":       "input",
		"request_id": "4b6e838e-fec5-44ff-ad26-2f63145f77f2", //
		"stream":     true,
		"prompt": map[string]any{
			"id":          "",
			"messageId":   "Ua7df011VARxwUd",
			"session_id":  "w9abb011VNQ",
			"data":        "翻译成中文：your name\n",
			"from":        0,
			"parent_id":   0,
			"created_at":  1706668079677,
			"attachments": []string{},
		},
		"app_info": map[string]any{
			"id":   10002,
			"name": "baichuan_wap",
		},
		"user_info": map[string]any{
			"id":     104316,
			"status": 1,
		},
		"session_info": map[string]any{
			"id":         "",
			"session_id": "w9abb011VNQ",
			"name":       "翻译文本\n",
			"created_at": 1706668079676,
		},
		"assistant_info": map[string]any{},
		"parameters": map[string]any{
			"repetition_penalty": -1,
			"temperature":        -1,
			"top_k":              -1,
			"top_p":              -1,
			"max_new_tokens":     -1,
			"do_sample":          -1,
			"regenerate":         0,
			"wse":                true,
		},
		"history":   []string{},
		"assistant": map[string]any{},
		"retry":     3,
	}
	header := map[string]string{
		"Content-Type": "application/json",
		"Cookie":       "next-auth.session-token=eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIn0..3FgtB3R3yIn_twmv.dUIk3vZL2uYo5puTikovXcrJ5bKlj91N22evZY3hODCelwaiE7A-NOa7dR0CMOx5xzN4j7SY2-12UNrjDAfBAB4nOnI2G_1IuvX2BuZZrM01D45RV9BUYGldJ57ZYlacSxFc5h-4ZrnWde5niVlebULhynjAZDyfLCwZu4Hpv138Ci39u5uPuq_03ibxNUywaX0-HxVJK-0hN2zhvmAbg2vy5V0v_zEcGjSaCEDUhLFANfQJtZ_lOe0SPg.TgO8ZouiFv5tibt-T1mpLA",
	}
	callback := func(name, text, mode string) {
		fmt.Println(name, text, mode)
	}
	err := SSE(url_, body, header, nil, 0, callback, "gwe", "append")
	if err != nil {
		t.Error(err)
	}
}

func TestSSETongyi(t *testing.T) {
	url_ := "https://qianwen.biz.aliyun.com/dialog/conversation"
	body := map[string]any{
		"requestId": "3a4ca2f77463401896b7449d4e97c0df",
		"sessionId": "",
		"contents": []any{
			map[string]string{
				"role":        "user",
				"contentType": "text",
				"content":     "帮我写一首诗歌",
			},
		},
		"action":      "next",
		"userAction":  "chat",
		"sessionType": "text_chat",
		"mode":        "chat",
		"parentMsgId": "0",
	}
	header := map[string]string{
		"Content-Type": "application/json",
		"Cookie":       "aliyun_choice=intl; l=fBOLVoP4PFw7bWSdBOfwPurza77OSIRAguPzaNbMi9fPOYCB5-3GW1BKAvT6C3GVF6lXR3yI_B6WBeYBcQd-nxvtGwBLE8DmndLHR35..; _samesite_flag_=true; cookie2=130228fade4e169112d0036e1e606fee; munb=2206480643244; csg=7a7670cf; t=30639aea8d91f5dff692b77826ed2872; _tb_token_=3e5535de3eae3; login_tongyi_ticket=WtFkQQTaqolC3*ZWgm_g5yzTRztfgB8uskxQwwPaHrrXUaIrWMUmL5WA5QsOXO7b0; cna=dPfqHYN17z4CAXTtxzzWxuXa; cnaui=1850209891717305; atpsida=d475dc7d3a3694d88658414e_1706684343_2; login_aliyunid_csrf=_csrf_tk_1202806667120835; aui=1850209891717305; sca=29dd768b; isg=BFRUP-kPzb6bnVki-ktsv4czJZTGrXiXliGs2O4z6V9D2fAjLr94JWwf2dHBFbDv; tfstk=ccTPBpawEwvjq8Ku7ZQF0DQKv_8caS_hWr5NqYW6gvTUYJbhQsb0psjCKs5UgJjl.",
		"Referer":      "https://tongyi.aliyun.com/",
		"X-Platform":   "pc_tongyi",
		"X-Xsrf-Token": "707fcde1-8ae9-4d75-ad2e-07d2231892fc",
	}
	callback := func(name, text, mode string) {
		fmt.Println(name, text, mode)
	}
	err := SSE(url_, body, header, nil, 0, callback, "gwe", "append")
	if err != nil {
		t.Error(err)
	}
}
