PLUGIN_SOURCES = $(wildcard *.go)
PLUGIN_OBJECTS = $(PLUGIN_SOURCES:.go=.so)
GO_BUILD_ARGS = -buildmode=plugin

.PHONY: all clean

all: $(PLUGIN_OBJECTS)

%.so: %.go
	go build $(GO_BUILD_ARGS) -o $@ $<

clean:
	rm -f $(PLUGIN_OBJECTS)