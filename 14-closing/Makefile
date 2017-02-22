SHELL=/bin/bash

.PHONY: all clean

all: pres_slides.pdf

%.pdf: %.tex pres.tex
	latexmk -xelatex -quiet "$*" || { cat "$*.log"; false; }

clean:
	latexmk -C
	$(RM) pres{,_slides}.{bbl,bib.bak,nav,run.xml,snm,vrb}
