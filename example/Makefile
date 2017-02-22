.PHONY: dredd

hooks/hooks:
	cd hooks/; go build; cd ..;

server:
	go build -o server

dredd: hooks/hooks server
	dredd ./api.apib http://127.0.0.1:8080 --language go --server ./server --hookfiles hooks/hooks --level silly

clean: 
	rm server hooks/hooks
