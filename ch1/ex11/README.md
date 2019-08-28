## Try fetchall with longer argument lists, such as samples from the top million web sites available at alexa.com. How does the program behave if a web site just doesn’t respond?

```shell script
$ go run gopl.io/ch1/fetchall http://8.8.8.8
... <hangs> ...
```

if website doesn't respon, the command will hang.