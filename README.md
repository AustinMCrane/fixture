# Fixture

Read and write structs to/from json files

### Save

save a struct to a file

```
a := Struct{
  Name: "test",
}

err := fixture.Save(&a, "example", "example-save.json")
if err != nil {
  log.Fataln(err)
}
```

### Read 

read a struct from a file

```
var a Struct

err := fixture.Read(&a, "example", "example-read.json")
if err != nil {
  log.Fataln(err)
}
```
