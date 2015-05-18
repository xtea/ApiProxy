package app

type ApiInfo struct {
	Method string
	Url    string
}

var apiMapping = map[string]ApiInfo{
	"/tianqi":    ApiInfo{"GET", "http://apistore.baidu.com/microservice/weather"}, // citypinyin=shanghai
	"/fenci":     ApiInfo{"GET", "http://api.pullword.com/get.php"},                // source=清华大学是好学校&param1=0&param2=1
	"/fencipost": ApiInfo{"POST", "http://api.pullword.com/post.php"},              // source=清华大学是好学校&param1=0&param2=1
}

func GetApiInfoById(k string) (ApiInfo, bool) {
	ret, ok := apiMapping[k]
	return ret, ok
}
