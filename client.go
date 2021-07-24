package main

type RawResponseToHttpResponseTransformer func(raw RawResponse) HttpResponse

type Client interface {
	GetClient() interface{}

	Request(request HttpRequest, httpResponse RawResponseToHttpResponseTransformer)

	//  default <T> HttpResponse<T> request(HttpRequest request, Function<RawResponse, HttpResponse<T>> transformer, Class<?> resultType){
	// 	 return request(request, transformer);
	//  }

	Close() error

	RegisterShutdownHook()
}
