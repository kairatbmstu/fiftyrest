package main

type MapBody func(interface{}) interface{}

type MapHttpResponse func(interface{}) interface{}

type HttpResponseConsumer func(response *HttpResponse)

/**
 * @param <T> a Http Response holding a specific type of body.
 */
type HttpResponse interface {

	/**
	 * @return the HTTP status code.
	 */
	GetStatus() int

	/**
	 * @return status text
	 */
	GetStatusText() string

	/**
	 * @return Response Headers (map) with <b>same case</b> as server response.
	 * For instance use <code>getHeaders().getFirst("Location")</code> and not <code>getHeaders().getFirst("location")</code> to get first header "Location"
	 */
	GetHeaders() Headers

	/**
	 * @return the body
	 */
	GetBody() interface{}

	/**
	 * If the transformation to the body failed by an exception it will be kept here
	 * @return a possible RuntimeException. Checked exceptions are wrapped in a UnirestException
	 */
	GetParsingError() error

	/**
	 * Map the body into another type
	 * @param func a function to transform a body type to something else.
	 * @param <V> The return type of the function
	 * @return the return type
	 */
	MapBody(f MapBody) interface{}

	/**
	 * Map the Response into another response with a different body
	 * @param func a function to transform a body type to something else.
	 * @param <V> The return type of the function
	 * @return the return type
	 */
	Map(f MapHttpResponse) HttpResponse

	/**
	 * If the response was a 200-series response. Invoke this consumer
	 * can be chained with ifFailure
	 * @param consumer a function to consume a HttpResponse
	 * @return the same response
	 */
	IfSuccess(consumer HttpResponseConsumer) HttpResponse

	/**
	 * If the response was NOT a 200-series response or a mapping exception happened. Invoke this consumer
	 * can be chained with ifSuccess
	 * @param consumer a function to consume a HttpResponse
	 * @return the same response
	 */
	IfFailure(consumer HttpResponseConsumer) HttpResponse

	/**
	 * If the response was NOT a 200-series response or a mapping exception happened. map the original body into a error type and invoke this consumer
	 * can be chained with ifSuccess
	 * @param <E> the type of error class to map the body
	 * @param errorClass the class of the error type to map to
	 * @param consumer a function to consume a HttpResponse
	 * @return the same response
	 */
	IfFailureWithError(err error, consumer HttpResponseConsumer) HttpResponse

	/**
	 * @return true if the response was a 200-series response and no mapping exception happened, else false
	 */
	IsSuccess() bool

	/**
	 * Map the body into a error class if the response was NOT a 200-series response or a mapping exception happened.
	 * Uses the system Object Mapper
	 * @param <E> the response type
	 * @param errorClass the class for the error
	 * @return the error object
	 */
	MapError(e error) error

	/**
	 * return a cookie collection parse from the set-cookie header
	 * @return a Cookies collection
	 */
	GetCookies() Cookies
}

type BytesHttpResponse interface {
	HttpResponse
}

type StringHttpResponse interface {
	HttpResponse
}

type ObjectHttpResponse interface {
	HttpResponse
}

type JsonHttpResponse interface {
	HttpResponse
}

type FileHttpResponse interface {
	HttpResponse
}
