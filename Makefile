serve_doc:
	godoc -http=:6060

save_doc:
	cd chesslib && godoc -http=:6060 && 
		wget -p -k http://localhost:6060/pkg/github.com/dchudik/chessapp/chesslib/
