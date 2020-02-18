.PHONY: site
site:
	rm -rf docs &&	cd site && hugo && mv public ../docs
