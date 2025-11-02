
# Player Service
provides functionality to handle player data and attachments like transcripts and player sheets.

## APIs

### Health

```curl
% curl --location --request GET 'http://localhost:8086/health' -v
Note: Unnecessary use of -X or --request, GET is already inferred.
* Host localhost:8086 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8086...
* Connected to localhost (::1) port 8086
> GET /health HTTP/1.1
> Host: localhost:8086
> User-Agent: curl/8.7.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Date: Sun, 02 Nov 2025 03:02:39 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
```

### Player

#### READ Player
```curl
% curl --location --request GET 'http://localhost:8086/player' \
--header 'x-pub-email: johndoe@example.com'
{"ID":1,"UUID":"2ba4396d-8b2d-4e18-aeeb-45d29790b6c7","Name":"John Doe","Email":"johndoe@example.com","Age":29,"Team":"Warriors","Score":1500.75,"CreatedAt":"2025-11-02T02:56:19.198435Z","UpdatedAt":"2025-11-02T02:57:59.612846Z","DeletedAt":null}
```

#### UPDATE Player
```
curl --location --request POST 'http://localhost:8086/player' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "age": 29,
  "team": "Warriors",
  "score": 1500.75
}'
```