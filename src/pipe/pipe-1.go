package main


import "io"


func pipe (app1 func(in io.Reader, out io.Writer), app2 func(in io.Reader, out io.Writer)) func(in io.Reader, out io.Writer) {
	
	return func(in io.Reader, out io.Writer) {
		
		pr, pw := io.Pipe()
		
		defer pw.Close()
		
		go func() {
			defer pr.Close()
			app2(pr, out)
		}()
		
		app1(in, pw)
	}
}

