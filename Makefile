NAME = Pinger


all:
	GOOS=linux GOARCH=amd64 go build -C cmd -o ../$(NAME).out

clear:
	#rm *.exe
	rm *.out

race:
	go build -C cmd --race -o ../$(NAME)_race.out

win:
	GOOS=windows GOARCH=amd64 go build -C cmd -o ../$(NAME).exe
