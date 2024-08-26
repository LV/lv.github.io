POSTS_DIR := content/posts

.PHONY: all
all:
	hugo server -D

post:
	hugo new $(POSTS_DIR)/$(wordlist 2, $(words $(MAKECMDGOALS)), $(MAKECMDGOALS)).md
	$(EDITOR) $(POSTS_DIR)/$(wordlist 2, $(words $(MAKECMDGOALS)), $(MAKECMDGOALS)).md

# Ignore arguments as targets
%:
	@:
