package main

type Interceptor interface {

	/**
	 * Called just before a request. This can be used to view or modify the request.
	 * this could be used for things like
	 *      - Logging the request
	 *      - Injecting tracer headers
	 *      - Record metrics
	 * The default implementation does nothing at all
	 * @param request the request
	 * @param config the current configuration
	 */
	OnRequest(request HttpRequest, config Config)

	/**
	 * Called just after the request. This can be used to view the response,
	 * Perhaps for logging purposes or just because you're curious.
	 *  @param response the response
	 *  @param request a summary of the request
	 *  @param config the current configuration
	 */
	OnResponse(response HttpResponse, request HttpRequestSummary, config Config)

	/**
	 * Called in the case of a total failure.
	 * This would be where Unirest was completely unable to make a request at all for reasons like:
	 *      - DNS errors
	 *      - Connection failure
	 *      - Connection or Socket timeout
	 *      - SSL/TLS errors
	 *
	 * The default implimentation simply wraps the exception in a UnirestException and throws it.
	 * It is possible to return a different response object from the original if you really
	 * didn't want to every throw exceptions. Keep in mind that this is a lie
	 *
	 * Nevertheless, you could return something like a kong.unirest.FailedResponse
	 *
	 * @param e the exception
	 * @param request the original request
	 * @param config the current config
	 * @return a alternative response.
	 */
	OnFail(e error, request HttpRequestSummary, config Config) (HttpResponse, error)
}
