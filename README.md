# Utils

This is very much a work in progress. Currently not a lot of tests, and it is just a flat module. I'll probably get to that soon.

The purpose of this package is for me to not having to write the same methods over and over for a bunch of golang projects.

I'll most likely re-structure it as I add more stuff. So don't consider it to never break.

## Methods
### func ExpandTilde(input string) string
Replaces the first ~ with the HOME env variable 
### func CreateUUID() string {
Generate a new UUID string
### func FileExist(filename string) bool {
Check if a file exists
### func RandomString() string
Generate a random string
### func GenerateRandomTempFilename() string
Generate a random filename in /tmp useful for unit tests
### func NilStringPointer() *string
Just returns a string pointer with nil
### func StringPointer(input string) *string
Converts a string to a string pointer. Mostly there for testing when I got really annoyed by how verbose it was to do it manually