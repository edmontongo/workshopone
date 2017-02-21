SHELL=/bin/bash

.PHONY: all clean

all: pres_slides.pdf pres_notes.pdf

%.pdf: %.tex pres.tex
	latexmk -xelatex -quiet "$*" || { cat "$*.log"; false; }

clean:
	latexmk -C
	$(RM) pres{,_notes,_slides}.{bbl,bib.bak,nav,run.xml,snm,vrb}
