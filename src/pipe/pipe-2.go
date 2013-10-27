package main

import "io"

func bind (app func(in io.Reader, out io.Writer, args []string), args []string) func(in io.Reader, out io.Writer) {
	return func (in io.Reader, out io.Writer) {
		app(in, out, args)
	}	
}

func pipe (apps ...func(in io.Reader, out io.Writer)) func(in io.Reader, out io.Writer) {
	if len(apps) == 0 {
		return nil
	}
	
	app := apps[0]
	
	for i := 1; i < len(apps); i++ {
		app1, app2 := app, apps[i]
		
		app = func(in io.Reader, out io.Writer) {
			pr, pw := io.Pipe()
						
			defer pw.Close()
			
			go func() {
				defer pr.Close()				
				app2(pr, out)
			}()
			
			app1(in, pw)
			
		}
	}
	
	return app
}


/*
 func tar(in io.Reader, out io.Writer, files []string)
 func gzip(in io.Reader, out io.Writer)
 
 pipe(bind(tar, files), gzip)
 
*/ 

func main() {
	
}