.PHONY: dredd

hooks/hooks:
	cd hooks/; rm -f foo.db; touch foo.db; go build; cd ..;

server:
	go build -o server

dredd: hooks/hooks server
	../node_modules/.bin/dredd api.apib http://localhost:8080 --language go --server ./server --hookfiles hooks/hooks

clean: 
	rm server hooks/hooks
