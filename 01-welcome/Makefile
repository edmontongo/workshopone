SHELL=/bin/bash

.PHONY: all clean

all: pres.pdf

pres.pdf: pres.tex pres.bib
	latexmk -xelatex -quiet pres || { cat pres.log; false; }

clean:
	latexmk -C
	$(RM) pres.{bbl,bib.bak,nav,run.xml,snm,vrb}
