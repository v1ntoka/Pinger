NAME = Pnigger

all:
	go build -C cmd -o ../$(NAME).out

clear:
	#rm *.exe
	rm *.out

race:
	go build -C cmd --race -o ../$(NAME)_race.out
