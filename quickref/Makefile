SHELL=/bin/bash

.PHONY: all clean

all: quickref.pdf

%.pdf: %.tex
	latexmk -quiet -xelatex "$*" || { cat "$*.log"; false; }

clean:
	latexmk -C
