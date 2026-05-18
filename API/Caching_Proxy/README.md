# _Caching Proxy_
<hr>
Cache Proxy is a simple, bare-bones caching proxy server implemented in Go. It serves as a basic solution for caching HTTP requests and responses, using an on-disk storage to store the cache.

### Features
<hr>

- On-disk Caching:

  - The server uses on-disk to store cached responses. This approach allows data stored on disk persists across application restarts and crashes. This ensures that cached data remains available for subsequent requests, reducing the need to recompute or re-fetch data.

 
- Proxy Functionality:

  - The server forwards client requests to the target server and caches the responses. If the same request is made again, the server returns the cached response, saving the time and resources of making a new request to the target server.


### _Usage_
<hr>
Start the server:

`go run cmd/proxy/main.go --port <PORT> --origin <ORIGIN_URL>`

_Note: the order of port and origin does not matter_

To clear the cache::

`go run cmd/proxy/main.go --clear-cache`

### Limitations
<hr>

#### No Cache Eviction Policy:
- This implementation does not include a cache eviction policy. The cache will grow indefinitely, which could lead to memory exhaustion in long-running scenarios. Adding an eviction policy like LRU (Least Recently Used) would be a recommended enhancement.


### Extras
<hr>

This Repo serves as a solution to Roadmap.sh [Caching Server Problem](https://roadmap.sh/projects/caching-server)