TARGET ?= wioterminal

build:
	tinygo build --target=$(TARGET)

flash:
	tinygo flash --target=$(TARGET)
