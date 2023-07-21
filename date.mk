$(info Current OS: $(OS))
ifeq ($(OS),Windows_NT)
DATE_STRING_RAW:=$(shell date /t)$(shell time /t)
nully:=
space:=${nully} ${nully}
DATE_TIME=$(subst :,_,$(subst /,_,$(subst $(space),_,$(DATE_STRING_RAW))))
else
DATE_TIME:=$(shell date -u +"%Y_%m_%d_%H_%M_%S")
endif
$(info Date-time stamp: $(DATE_TIME))
