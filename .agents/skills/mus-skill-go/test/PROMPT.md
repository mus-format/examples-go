# Generation Prompts

Prompts used to generate MUS serializers and unit tests for the project 
verification:

1. Using instructions from SKILL.md, generate MUS serializers for all types 
   found in `file_name`, the output save into `file_name_mus.ai.gen.go` and 
   `file_name_mus.ai.gen_test.go`.

2. Using instructions from SKILL.md, generate [stream?] MUS serializers in 
   `mode_name` mode for all types found in `full_struct.go`, the output save 
   into `test/mode/[stream_?]mode_name/mus.ai.gen.go` and 
   `test/mode/[stream_?]mode_name/mus.ai.gen_test.go`.

3. Using instructions from SKILL.md, generate MUS serializers for all types found in `test/genopts/pkg_name`.