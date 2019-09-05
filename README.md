# go-http-tilezen

Work in progress.

## Example

### TilezenProxyHandler

```
import (
	"net/http"
	tz_http "github.com/sfomuseum/go-http-tilezen/http"
	"github.com/whosonfirst/go-cache-blob"	
)

func main() {

	blob_dsn := "s3://your-bucket?region=us-east-1&prefix=tilezen&credentials=iam:"     
	proxy_timeout := 30
	proxy_url := "/tiles/"
	
	mux := http.NewServeMux()
			
	blob_cache, _ := blob.NewBlobCacheWithDSN(blob_dsn)

	timeout := time.Duration(proxy_timeout) * time.Second
		
	proxy_opts := &tz_http.TilezenProxyHandlerOptions{
		Cache: blob_cache,
		Timeout: timeout,
	}

	proxy_handler, _ := tz_http.TilezenProxyHandler(proxy_opts)

	mux.Handle(proxy_url, proxy_handler)

	http.ListenAndServe(":8080", mux)
```

_Error handling removed for the sake of brevity._

## See also

* https://github.com/sfomuseum/go-tilezen
* https://github.com/whosonfirst/go-cache
* https://github.com/whosonfirst/go-cache-blob