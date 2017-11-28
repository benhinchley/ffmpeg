# ffmpeg
> a go library interface around the [ffmpeg] command line tool

[![GoDoc](https://godoc.org/github.com/benhinchley/ffmpeg?status.svg)](https://godoc.org/github.com/benhinchley/ffmpeg)

## Usage
```go
// For creating a video from many images
cmd, err := ffmpeg.Command(nil,
	ffmpeg.Input("foo-%03d.jpeg"),
	ffmpeg.Output("foo.avi",
		ffmpeg.WithSize(ffmpeg.AllStreamSpecifier(), 1920,1080)))
if err != nil {
	// ... handle error
}

if err := cmd.Run(); err != nil {
	// ... handle error
}
```

## License
[MIT](LICENSE)

<!-- links -->
[ffmpeg]: http://ffmpeg.org/ffmpeg.html
