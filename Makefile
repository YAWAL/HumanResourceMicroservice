.PHONY: dependencies
dependencies:
	echo "Installing dependencies"
	dep ensure

.PHONY: code-quality
code-quality:
	gometalinter --vendor --tests --skip=mock --exclude='_gen.go' --disable=gotype --disable=errcheck --disable=gas --disable=dupl --deadline=1500s --checkstyle --sort=linter ./... > static-analysis.xml

install-helpers:
	echo "Installing GoMetaLinter"
	go get -u github.com/alecthomas/gometalinter
	echo "Installing linters"
	gometalinter --install
	echo "Installing Dep"
	go get -u github.com/golang/dep/cmd/dep

jaeger:
	docker run -d -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp \
	-p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:0.8.0