.PHONY: site
site:
	rm -rf docs &&	cd site && hugo --minify -t aether && mv public ../docs
