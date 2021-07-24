package main

type Headers struct {
	Headers []Header
}

type Entry struct {
	Name  string
	Value string
}



func NewHeaders() *Headers {
	var headers = new(Headers)
	headers.Headers = make([]Header, 0)
	return headers
}

func NewEntry(name string, value string){
	var entry Entry
	entry.Name = name
	entry.Value = value
}


    /**
     * Add a header element
     * @param name the name of the header
     * @param value the value for the header
     */
	 public (h *Headers) void add(String name, String value) {
        add(name, () -> value);
    }

    /**
     * Add a header element with a supplier which will be evaluated on request
     * @param name the name of the header
     * @param value the value for the header
     */
    public void add(String name, Supplier<String> value) {
        if (Objects.nonNull(name)) {
            headers.add(new Entry(name, value));
        }
    }

    /**
     * Replace a header value. If there are multiple instances it will overwrite all of them
     * @param name the name of the header
     * @param value the value for the header
     */
    public void replace(String name, String value) {
        remove(name);
        add(name, value);
    }

    private void remove(String name) {
        headers.removeIf(h -> isName(h, name));
    }

    /**
     * Get the number of header keys.
     * @return the size of the header keys
     */
    public int size() {
        return headers.stream().map(Header::getName).collect(toSet()).size();
    }

    /**
     * Get all the values for a header name
     * @param name name of the header element
     * @return a list of values
     */
    public List<String> get(String name) {
        return headers.stream()
                .filter(h -> isName(h, name))
                .map(Header::getValue)
                .collect(toList());
    }

    /**
     * Add a bunch of headers at once
     * @param header a header
     */
    public void putAll(Headers header) {
        this.headers.addAll(header.headers);
    }

    /**
     * Check if a header is present
     * @param name a header
     * @return if the headers contain this name.
     */
    public boolean containsKey(String name) {
        return this.headers.stream().anyMatch(h -> isName(h, name));
    }

    /**
     * Clear the headers!
     */
    public void clear() {
        this.headers.clear();
    }

    /**
     * Get the first header value for a name
     * @param key the name of the header
     * @return the first value
     */
    public String getFirst(String key) {
        return headers
                .stream()
                .filter(h -> isName(h, key))
                .findFirst()
                .map(Header::getValue)
                .orElse("");
    }

    /**
     * Get all of the headers
     * @return all the headers, in order
     */
    public List<Header> all() {
        return new ArrayList<>(this.headers);
    }

    private boolean isName(Header h, String name) {
        return Util.nullToEmpty(name).equalsIgnoreCase(h.getName());
    }

    void remove(String key, String value) {
        List<Header> header = headers.stream().
                filter(h -> key.equalsIgnoreCase(h.getName()) && value.equalsIgnoreCase(h.getValue()))
                .collect(toList());
        headers.removeAll(header);
    }

    /**
     * @return list all headers like this: <pre>Content-Length: 42
     * Cache-Control: no-cache
     * ...</pre>
     */
    @Override
    public String toString() {
       final StringJoiner sb = new StringJoiner(System.lineSeparator());
        headers.forEach(header -> sb.add(header.toString()));
        return sb.toString();
    }

    public void cookie(Cookie cookie) {
        headers.add(new Entry("cookie", cookie.toString()));
    }

    public void cookie(Collection<Cookie> cookies) {
        cookies.forEach(this::cookie);
    }
