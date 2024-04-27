all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./tfgen.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.22.2 amd64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/tfgen.exe $(DESTDIR)$(PREFIX)/bin/tfgen
	strip $(DESTDIR)$(PREFIX)/bin/tfgen
