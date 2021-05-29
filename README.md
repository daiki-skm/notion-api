# notion-api

# Setup
Please go to [here](https://developers.notion.com/docs) to create Secrets ID and Database ID.

# Usage
```
export NOTION_KEY=your_secret_key
export BLOCK_ID=your_database_id
```

- all text
```
go run main.go
```

- only paragraph
```
go run main.go paragraph
```

- only todo
```
go run main.go todo
```

- only bullet
```
go run main.go bullet
```

- only toggle
```
go run main.go paragraph
```
