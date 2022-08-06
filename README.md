# clip-list

can be used on macOS.

## Usage
### Show clips
```
$ clip-list show  
hoge
fuga
test
```
### Add clip
```
$ clip-list add example
Added example.

$ clip-list show  
hoge
fuga
test
example
```

### Select clip
```
$ clip-list select
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select clip: 
  ▸ hoge
    fuga
    test
 
 ✔ hoge
Copied to clipboard.
```

### Delete clip
```
$ clip-list delete
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select clip: 
  ▸ hoge
    fuga
    test

✔ hoge
Selected clip is deleted.

$ clip-list show
fuga
test
```
