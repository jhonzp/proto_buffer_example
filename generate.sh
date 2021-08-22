protoc -I src\ --go_out=src\ src\simple\simple.proto
protoc -I src\ --go_out=src\ src\enum_example\enum_example.proto
protoc -I src\ --go_out=src\ src\compex\complex.proto
protoc -I src\ --go_out=src\ src\example\addressbook.proto