# Multi-Package Support

If the user's code is split across multiple packages, generate serializers for 
all types in all packages. Corresponding `mus.ai.gen_test.go` MUST be generated 
for every `mus.ai.gen.go` file, including any subpackages or external packages.
