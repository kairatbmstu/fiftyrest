package main

type Client interface {

	 Object getClient()

	 HttpResponse request(HttpRequest request, Function<RawResponse, HttpResponse<T>> transformer);
 
	 default <T> HttpResponse<T> request(HttpRequest request, Function<RawResponse, HttpResponse<T>> transformer, Class<?> resultType){
		 return request(request, transformer);
	 }

	 Stream<Exception> close();
 
	 void registerShutdownHook();

}
