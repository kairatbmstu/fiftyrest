package main

type Client interface {

	 GetClient() interface{}

	 Request(HttpRequest request, Function<RawResponse, HttpResponse<T>> transformer) HttpResponse 
 
	 default <T> HttpResponse<T> request(HttpRequest request, Function<RawResponse, HttpResponse<T>> transformer, Class<?> resultType){
		 return request(request, transformer);
	 }

	 Stream<Exception> close();
 
	 RegisterShutdownHook()

}
