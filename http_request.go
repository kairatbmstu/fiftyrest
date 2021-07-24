package main

import "time"

type HttpRequest interface {

	/**
	 * add a route param that replaces the matching {name}
	 * For example routeParam("name", "fred") will replace {name} in
	 * https://localhost/users/{user}
	 * to
	 * https://localhost/users/fred
	 *
	 * @param name the name of the param (do not include curly braces {}
	 * @param value the value to replace the placeholder with
	 * @return this request builder
	 */
	RouteParam(name string, value string) HttpRequest

	/**
	 * add a route param map that replaces the matching {name}
	 * For example routeParam(Map.of("name", "fred")) will replace {name} in
	 * https://localhost/users/{user}
	 * to
	 * https://localhost/users/fred
	 *
	 * @param params a map of path params
	 * @return this request builder
	 */
	RouteParamWithParameters(params map[string]interface{}) HttpRequest

	/**
	 * Basic auth credentials
	 * @param username the username
	 * @param password the password
	 * @return this request builder
	 */
	BasicAuth(username string, password string) HttpRequest

	/**
	 * The Accept header to send (e.g. application/json
	 * @param value a valid mime type for the Accept header
	 * @return this request builder
	 */
	Accept(value string) HttpRequest

	/**
	 * The encoding to expect the response to be for cases where the server fails to respond with the proper encoding
	 * @param encoding a valid mime type for the Accept header
	 * @return this request builder
	 */
	ResponseEncoding(encoding string) HttpRequest

	/**
	 * Add a http header, HTTP supports multiple of the same header. This will continue to append new values
	 * @param name name of the header
	 * @param value value for the header
	 * @return this request builder
	 */
	Header(name string, value string) HttpRequest

	/**
	 * Replace a header value or add it if it doesn't exist
	 * @param name name of the header
	 * @param value value for the header
	 * @return this request builder
	 */
	HeaderReplace(name string, value string) HttpRequest

	/**
	 * Add headers as a map
	 * @param headerMap a map of headers
	 * @return this request builder
	 */
	Headers(headerMap map[string]interface{}) HttpRequest

	/**
	 * Add a simple cookie header
	 * @param name the name of the cookie
	 * @param value the value of the cookie
	 * @return this request builder
	 */
	Cookie(name string, value string) HttpRequest

	/**
	 * Add a simple cookie header
	 * @param cookie a cookie
	 * @return this request builder
	 */
	CookieAsCookie(cookie Cookie) HttpRequest

	/**
	 * Add a collection of cookie headers
	 * @param cookies a cookie
	 * @return this request builder
	 */
	Cookies(cookies []Cookie) HttpRequest

	/**
	 * add a query param to the url. The value will be URL-Encoded
	 * @param name the name of the param
	 * @param value the value of the param
	 * @return this request builder
	 */
	QueryString(name string, value interface{}) HttpRequest

	/**
	 * Add multiple param with the same param name.
	 * queryString("name", Arrays.asList("bob", "linda")) will result in
	 * ?name=bob&amp;name=linda
	 * @param name the name of the param
	 * @param value a collection of values
	 * @return this request builder
	 */
	QueryStringWithValues(name string, values []interface{}) HttpRequest

	/**
	 * Add query params as a map of name value pairs
	 * @param parameters a map of params
	 * @return this request builder
	 */
	QueryStringWithParameters(parameters map[string]interface{}) HttpRequest

	/**
	 * Pass a ObjectMapper for the request. This will override any globally
	 * configured ObjectMapper
	 * @param mapper the ObjectMapper
	 * @return this request builder
	 */
	WithObjectMapper(mapper ObjectMapper) HttpRequest

	/**
	 * Set a socket timeout for this request
	 * @param millies the time in millies
	 * @return this request builder
	 */
	SocketTimeout(millies int) HttpRequest

	/**
	 * Set a connect timeout for this request
	 * @param millies the time in millies
	 * @return this request builder
	 */
	ConnectTimeout(millies int) HttpRequest

	/**
	 * Set a proxy for this request. Only basic proxies are supported.
	 * @param host the host url
	 * @param port the proxy port
	 * @return this request builder
	 */
	Proxy(host string, port int) HttpRequest

	/**
	 * sets a download monitor for monitoring the response. this could be used for drawing a progress bar
	 * @param monitor a ProgressMonitor
	 * @return this request builder
	 */
	DownloadMonitor(monitor ProgressMonitor) HttpRequest

	/**
	 * Executes the request and returns the response with the body mapped into a String
	 * @return response
	 */
	AsString() StringHttpResponse

	/**
	 * Executes the request and returns the response with the body mapped into a byte[]
	 * @return response
	 */
	AsBytes() BytesHttpResponse

	/**
	 * Executes the request and returns the response with the body mapped into a JsonNode
	 * @return response
	 */
	AsJson() JsonHttpResponse

	/**
	 * Executes the request and returns the response with the body mapped into T by a configured ObjectMapper
	 * @param responseClass the class to return. This will be passed to the ObjectMapper
	 * @param <T> the return type
	 * @return a response
	 */
	AsObject() HttpResponse

	/**
	 * Execute the request and pass the raw response to a function for mapping.
	 * This raw response contains the original InputStream and is suitable for
	 * reading large responses.
	 * @param function the function to map the response into a object of T
	 * @param <T> The type of the response mapping
	 * @return A HttpResponse containing T as the body
	 */
	//AsObject(Function<RawResponse, T> function) HttpResponse

	/**
	 * Executes the request and writes the contents into a file
	 * @param path The path to the file.
	 * @param copyOptions options specifying how the copy should be done
	 * @return a HttpResponse with the file containing the results
	 */
	AsFile(path string, copyOptions []CopyOption) FileHttpResponse

	/**
	 * Allows for following paging links common in many APIs.
	 * Each request will result in the same request (headers, etc) but will use the "next" link provided by the extract function.
	 *
	 * @param <T> the type of response.
	 * @param mappingFunction a function to return the desired return type leveraging one of the as* methods (asString, asObject, etc).
	 * @param linkExtractor a function to extract a "next" link to follow. Retuning a null or empty string ends the paging
	 * @return a PagedList of your type
	 */
	//  <T> PagedList<T> asPaged(Function<HttpRequest, HttpResponse> mappingFunction,
	//   Function<HttpResponse<T>, String> linkExtractor);

	/**
	 * Executes the request and returns the response without parsing the body
	 * @return the basic HttpResponse
	 */
	AsEmpty() HttpResponse

	/**
	 * Execute the request and pass the raw response to a consumer.
	 * This raw response contains the original InputStream and is suitable for
	 * reading large responses
	 * @param consumer a consumer function
	 */
	//void thenConsume(Consumer<RawResponse> consumer);

	/**
	 * Execute the request asynchronously and pass the raw response to a consumer.
	 * This raw response contains the original InputStream and is suitable for
	 * reading large responses
	 * @param consumer a consumer function
	 */
	//void thenConsumeAsync(Consumer<RawResponse> consumer);

	/**
	 * @return The HTTP method of the request
	 */
	getHttpMethod() HttpMethod

	/**
	 * @return The current URL string for the request
	 */
	GetUrl() string

	/**
	 * @return the current headers for the request
	 */
	GetHeaders() Headers

	/**
	 * @return if the request has a body it will be here.
	 */
	getBody() Body

	/**
	 * @return socket timeout for this request
	 */
	GetSocketTimeout() int

	/**
	 * @return the connect timeout for this request
	 */
	GetConnectTimeout() int

	/**
	 * @return the proxy for this request
	 */
	GetProxy() Proxy

	/**
	 * @return a summary for the response, used in metrics
	 */
	ToSummary() HttpRequestSummary

	/**
	 * @return the instant the request object was created in UTC (not when it was sent).
	 */
	GetCreationTime() time.Time
}
