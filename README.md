
# Player Service
provides functionality to handle player data and attachments like transcripts and player sheets.

## APIs

### Health

```curl
% curl --location --request GET 'http://localhost:8086/health' -v

< HTTP/1.1 200 OK
```

### Player

#### Read Player
Takes in email and exposes player info
```curl
% curl --location --request GET 'http://localhost:8086/player' \
--header 'x-pub-email: johndoe@example.com'

{
    "ID": 1,
    "UUID": "2ba4396d-8b2d-4e18-aeeb-45d29790b6c7",
    "Name": "John Doe",
    "Email": "johndoe@example.com",
    "Age": 29,
    "Team": "Warriors",
    "Score": 1500.75,
    "CreatedAt": "2025-11-02T02:56:19.198435Z",
    "UpdatedAt": "2025-11-02T03:07:10.688604Z",
    "DeletedAt": null
}
```

#### Update Player
Takes in player information and updates it in the backend
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

< HTTP/1.1 200 OK
```
### Attachment
#### Upload
Takes in email, file-category, file and uploads it to backend.
```curl
curl --location --request POST 'http://localhost:8086/attachment' \
--header 'x-pub-email: johndoe@example.com' \
--header 'X-Pub-File-Category: cv' \
--form 'File=@"/Users/i/Documents/coverLetter.doc"'

{
    "status": "Success",
    "description": "Upload complete for johndoe@example.com"
}
```

#### Download

Takes in email, file-category, filename and downloads the file.

```curl
 % curl --location --request GET 'http://localhost:8086/attachment' \
--header 'x-pub-email: johndoe@example.com' \
--header 'x-pub-file-category: cv' \
--header 'x-pub-file-name: coverLetter.doc' -o coverLetter.doc
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 14848  100 14848    0     0  6087k      0 --:--:-- --:--:-- --:--:-- 7250k
```



## Mocking

Please run mockery to update the mocks for unitesting
```bash
% mockery
2025-11-03T13:12:47.039380000-06:00 INF Starting mockery config-file=/Users/i/Src/pub-a-player/.mockery.yml version=v3.5.5
2025-11-03T13:12:47.039543000-06:00 INF Parsing configured packages... version=v3.5.5
2025-11-03T13:12:47.295027000-06:00 INF Done parsing configured packages. version=v3.5.5
2025-11-03T13:12:47.295229000-06:00 INF adding interface to collection collection=/Users/i/Src/pub-a-player/core/mocks.go interface=Handler package-path=github.com/ckalagara/pub-a-player/core version=v3.5.5
2025-11-03T13:12:47.295364000-06:00 INF adding interface to collection collection=/Users/i/Src/pub-a-player/core/mocks.go interface=store package-path=github.com/ckalagara/pub-a-player/core version=v3.5.5
2025-11-03T13:12:47.295518000-06:00 INF Executing template file=/Users/i/Src/pub-a-player/core/mocks.go version=v3.5.5
2025-11-03T13:12:47.299119000-06:00 INF Writing template to file file=/Users/i/Src/pub-a-player/core/mocks.go version=v3.5.5
```

## Docker

```bash
docker build -t pub-a-player .  
docker run --network app-network --name pub-a-player -p 8089:8089 pub-a-player
```