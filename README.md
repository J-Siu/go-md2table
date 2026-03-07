# go-md2table

Transform markdown file into markdown table.

## Example

```sh
go run main.go example/test.md
```

`example/test.md`:

```md
# a
line 1
## b
line 2
line 3
### c
line 4
line 5
line 6
#### d
line 7
line 8
line 9
line 10
```

To:

```md
||||||
|--|--|--|--|--|
|a|||| line 1|
||b||| line 2 line 3|
|||c|| line 4 line 5 line 6|
||||d| line 7 line 8 line 9 line 10
```
