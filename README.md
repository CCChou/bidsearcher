# bidsearcher

Use keyword to search the bids' information

## Usage

Install js dependencies
```
cd frontend
npm install
```

Compile frontend page
```
npm run build
```

Modify the config.yml
```
username: xxx
password: yyy
```

Run application
```
go run ./cmd/main.go
```

Register the application as service
```
[Unit]
Description=BidSearcher Service
After=multi-user.target

[Service]
User=root
Group=root
WorkingDirectory=<directory_path_of_project>
ExecStart=/opt/go/bin/go run <directory_path_of_project>/cmd/main.go

[Install]
WantedBy=multi-user.target
```

## Reference
- [台灣採購公報網](https://www.taiwanbuying.com.tw/)
